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

## Deploy a Single Service

After building and pushing images (`make docker-build-all`):

```bash
make k8s-apply-user-service          # or order-service, payment-service, ml-service, notification-service
```

Delete one:
```bash
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

