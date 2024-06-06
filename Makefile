VERSION="v1.0.3"
APPNAME="fluid"
OASNAME="oas"

install:
	podman build . --no-cache -f build/openapi.Containerfile -t ${OASNAME}:v2.0

build::
	podman build . -f build/Containerfile -t ${APPNAME}:${VERSION} --build-arg VERSION=${VERSION}

run:
	podman run ${APPNAME}:${VERSION}

bundle:
	podman run --name redocly-bundle --rm -v ${PWD}/api:/spec docker.io/redocly/cli:1.14.0 bundle openapi.yaml -o raw.openapi.yaml --ext yaml

gen:
	podman run --name oas --rm -v ${PWD}:/app ${OASNAME}:v2.0 --config=api/openapi.cfg.types.yaml api/raw.openapi.yaml 
	podman run --name oas --rm -v ${PWD}:/app ${OASNAME}:v2.0 --config=api/openapi.cfg.server.yaml api/raw.openapi.yaml 