# Variables
KIND_CLUSTER_NAME := polyaxon-cluster
NAMESPACE := polyaxon
HELM_RELEASE_NAME := polyaxon
HELM_CHART := ./charts/polyaxon
INGRESS_URL := polyaxon.local
VALUES_FILE := polyaxon.yml

# Paths
KIND_CONFIG := kind.yml

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build deploy

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## Build Docker image
.PHONY: build
build:
	docker build -t polyaxon/cli:latest -f cli/.docker/Dockerfile cli
	docker build -t polyaxon/agent:latest -f cli/.docker/Dockerfile.agent cli
	docker build -t polyaxon/init:latest -f cli/.docker/Dockerfile.init cli
	docker build -t polyaxon/notifier:latest -f cli/.docker/Dockerfile.notifier cli
	docker build -t polyaxon/scheduler:latest -f cli/.docker/Dockerfile.scheduler cli
	docker build -t polyaxon/sidecar:latest -f cli/.docker/Dockerfile.sidecar cli

## Create Kind cluster
.PHONY: create-cluster
create-cluster: 
	@echo "Creating Kind cluster with configuration from $(KIND_CONFIG)..."
	kind create cluster --name $(KIND_CLUSTER_NAME) --config $(KIND_CONFIG)
	@echo "Kind cluster $(KIND_CLUSTER_NAME) created."

## Delete Kind cluster
.PHONY: delete-cluster
delete-cluster:
	@echo "Deleting Kind cluster..."
	kind delete cluster --name $(KIND_CLUSTER_NAME)
	@echo "Kind cluster $(KIND_CLUSTER_NAME) deleted."

## Deploy Polyaxon CE using Helm
.PHONY: install-polyaxon
install-polyaxon:
	@echo "Installing Polyaxon Helm chart..."
	helm install $(HELM_RELEASE_NAME) $(HELM_CHART) -f $(VALUES_FILE) --namespace $(NAMESPACE) --create-namespace
	@echo "Polyaxon deployed successfully."

## Uninstall Polyaxon CE
.PHONY: uninstall-polyaxon
uninstall-polyaxon:
	@echo "Uninstalling Polyaxon Helm release..."
	helm uninstall $(HELM_RELEASE_NAME) --namespace $(NAMESPACE)
	@echo "Polyaxon uninstalled."

## Deploy everything
.PHONY: deploy
deploy: create-cluster install-polyaxon
	@echo "Mapping $(INGRESS_URL) to localhost..."
	@echo "127.0.0.1 $(INGRESS_URL)" | sudo tee -a /etc/hosts > /dev/null
	@echo "Polyaxon is accessible at http://$(INGRESS_URL)"

## Clean up the environment
.PHONY: clean
clean: uninstall-polyaxon delete-cluster
	@echo "Cleaning up /etc/hosts..."
	sudo sed -i '/$(INGRESS_URL)/d' /etc/hosts
	@echo "Cleanup complete."