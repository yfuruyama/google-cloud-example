Client 

```sh
# local
grpc_cli --protofiles=proto/hello.proto call localhost:8080 Hello.Echo 'msg: "Hello"'

# remote
docker run -it --entrypoint=/client gcr.io/${PROJECT_ID}/cloudrun_grpc_unary -server_addr="${HOST}:443"
```
