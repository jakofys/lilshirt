VERSION="v1.0.3"

build:
	podman build . -f build/Containerfile -t lilshirt:${VERSION} --build-arg VERSION=${VERSION}

run:
	podman run lilshirt:${VERSION}

up: