MICROS
======

micro-service startup based on web framework of [gin](https://github.com/gin-gonic/gin)


### Install go:

````shell
export GOROOT=/usr/local/go
export GOPATH=$HOME/your/go/code
````
for more [go install](https://golang.org/doc/install#install) help

### Entry ```GOPATH/src``` dir and clone:

````sh
git clone git@github.com:kevinchendev/micros.git
````

### Install tools:

````sh
go get -u github.com/ddollar/forego
go get -u github.com/ddollar/rerun
````

### Start:

````sh
forego start
````

![image](http://i1.piimg.com/1949/6f0aadaf6d1de79f.png)

TODO:

- [x] Add RPC service
- [x] Service start
- [ ] Add CI test
- [ ] Add Dockerfile



