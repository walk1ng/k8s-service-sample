PROJECT?=github.com/walk1ng/k8s-service-sample
APP?=mysvc
PORT?=8081

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
AZ_REGISTRY?=weiacr.azurecr.io
IMAGE?=${AZ_REGISTRY}/${APP}

GOOS?=linux
GOARCH?=amd64

clean: 
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t ${IMAGE}:${RELEASE} .

run: container
	docker stop $(APP) || true && docker rm $(APP) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		${IMAGE}:${RELEASE}

push: container
	docker push ${IMAGE}:${RELEASE}

test:
	go test -v ./...
