# Variables
PACKAGE=truestudent_network_login/pkg
VERSION=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')

# ldflags
LDFLAGS="-X '${PACKAGE}/version.commit=${COMMIT}' \
         -X '${PACKAGE}/version.version=${VERSION}' \
         -X '${PACKAGE}/version.buildDate=${BUILD_DATE}'"

# Target rules
build-linux-amd64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=${LDFLAGS} -o output/truestudent_network_login_amd64 cmd/main/main.go

build-linux-amd64-static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=${LDFLAGS} -o output/truestudent_network_login_amd64_static cmd/main/main.go

build-linux-arm:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -a -installsuffix cgo -ldflags=${LDFLAGS} -o output/truestudent_network_login_arm64 cmd/main/main.go

build-linux-arm-static:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -installsuffix cgo -ldflags=${LDFLAGS} -o output/truestudent_network_login_arm64_static cmd/main/main.go

all: build-linux-amd64 build-linux-amd64-static build-linux-arm build-linux-arm-static
