# Simple GRPC client
This app is a client for [this](https://github.com/Ali-Farhadnia/serverGRPC) app.

## Usage
### help
```sh
go run main.go --help
```
### key options
```sh
go run main.go
```

or

```sh
go run main.go -k=keylist
```
###Insert one book
```sh
go run main.go -k=insert_one v=(books)
```

e.g.

```sh
go run main.go -k=insert_one -v={name:test1,author:test1,pagecount:50,inventory:50}
```
###Insert many books
```sh
go run main.go -k=insert_many v=(book)
```

e.g.

```sh
go run main.go -k=insert_many -v=[{name:test1,author:test1,pagecount:50,inventory:50}-{name:test2,author:test1,pagecount:50,inventory:50}-{name:test3,author:test1,pagecount:50,inventory:50}]
```
