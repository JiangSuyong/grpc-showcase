# gRPC Showcase

This is a gRPC best practise repo, which contains:

- [Google API design](https://cloud.google.com/apis/design) for gRPC and HTTP REST service
- protocol buffers code gen via docker image, so you don't need to install all the toolchain
- best practise for writing a gRPC service
- multi-stage docker build with distroless image
- runs on Kubernetes and Istio
- integrates gRPC health probe protocol
- gRPC-JSON transcoder in Istio envoy sidecar, so you can call gRPC service via HTTP REST JSON requests
- cutting edge technology
- automation all the workflows

# Up and Running

You must have a working Kubernetes and Istio environment and config its ingress, simply run,

```shell script
$ make deploy
```

Then test the gRPC client(requires Go),

```shell script
$ go build -o cli ./cmd/cli
$ ./cli -addr=localhost:80 -name=World # replace with your server address
2020/05/05 13:35:48 Greeting: Hello World
```

Or just play with curl,

```shell script
$ curl -d '{"name": "World"}' "localhost/v1/say" # replace with your server address
{
 "message": "Hello World"
}
```

If you have any questions or contributions feel free to file an issue.

Have fun!