# Go Bindings for Rancher API

# Building

```sh
cd client
go build
```

# Tests

```sh
cd client
go test
```

# Run

## List Container
```sh
go run main.go --url=http://localhost:8080/v1/schema container list
```
## Apply Filters
```sh
go run main.go --url=http://localhost:8080/v1/schema container list --accountId 1a1 --imageId 1i34
```
# License

Apache Software License 2.0
