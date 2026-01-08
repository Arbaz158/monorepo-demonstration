GO_SERVICES := services/go/user-service services/go/order-service services/go/payment-service
PY_SERVICES := services/python/ml-service services/python/notification-service
K8S_SERVICES := $(GO_SERVICES) $(PY_SERVICES)

# Docker Hub repository (change if you use a different registry)
# Example: DOCKER_REPO = myuser344/monorepo-demonstration
DOCKER_REPO ?= arbaz344/monorepo-demonstration

.PHONY: build-go build-python docker-build-all docker-build-go docker-build-python docker-build-push-go docker-build-push-python docker-build-% docker-build-push-% docker-login docker-check-login

docker-login:
	@echo "Logging into Docker Hub..."
	@docker login -u arbaz344
	@echo "Login successful!"

docker-check-login:
	@echo "Checking Docker Hub login status..."
	@if [ -f ~/.docker/config.json ] && grep -q "https://index.docker.io/v1/" ~/.docker/config.json 2>/dev/null; then \
		echo "✓ Docker Hub credentials found"; \
	else \
		echo "✗ Not logged in to Docker Hub. Run 'make docker-login' first."; \
		exit 1; \
	fi

build-go:
	@echo "Building all Go services (single module at services/go)"
	@cd services/go && go build ./...

build-python:
	@for s in $(PY_SERVICES); do \
		echo "Checking $$s"; \
		python -m py_compile $$(find $$s -maxdepth 1 -name '*.py'); \
	done

# Build all services and push to Docker Hub
docker-build-all: docker-check-login
	@for s in $(GO_SERVICES) $(PY_SERVICES); do \
		name=$$(basename $$s); \
		echo "Building docker image $(DOCKER_REPO):$$name"; \
		docker build -t $(DOCKER_REPO):$$name -f $$s/Dockerfile .; \
		echo "Pushing docker image $(DOCKER_REPO):$$name"; \
		docker push $(DOCKER_REPO):$$name; \
	done

# Build all Go services (without pushing)
docker-build-go:
	@for s in $(GO_SERVICES); do \
		name=$$(basename $$s); \
		echo "Building docker image $(DOCKER_REPO):$$name"; \
		docker build -t $(DOCKER_REPO):$$name -f $$s/Dockerfile .; \
	done

# Build all Python services (without pushing)
docker-build-python:
	@for s in $(PY_SERVICES); do \
		name=$$(basename $$s); \
		echo "Building docker image $(DOCKER_REPO):$$name"; \
		docker build -t $(DOCKER_REPO):$$name -f $$s/Dockerfile .; \
	done

# Build and push all Go services to Docker Hub
docker-build-push-go: docker-check-login
	@for s in $(GO_SERVICES); do \
		name=$$(basename $$s); \
		echo "Building docker image $(DOCKER_REPO):$$name"; \
		docker build -t $(DOCKER_REPO):$$name -f $$s/Dockerfile .; \
		echo "Pushing docker image $(DOCKER_REPO):$$name"; \
		docker push $(DOCKER_REPO):$$name; \
	done

# Build and push all Python services to Docker Hub
docker-build-push-python: docker-check-login
	@for s in $(PY_SERVICES); do \
		name=$$(basename $$s); \
		echo "Building docker image $(DOCKER_REPO):$$name"; \
		docker build -t $(DOCKER_REPO):$$name -f $$s/Dockerfile .; \
		echo "Pushing docker image $(DOCKER_REPO):$$name"; \
		docker push $(DOCKER_REPO):$$name; \
	done

# Build and push a single service to Docker Hub (MUST come before docker-build-%)
# Usage: make docker-build-push-user-service OR make docker-build-push-notification-service
docker-build-push-%: docker-check-login
	@service_path=$$(if [ -f services/go/$*/Dockerfile ]; then echo services/go/$*; elif [ -f services/python/$*/Dockerfile ]; then echo services/python/$*; else echo ""; fi); \
	if [ -z "$$service_path" ]; then \
		echo "✗ Service '$*' not found. Available services:"; \
		echo "  Go: user-service, order-service, payment-service"; \
		echo "  Python: ml-service, notification-service"; \
		exit 1; \
	else \
		echo "Building docker image $(DOCKER_REPO):$*"; \
		docker build -t $(DOCKER_REPO):$* -f $$service_path/Dockerfile .; \
		echo "Pushing docker image $(DOCKER_REPO):$*"; \
		docker push $(DOCKER_REPO):$*; \
	fi

# Build a single service (without pushing)
# Usage: make docker-build-user-service OR make docker-build-notification-service
docker-build-%:
	@service_path=$$(if [ -f services/go/$*/Dockerfile ]; then echo services/go/$*; elif [ -f services/python/$*/Dockerfile ]; then echo services/python/$*; else echo ""; fi); \
	if [ -z "$$service_path" ]; then \
		echo "✗ Service '$*' not found. Available services:"; \
		echo "  Go: user-service, order-service, payment-service"; \
		echo "  Python: ml-service, notification-service"; \
		exit 1; \
	else \
		echo "Building docker image $(DOCKER_REPO):$*"; \
		docker build -t $(DOCKER_REPO):$* -f $$service_path/Dockerfile .; \
	fi

.PHONY: k8s-apply-all k8s-delete-all k8s-apply-go k8s-apply-python k8s-delete-go k8s-delete-python k8s-apply-% k8s-delete-%

k8s-apply-all:
	@for s in $(K8S_SERVICES); do \
		echo "Applying $$s k8s manifests"; \
		kubectl apply -f $$s/k8s; \
	done

k8s-delete-all:
	@for s in $(K8S_SERVICES); do \
		echo "Deleting $$s k8s manifests"; \
		kubectl delete -f $$s/k8s --ignore-not-found; \
	done

# Apply all Go service manifests
k8s-apply-go:
	@for s in $(GO_SERVICES); do \
		echo "Applying $$s k8s manifests"; \
		kubectl apply -f $$s/k8s; \
	done

# Apply all Python service manifests
k8s-apply-python:
	@for s in $(PY_SERVICES); do \
		echo "Applying $$s k8s manifests"; \
		kubectl apply -f $$s/k8s; \
	done

# Delete all Go service manifests
k8s-delete-go:
	@for s in $(GO_SERVICES); do \
		echo "Deleting $$s k8s manifests"; \
		kubectl delete -f $$s/k8s --ignore-not-found; \
	done

# Delete all Python service manifests
k8s-delete-python:
	@for s in $(PY_SERVICES); do \
		echo "Deleting $$s k8s manifests"; \
		kubectl delete -f $$s/k8s --ignore-not-found; \
	done

# Usage: make k8s-apply-user-service  OR make k8s-delete-user-service
k8s-apply-%:
	@kubectl apply -f services/go/$*/k8s 2>/dev/null || kubectl apply -f services/python/$*/k8s

k8s-delete-%:
	@kubectl delete -f services/go/$*/k8s --ignore-not-found 2>/dev/null || kubectl delete -f services/python/$*/k8s --ignore-not-found
