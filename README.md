# Monorepo Demonstration

Polyglot demo with Go and Python services. Works with **any Kubernetes cluster** (minikube, EKS, GKE, AKS, on-prem, etc.). Uses a single Go module (`services/go/go.mod`) and shared Python requirements (`services/python/requirements.txt`). Images are built and pushed to Docker Hub for universal cluster access.

## Layout
- `services/go/` — Go module root  
  - `common/` shared Go libs  
  - `{user-service,order-service,payment-service}/` Go services + Dockerfile + k8s/  
- `services/python/` — Python services + shared `requirements.txt`  
  - `{ml-service,notification-service}/` with Dockerfile + k8s/  
- `Makefile` — build & deploy helpers  
- `scripts/` — reserved for tooling hooks

## Prerequisites
- Docker
- Go 1.21+
- Python 3.12+ (if running Python services directly)
- kubectl (configured for your target cluster)
- make
- Docker Hub account (for pushing images)

## Clone
```bash
git clone <repo-url>
cd monorepo-demonstration
```

## Build & Deploy All (Any Kubernetes Cluster)

1) **Login to Docker Hub** (one-time per session):
```bash
make docker-login
# Enter your Docker Hub password/Personal Access Token when prompted
```

2) **Build and push all images to Docker Hub**:
```bash
make docker-build-all
```
This builds and pushes images to `arbaz344/monorepo-demonstration:<service-name>` (configurable via `DOCKER_REPO` in Makefile).

3) **Deploy to your Kubernetes cluster**:
```bash
# Make sure kubectl is pointing to your target cluster
kubectl config current-context  # verify your cluster

# Apply all manifests
make k8s-apply-all
```

4) **Verify deployment**:
```bash
kubectl get pods,svc
kubectl get pods -w  # watch pod status
```

## Local Development (Minikube)

For local minikube development, you can optionally skip Docker Hub and use local images:

```bash
# Point Docker to minikube daemon
eval $(minikube -p minikube docker-env)

# Build images locally (without pushing)
make docker-build-all

# Load images into minikube
minikube image load arbaz344/monorepo-demonstration:user-service
minikube image load arbaz344/monorepo-demonstration:order-service
minikube image load arbaz344/monorepo-demonstration:payment-service
minikube image load arbaz344/monorepo-demonstration:ml-service
minikube image load arbaz344/monorepo-demonstration:notification-service

# Deploy
make k8s-apply-all
```

**Note:** The standard workflow (Docker Hub) works for minikube too and is recommended for consistency.

## Teardown All
```bash
make k8s-delete-all
```

## Build a Single Service

You can build individual services without building everything:

**From the monorepo root:**

```bash
# Build only (no push)
make docker-build-notification-service
make docker-build-ml-service
make docker-build-user-service
make docker-build-order-service
make docker-build-payment-service

# Build and push to Docker Hub
make docker-build-push-notification-service
make docker-build-push-ml-service
# ... etc
```

**Important:** Dockerfiles expect the build context to be the monorepo root (because they reference `services/python/requirements.txt` and similar paths). Always run `docker build` commands from the **root directory**, not from within service directories.

If you want to build from the command line directly:
```bash
# From the monorepo root directory:
docker build -t arbaz344/monorepo-demonstration:notification-service -f services/python/notification-service/Dockerfile .
```

## Build Services by Language

You can build and push services grouped by language:

**Go services only:**
```bash
# Build all Go services (no push)
make docker-build-go

# Build and push all Go services
make docker-build-push-go
```

**Python services only:**
```bash
# Build all Python services (no push)
make docker-build-python

# Build and push all Python services
make docker-build-push-python
```

## Deploy Services

**Deploy all services:**
```bash
make k8s-apply-all
```

**Deploy by language:**
```bash
# Deploy all Go services
make k8s-apply-go

# Deploy all Python services
make k8s-apply-python
```

**Deploy a single service:**
```bash
make k8s-apply-user-service          # or order-service, payment-service, ml-service, notification-service
```

**Delete services:**
```bash
# Delete all
make k8s-delete-all

# Delete by language
make k8s-delete-go
make k8s-delete-python

# Delete single service
make k8s-delete-user-service
```

## Service Ports & Access

- **Go services**: Listen on port 8080 in-container; Service objects expose port 80.
- **Python services**: `ml-service` listens on 5000; `notification-service` on 5001; each Service exposes port 80.

**Accessing services:**
- **Any cluster**: Use `kubectl port-forward svc/<service-name> 8080:80` to access locally
- **Minikube**: `minikube service <service-name> --url`
- **Cloud clusters**: Configure LoadBalancer/Ingress as needed

## Configuration

- **Docker Hub repository**: Set `DOCKER_REPO` in Makefile (default: `arbaz344/monorepo-demonstration`)
- **Go module root**: `services/go/go.mod`. Build all Go services together: `cd services/go && go build ./...` (Makefile automates).
- **Python shared deps**: `services/python/requirements.txt` (both Python Dockerfiles install from it).
- **Image tags**: Images are tagged as `<DOCKER_REPO>:<service-name>` (e.g., `arbaz344/monorepo-demonstration:user-service`)

## Cloud Deployment Notes

✅ **Works out-of-the-box** with:
- **EKS (AWS)**: No additional configuration needed for public Docker Hub images
- **GKE (Google Cloud)**: No additional configuration needed
- **AKS (Azure)**: No additional configuration needed
- **On-prem/Private clusters**: Ensure cluster nodes can access Docker Hub (or configure private registry)

For **private registries**, update `DOCKER_REPO` in Makefile and add `imagePullSecrets` to deployments if required.


# Command to delete all the images of a particluar docker hub repo
docker images arbaz344/monorepo-demonstration -q | xargs docker rmi -f

# Forward order-service to local port 8080
kubectl port-forward svc/order-service 8080:80

# Test health endpoint
curl http://localhost:8080/health

# Or test orders endpoint
curl http://localhost:8080/orders


