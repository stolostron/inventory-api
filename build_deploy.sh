set -exv

if [[ -z "$IMAGE" ]]; then
    IMAGE="quay.io/cloudservices/kessel-inventory"
fi
IMAGE_TAG=$(git rev-parse --short=7 HEAD)
GIT_COMMIT=$(git rev-parse --short HEAD)

if [[ -z "$QUAY_USER" || -z "$QUAY_TOKEN" ]]; then
    echo "QUAY_USER and QUAY_TOKEN must be set"
    exit 1
fi

if [[ -z "$RH_REGISTRY_USER" || -z "$RH_REGISTRY_TOKEN" ]]; then
    echo "RH_REGISTRY_USER and RH_REGISTRY_TOKEN  must be set"
    exit 1
fi

DOCKER_CONF="$PWD/.docker"
mkdir -p "$DOCKER_CONF"
docker --config="$DOCKER_CONF" login -u="$QUAY_USER" -p="$QUAY_TOKEN" quay.io
docker --config="$DOCKER_CONF" login -u="$RH_REGISTRY_USER" -p="$RH_REGISTRY_TOKEN" registry.redhat.io
docker --config="$DOCKER_CONF" build --build-arg GIT_COMMIT=$GIT_COMMIT --no-cache -t "${IMAGE}:${IMAGE_TAG}" . -f ./Dockerfile
docker --config="$DOCKER_CONF" build --build-arg GIT_COMMIT=$GIT_COMMIT --no-cache -t "${IMAGE}:latest" . -f ./Dockerfile
docker --config="$DOCKER_CONF" push "${IMAGE}:${IMAGE_TAG}"
docker --config="$DOCKER_CONF" push "${IMAGE}:latest"
