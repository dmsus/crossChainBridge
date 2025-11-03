# Kubernetes Deployment for Cross-Chain Bridge

## Structure
k8s/
├── base/ # Base configuration
├── staging/ # Staging environment
├── production/ # Production environment
└── README.md


## Prerequisites
- Kubernetes 1.25+
- kubectl configured
- Kustomize 4.5+

## Quick Start

### Apply Base Configuration
```bash
kubectl apply -k k8s/base
```
## Deploy to Staging
```bash
kubectl apply -k k8s/staging
```
## Deploy to Production
```bash
kubectl apply -k k8s/production
```
## Components
## Relayer
- Deployment with 2 replicas (staging: 1, production: 3)

- Horizontal Pod Autoscaler (2-10 pods)

- Resource limits: 128Mi-512Mi memory, 100m-500m CPU

- Health checks on port 8080 (/health, /ready)

- Metrics on port 9090 (/metrics)

## PostgreSQL
- Single replica deployment

- Persistent volume (10Gi)

- Health checks with pg_isready

## Security
- Non-root execution

- Network policies

- Read-only root filesystem

- No privilege escalation

## Monitoring
- Prometheus metrics automatically discovered

- Service annotations for scraping

- RBAC for cluster monitoring

## Secrets Management
Replace placeholder values in secrets before deployment:

```bash
# Generate base64 values
echo -n "your-password" | base64
echo -n "your-private-key" | base64

# Update k8s/base/relayer/secret.yaml
# Update k8s/base/postgres/secret.yaml
```
## Image Registry
Update image references in kustomization.yaml:

```yaml
images:
  - name: crosschainbridge-relayer
    newName: your-registry/crosschainbridge-relayer
    newTag: latest
```