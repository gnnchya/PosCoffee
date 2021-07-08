# Authorization Service

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![go-version](https://img.shields.io/badge/Go-v.1.14-blue.svg)](https://git.touchdevops.com/touch-product/homecare/ehr-service.git)
[![go-framework](https://img.shields.io/badge/Framework-gin-green.svg)](https://github.com/gin-gonic/gin)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://git.touchdevops.com/touch-product/homecare/ehr-service.git)

### Feature

- [x] Role Create
- [x] Role Update
- [x] Role Delete
- [x] Role Get Data Info
- [x] Role Get Lists
- [x] Permission Create
- [x] Permission Update
- [x] Permission Delete
- [x] Permission Get Data Info
- [x] Permission Get Lists

### System Requirements Development

- [x] Docker
- [x] Jaeger
- [x] MongoDB
- [x] GRPC

### Prototype

<p>
    <a href="https://github.com/gnnchya/PosCoffee/authorize">Touch Go Blueprint</a>
</p>

### Api Specification

URL: <a href="http://example.swagger-api-touch.com">example.swagger-api-touch.com</a>

### Pre-Require

Mockery
```
GO111MODULE=off go get github.com/vektra/mockery/.../
```

Swagger
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### Installation

```
git clone https://github.com/gnnchya/PosCoffee/authorize
cd authorization-service
go mod download
```

### Testing

unit testing command

```
  go test ./... -cover
```

integrating testing command

```
  go test ./... -tags integration
```

### Generate Mocks

generate mocks from interfaces for unit testing

```
  go generate ./...
```

### Local Development

development in local start mongodb jaeger

```
cd development
source ./local.env
docker-compose up -d
```

### Tracing With Jaeger

please see in the example code implement jaeger wrap service ```service/company/withtracer```

### Others

- Uber golang style guide [link](https://github.com/uber-go/guide)
- Practical Go: Real world advice for writing maintainable Go
  programs [link](https://dave.cheney.net/practical-go/presentations/qcon-china.html?fbclid=IwAR2_D2Y2HXVYUNiG3LctB0kF64YKzGUatcIHm_sLYwm9SEqEKWAd76G7NAU)