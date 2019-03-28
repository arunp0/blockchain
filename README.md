# blockchain

Simple Block Chain implemenation in Go with Proof of Work

Go to app folder
    
```sh
    $cd server/src/app/server
```

Now we need to set our GOPATH, so run the following command:

```sh
    $export GOPATH={path to app directory}/server
```

Great, now that we’ve set that, we can setup our app for dependencies.

Run the following:

```sh
    $cd src/app
    $dep init
    $dep ensure
```

To Run the server
```sh
    $go run server.go
```

To Mine:

```R
GET http://localhost:3000/mine_block
```

```json
{"index":2,"timeStamp":"2019-03-29T02:51:53.697289+05:30","pHash":"4b1c33439ca926762d25438b9a62c1b1898de7bbf5fa9af5abbc0b778aa15c02","proof":61840}
```

To Get the chain:

```R
GET http://localhost:3000/get_chain
```

```json
    [{"index":1,"timeStamp":"2019-03-29T02:51:39.540236+05:30","pHash":"0","proof":1},{"index":2,"timeStamp":"2019-03-29T02:51:53.697289+05:30","pHash":"4b1c33439ca926762d25438b9a62c1b1898de7bbf5fa9af5abbc0b778aa15c02","proof":61840}]
```

To check if chain is valid:

```R
GET http://localhost:3000/is_valid
```
```json
"Block chain is valid"
```


Run Test cases
```sh
    $cd blockchain
    $go test
```

