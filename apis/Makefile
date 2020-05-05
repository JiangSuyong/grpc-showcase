# 构建文件

srcDir = $(shell pwd)
outDir = $(shell pwd)

protobuf:
	docker run --rm \
		-v $(srcDir):/protos \
		-v $(outDir):/out \
		longkai/dockerfiles:grpc-toolkit-3.11.4 \
		sh -c '\
		protoc \
		-I/protos \
		-I/usr/local/include/googleapis \
		--include_imports \
		--include_source_info \
		--go_out=plugins=grpc:/out \
		--descriptor_set_out=/out/api_descriptor.pb \
		$$(find /protos -name '*.proto') \
		'
