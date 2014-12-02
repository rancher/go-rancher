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

## Bash Completion

Go to the last line of ~/.bashrc
and copy the contents of bash-complete.sh and source it

If you're using vim, you could execute this command on opening ~/.bashrc, and moving cursor to the last line
```
:r ! cat $(GO-RANCHER_PATH)/bash-complete.sh 
:wq (exit after saving)
source ~/.bashrc
```

# License

Apache Software License 2.0
