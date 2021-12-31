# Simple GRPC client
This app is a client for [this](https://github.com/Ali-Farhadnia/serverGRPC) app.

##Config
```json
{
    "grpc_config":{
        "host":"37.152.177.253",
        "port":"8083"
    }
}
```
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

### Insert one book
```sh
go run main.go -k=insert_one v=(books)
```

e.g.:

```sh
go run main.go -k=insert_one -v={name:test1,author:test1,pagecount:50,inventory:50}
```

### Insert many books
```sh
go run main.go -k=insert_many v=(book)
```

e.g.:

```sh
go run main.go -k=insert_many -v=[{name:test1,author:test1,pagecount:50,inventory:50}-{name:test2,author:test1,pagecount:50,inventory:50}-{name:test3,author:test1,pagecount:50,inventory:50}]
```

### Update one book(with book id)
```sh
go run main.go -k=update v=(book(with book id))
```

e.g.:

```sh
go run main.go -k=update -v={id:(some id),name:test1,author:test1,pagecount:50,inventory:50}
```

### Delete one book(with book id)
```sh
go run main.go -k=delete v=(id)
```

### Finde one book by id
```sh
go run main.go -k=find_by_id v=(id)
```
