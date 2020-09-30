outDir = $(shell pwd)
port = 80
app = grpc

protobuf:
	cd apis && $(MAKE) protobuf outDir=$(outDir)

transcoder: protobuf
	docker run --rm \
		-v $(outDir)/api_descriptor.pb:/out/api_descriptor.pb \
		longkai/dockerfiles:grpc-toolkit \
		grpc-transcoder \
			--port $(port) \
			--service $(app) \
			--descriptor /out/api_descriptor.pb | \
	kubectl apply -f -

build: protobuf
	# local machine build, require go
	GOOS=linux GOARCH=amd64 go build -o serv ./cmd/serv

container: protobuf
	docker build -t longkai/grpc-showcase:v1.0.0 .

deploy: container transcoder
	kubectl apply -f k8s.yaml

delete:
	kubectl delete envoyfilters.networking.istio.io $(app)
	kubectl delete -f k8s.yaml
	rm -r genproto
	rm api_descriptor.pb
