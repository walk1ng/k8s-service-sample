APP?=k8s-service-sample.exe
PORT?=8081

clean:
rm -f ${APP}

build: clean
go build -o ${APP}

run: build
PORT=${PORT} ./${APP}

test:
go test -v ./...