#!/bin/bash
set -e

echo "🔍 Testing CI/CD configuration..."

# Check workflow syntax
echo "✅ Validating workflow syntax..."
find .github/workflows -name "*.yml" -exec python3 -c "import yaml; yaml.safe_load(open('{}'))" \; > /dev/null

# Check required files exist
echo "✅ Checking required files..."
[ -f ".github/workflows/ci.yml" ] || exit 1
[ -f ".github/workflows/cd.yml" ] || exit 1
[ -f ".github/environments/staging.md" ] || exit 1
[ -f ".github/environments/production.md" ] || exit 1

# Validate Kubernetes manifests
echo "✅ Validating Kubernetes manifests..."
cd k8s
kustomize build base/ > /dev/null
kustomize build staging/ > /dev/null  
kustomize build production/ > /dev/null
cd ..

echo "🎉 CI/CD configuration is valid!"
