FROM --plateform=linux/amd64 docker.io/golang:1.22.3-alpine3.20

RUN apk -U upgrade

WORKDIR /app 

RUN go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@master

ENTRYPOINT ["oapi-codegen"]