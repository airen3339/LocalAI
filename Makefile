GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=local-ai
GOLLAMA_VERSION?=llama.cpp-8687c1f
GOGPT4ALLJ_VERSION?=1f548782d80d48b9a0fac33aae6f129358787bc0
GOGPT2_VERSION?=1c24f5b86ac428cd5e81dae1f1427b1463bd2b06

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

C_INCLUDE_PATH=$(shell pwd)/go-llama:$(shell pwd)/go-gpt4all-j:$(shell pwd)/go-gpt2 
LIBRARY_PATH=$(shell pwd)/go-llama:$(shell pwd)/go-gpt4all-j:$(shell pwd)/go-gpt2

# Use this if you want to set the default behavior
ifndef BUILD_TYPE
	BUILD_TYPE:=default
endif

ifeq ($(BUILD_TYPE), "generic")
	GENERIC_PREFIX:=generic-
else
	GENERIC_PREFIX:=
endif

.PHONY: all test build vendor

all: help

## Build:

build: prepare ## Build the project
	$(info ${GREEN}I local-ai build info:${RESET})
	$(info ${GREEN}I BUILD_TYPE: ${YELLOW}$(BUILD_TYPE)${RESET})
	C_INCLUDE_PATH=${C_INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} $(GOCMD) build -o $(BINARY_NAME) ./

generic-build: ## Build the project using generic
	BUILD_TYPE="generic" $(MAKE) build

## GPT4ALL-J
go-gpt4all-j:
	git clone --recurse-submodules https://github.com/go-skynet/go-gpt4all-j.cpp go-gpt4all-j
	cd go-gpt4all-j && git checkout -b build $(GOGPT4ALLJ_VERSION)
	# This is hackish, but needed as both go-llama and go-gpt4allj have their own version of ggml..
	@find ./go-gpt4all-j -type f -name "*.c" -exec sed -i'' -e 's/ggml_/ggml_gptj_/g' {} +
	@find ./go-gpt4all-j -type f -name "*.cpp" -exec sed -i'' -e 's/ggml_/ggml_gptj_/g' {} +
	@find ./go-gpt4all-j -type f -name "*.h" -exec sed -i'' -e 's/ggml_/ggml_gptj_/g' {} +
	@find ./go-gpt4all-j -type f -name "*.cpp" -exec sed -i'' -e 's/gpt_/gptj_/g' {} +
	@find ./go-gpt4all-j -type f -name "*.h" -exec sed -i'' -e 's/gpt_/gptj_/g' {} +
	@find ./go-gpt4all-j -type f -name "*.cpp" -exec sed -i'' -e 's/json_/json_gptj_/g' {} +
	@find ./go-gpt4all-j -type f -name "*.cpp" -exec sed -i'' -e 's/void replace/void json_gptj_replace/g' {} +
	@find ./go-gpt4all-j -type f -name "*.cpp" -exec sed -i'' -e 's/::replace/::json_gptj_replace/g' {} +

go-gpt4all-j/libgptj.a: go-gpt4all-j
	$(MAKE) -C go-gpt4all-j $(GENERIC_PREFIX)libgptj.a

# CEREBRAS GPT
go-gpt2:
	git clone --recurse-submodules https://github.com/go-skynet/go-gpt2.cpp go-gpt2
	cd go-gpt2 && git checkout -b build $(GOGPT2_VERSION)
	# This is hackish, but needed as both go-llama and go-gpt4allj have their own version of ggml..
	@find ./go-gpt2 -type f -name "*.c" -exec sed -i'' -e 's/ggml_/ggml_gpt2_/g' {} +
	@find ./go-gpt2 -type f -name "*.cpp" -exec sed -i'' -e 's/ggml_/ggml_gpt2_/g' {} +
	@find ./go-gpt2 -type f -name "*.h" -exec sed -i'' -e 's/ggml_/ggml_gpt2_/g' {} +
	@find ./go-gpt2 -type f -name "*.cpp" -exec sed -i'' -e 's/gpt_/gpt2_/g' {} +
	@find ./go-gpt2 -type f -name "*.h" -exec sed -i'' -e 's/gpt_/gpt2_/g' {} +
	@find ./go-gpt2 -type f -name "*.cpp" -exec sed -i'' -e 's/json_/json_gpt2_/g' {} +

go-gpt2/libgpt2.a: go-gpt2
	$(MAKE) -C go-gpt2 $(GENERIC_PREFIX)libgpt2.a
	

go-llama:
	git clone -b $(GOLLAMA_VERSION) --recurse-submodules https://github.com/go-skynet/go-llama.cpp go-llama

go-llama/libbinding.a: go-llama
	$(MAKE) -C go-llama $(GENERIC_PREFIX)libbinding.a

replace:
	$(GOCMD) mod edit -replace github.com/go-skynet/go-llama.cpp=$(shell pwd)/go-llama
	$(GOCMD) mod edit -replace github.com/go-skynet/go-gpt4all-j.cpp=$(shell pwd)/go-gpt4all-j
	$(GOCMD) mod edit -replace github.com/go-skynet/go-gpt2.cpp=$(shell pwd)/go-gpt2

prepare: go-llama/libbinding.a go-gpt4all-j/libgptj.a go-gpt2/libgpt2.a replace

clean: ## Remove build related file
	rm -fr ./go-llama
	rm -rf ./go-gpt4all-j
	rm -rf ./go-gpt2
	rm -rf $(BINARY_NAME)

## Run:
run: prepare
	C_INCLUDE_PATH=${C_INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} $(GOCMD) run ./main.go

test-models/testmodel:
	mkdir test-models
	wget https://huggingface.co/concedo/cerebras-111M-ggml/resolve/main/cerberas-111m-q4_0.bin -O test-models/testmodel

test: prepare test-models/testmodel
	@C_INCLUDE_PATH=${C_INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} MODELS_PATH=$(abspath ./)/test-models $(GOCMD) test -v ./...

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
