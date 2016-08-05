.SILENT :
.PHONY : main clean dist package

WITH_ENV = env `cat .env 2>/dev/null | xargs`
DATE := `date '+%Y%m%d'`

NAME:=micros
ROOF:=$(NAME)
TAG:=`git describe --tags --always`
LDFLAGS:=-X $(ROOF)/settings.buildVersion=$(TAG)-$(DATE)

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

dist:
	echo "building $(NAME)"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux_amd64/$(NAME)

package:
	tar -cvJf $(NAME)-linux-amd64-$(TAG).tar.xz -C dist/linux_amd64 .

rpc-go:
	echo "generating rpc go server and client"
	go get github.com/golang/protobuf/...
	go get google.golang.org/grpc
	protoc -I ../../../ -I rpc/proto --go_out=plugins=grpc:rpc/proto rpc/proto/*.proto

rpc-python:
	echo "generating rpc python server and client"
	protoc -I ../../protos --python_out=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_python_plugin` ../../protos/route_guide.proto
