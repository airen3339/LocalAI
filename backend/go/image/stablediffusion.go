package main

// This is a wrapper to statisfy the GRPC service interface
// It is meant to be used by the main executable that is the server for the specific backend type (falcon, gpt3, etc)
import (
	"encoding/base64"
	"os"

	"github.com/go-skynet/LocalAI/pkg/grpc/base"
	pb "github.com/go-skynet/LocalAI/pkg/grpc/proto"
	"github.com/go-skynet/LocalAI/pkg/stablediffusion"
)

type StableDiffusion struct {
	base.SingleThread
	stablediffusion *stablediffusion.StableDiffusion
}

func (sd *StableDiffusion) Load(opts *pb.ModelOptions) error {
	var err error
	// Note: the Model here is a path to a directory containing the model files
	sd.stablediffusion, err = stablediffusion.New(opts.ModelFile)
	return err
}

func (sd *StableDiffusion) GenerateImage(opts *pb.GenerateImageRequest) (string, error) {
	// Temporary file from os.Tempdir
	f, err := os.CreateTemp("", "stablediffusion")
	if err != nil {
		return "", err
	}

	f.Close()
	output := f.Name() + ".jpg"

	// Rename the temporary file
	err = os.Rename(f.Name(), output)
	if err != nil {
		return "", err
	}

	err = sd.stablediffusion.GenerateImage(
		int(opts.Height),
		int(opts.Width),
		int(opts.Mode),
		int(opts.Step),
		int(opts.Seed),
		opts.PositivePrompt,
		opts.NegativePrompt,
		output)
	if err != nil {
		return "", err
	}

	d, err := os.ReadFile(output)
	if err != nil {
		return "", err
	}
	//Base64 encode it
	return base64.StdEncoding.EncodeToString(d), nil
}
