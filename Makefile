.SILENT :
.PHONY : main clean dist package

WITH_ENV = env `cat .env 2>/dev/null | xargs`
DATE := `date '+%Y%m%d'`

NAME:=micros
ROOF:=$(NAME)
TAG:=`git describe --tags --always`
LDFLAGS:=-X $(ROOF)/settings.buildVersion=$(TAG)-$(DATE)

all: main dist rpc package

main: vet
	echo "building $(NAME)"
	go build -ldflags "$(LDFLAGS)"

vet:
	echo "checking ."
	go tool vet -all .

clean:
	echo "cleaning dist"
	rm -rf dist
	rm -f $(NAME)
	rm -f $(NAME)-*.tar.xz

dist: clean
	echo "building $(NAME)"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux_amd64/$(NAME)
	mkdir -p dist/darwin_amd64 && GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/darwin_amd64/$(NAME)

package:
	tar -cvJf $(NAME)-linux-amd64-$(TAG).tar.xz -C dist/linux_amd64 .

rpc:
	echo "building $@"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux_amd64/mircos-$@ $(ROOF)/rpc/grpcserver
	mkdir -p dist/darwin_amd64 && GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/darwin_amd64/mircos-$@ $(ROOF)/rpc/grpcserver
.PHONY: rpc

client:
	echo "Building $@"
	go get $(ROOF)/rpc/client
.PHONY: client

protoc-gen-go:
	echo "generating rpc go server and client code"
	go get github.com/golang/protobuf/...
	go get google.golang.org/grpc
	protoc -I ../../../ -I rpc/proto --go_out=plugins=grpc:rpc/proto rpc/proto/*.proto

protoc-gen-python:
	echo "generating rpc python server and client code"
	protoc -I ../../proto --python_out=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_python_plugin` ../../proto/*.proto
