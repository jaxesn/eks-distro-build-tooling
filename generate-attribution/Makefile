# Disable built-in rules and variables
MAKEFLAGS+=--no-builtin-rules --warn-undefined-variables
SHELL=bash
.SHELLFLAGS:=-eu -o pipefail -c
.SUFFIXES:
.SECONDEXPANSION:

BASE_DIRECTORY=$(shell git rev-parse --show-toplevel)
MAKE_ROOT=$(BASE_DIRECTORY)/generate-attribution

AWS_ACCOUNT_ID?=$(shell aws sts get-caller-identity --query Account --output text)
AWS_REGION?=us-west-2

IMAGE_REPO?=$(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com
BASE_IMAGE_REPO?=$(IMAGE_REPO)
JOB_TYPE?=

PLATFORMS?=linux/amd64,linux/arm64

IMAGE_TAG?=$(shell date "+%F-%s")
IMAGE?=$(IMAGE_REPO)/$(IMAGE_NAME):$(call ADD_VERSIONED_TO_TAG,$(IMAGE_TAG))

LATEST_TAGS=$(foreach tag,latest,$(IMAGE_REPO)/$(IMAGE_NAME):$(call ADD_VERSIONED_TO_TAG,$(tag)))
LATEST_IMAGE=$(shell echo $(LATEST_TAGS) | sed "s/ \+/,/g")

BASE_IMAGE=$(BASE_IMAGE_REPO)/$(BASE_IMAGE_NAME):$(BASE_IMAGE_TAG)

BUILDER_IMAGE_NAME?=$(BASE_IMAGE_NAME)-builder
BUILDER_IMAGE=$(BASE_IMAGE_REPO)/$(BUILDER_IMAGE_NAME):$(BASE_IMAGE_TAG)

IMAGE_OUTPUT_TYPE=image
IMAGE_OUTPUT=push=true

GOPROXY?=direct


ADD_VERSIONED_TO_TAG=$(if $(VERSIONED_VARIANT),$(VERSIONED_VARIANT)-$(1),$(1))

define BUILDCTL
	buildctl \
		build \
		--frontend dockerfile.v0 \
		--opt filename=$(DOCKERFILE) \
		--opt platform=$(PLATFORMS) \
		--opt build-arg:BASE_IMAGE=$(BASE_IMAGE) \
		--opt build-arg:GOPROXY=$(GOPROXY) \
		--opt build-arg:BUILDER_IMAGE=$(BUILDER_IMAGE) \
		--progress plain \
		--local dockerfile=./ \
		--local context=. \
		--output type=$(IMAGE_OUTPUT_TYPE),oci-mediatypes=true,\"name=$(IMAGE),$(LATEST_IMAGE)\",$(IMAGE_OUTPUT)
endef


## --------------------------------------
## Help
## --------------------------------------
##@ Helpers
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-35s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: buildkit-check
buildkit-check:
	$(MAKE_ROOT)/../scripts/buildkit_check.sh

.PHONY: generate-attribution-image
generate-attribution-image: IMAGE_NAME=generate-attribution
generate-attribution-image: BASE_IMAGE_NAME=eks-distro-minimal-base-nodejs
generate-attribution-image: DOCKERFILE=Dockerfile.generate-attribution
generate-attribution-image: BASE_IMAGE_TAG=$(shell yq e ".al2022.\"eks-distro-minimal-base-nodejs18\"" $(MAKE_ROOT)/../EKS_DISTRO_TAG_FILE.yaml)
generate-attribution-image: buildkit-check
	$(call BUILDCTL)


.PHONY: go-licenses-image-%
go-licenses-image-%: IMAGE_NAME=go-licenses
go-licenses-image-%: VERSIONED_VARIANT=$*
go-licenses-image-%: BASE_IMAGE_NAME=eks-distro-minimal-base-golang
go-licenses-image-%: DOCKERFILE=Dockerfile.go-licenses
go-licenses-image-%: BASE_IMAGE_TAG=$(shell yq e ".al2.\"eks-distro-minimal-base-golang$(VERSIONED_VARIANT)\"" $(MAKE_ROOT)/../EKS_DISTRO_TAG_FILE.yaml)
go-licenses-image-%: buildkit-check
	$(call BUILDCTL)

.PHONY: go-licenses-images
go-licenses-images: $(addprefix go-licenses-image-,1.19 1.18 1.17 1.16)

.PHONY: images
images: generate-attribution-image go-licenses-images

.PHONY: build
build: images

.PHONY: release
release: images

.PHONY: all
all: release
