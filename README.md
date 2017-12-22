# rpc-http-json-golang
RPC - HTTP Server for JSON request

## How to run

### Using Docker
Build the image:
```
docker build -t rpc-http-json-golang .
```

Run image:
```
docker run --publish 1234:1234 --name test --rm rpc-http-json-golang
```

Testing server if run:
```
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234
```
