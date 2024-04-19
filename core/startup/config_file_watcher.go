package startup

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-skynet/LocalAI/core/config"
	"github.com/imdario/mergo"
	"github.com/rs/zerolog/log"
)

type fileHandler func(fileContent []byte, appConfig *config.ApplicationConfig) error

type configFileHandler struct {
	handlers map[string]fileHandler

	watcher *fsnotify.Watcher

	configDir string
	appConfig *config.ApplicationConfig
}

// TODO: This should be a singleton eventually so other parts of the code can register config file handlers,
// then we can export it to other packages
func newConfigFileHandler(appConfig *config.ApplicationConfig) configFileHandler {
	c := configFileHandler{
		handlers:  make(map[string]fileHandler),
		configDir: appConfig.DynamicConfigsDir,
		appConfig: appConfig,
	}
	c.Register("api_keys.json", readApiKeysJson(*appConfig), true)
	c.Register("external_backends.json", readExternalBackendsJson(*appConfig), true)
	return c
}

func (c *configFileHandler) Register(filename string, handler fileHandler, runNow bool) error {
	_, ok := c.handlers[filename]
	if ok {
		return fmt.Errorf("handler already registered for file %s", filename)
	}
	c.handlers[filename] = handler
	if runNow {
		c.callHandler(path.Join(c.appConfig.DynamicConfigsDir, filename), handler)
	}
	return nil
}

func (c *configFileHandler) callHandler(filename string, handler fileHandler) {
	fileContent, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		log.Error().Err(err).Str("filename", filename).Msg("could not read file")
	}

	if err = handler(fileContent, c.appConfig); err != nil {
		log.Error().Err(err).Msg("WatchConfigDirectory goroutine failed to update options")
	}
}

func (c *configFileHandler) Watch() error {
	configWatcher, err := fsnotify.NewWatcher()
	c.watcher = configWatcher
	if err != nil {
		log.Fatal().Err(err).Str("configdir", c.configDir).Msg("wnable to create a watcher for configuration directory")
	}

	if c.appConfig.DynamicConfigsDirPollInterval > 0 {
		log.Debug().Msg("Poll interval set, falling back to polling for configuration changes")
		ticker := time.NewTicker(c.appConfig.DynamicConfigsDirPollInterval)
		go func() {
			for {
				<-ticker.C
				for file, handler := range c.handlers {
					log.Debug().Str("file", file).Msg("processing config file")
					c.callHandler(file, handler)
				}
			}
		}()
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-c.watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write | fsnotify.Create | fsnotify.Remove) {
					handler, ok := c.handlers[path.Base(event.Name)]
					if !ok {
						continue
					}

					c.callHandler(event.Name, handler)
				}
			case err, ok := <-c.watcher.Errors:
				log.Error().Err(err).Msg("config watcher error received")
				if !ok {
					return
				}
			}
		}
	}()

	// Add a path.
	err = c.watcher.Add(c.appConfig.DynamicConfigsDir)
	if err != nil {
		return fmt.Errorf("unable to establish watch on the LocalAI Configuration Directory: %+v", err)
	}

	return nil
}

// TODO: When we institute graceful shutdown, this should be called
func (c *configFileHandler) Stop() {
	c.watcher.Close()
}

func readApiKeysJson(startupAppConfig config.ApplicationConfig) fileHandler {
	handler := func(fileContent []byte, appConfig *config.ApplicationConfig) error {
		log.Debug().Msg("processing api_keys.json")

		if len(fileContent) > 0 {
			// Parse JSON content from the file
			var fileKeys []string
			err := json.Unmarshal(fileContent, &fileKeys)
			if err != nil {
				return err
			}

			appConfig.ApiKeys = append(startupAppConfig.ApiKeys, fileKeys...)
		} else {
			appConfig.ApiKeys = startupAppConfig.ApiKeys
		}
		log.Debug().Msg("api keys loaded from api_keys.json")
		return nil
	}

	return handler
}

func readExternalBackendsJson(startupAppConfig config.ApplicationConfig) fileHandler {
	handler := func(fileContent []byte, appConfig *config.ApplicationConfig) error {
		log.Debug().Msg("processing external_backends.json")

		if len(fileContent) > 0 {
			// Parse JSON content from the file
			var fileBackends map[string]string
			err := json.Unmarshal(fileContent, &fileBackends)
			if err != nil {
				return err
			}
			appConfig.ExternalGRPCBackends = startupAppConfig.ExternalGRPCBackends
			err = mergo.Merge(&appConfig.ExternalGRPCBackends, &fileBackends)
			if err != nil {
				return err
			}
		} else {
			appConfig.ExternalGRPCBackends = startupAppConfig.ExternalGRPCBackends
		}
		log.Debug().Msg("external backends loaded from external_backends.json")
		return nil
	}
	return handler
}
