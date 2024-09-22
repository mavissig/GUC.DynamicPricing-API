dir_api_main = ./cmd/api
dir_api_transport = ./internal/api/transport/http
dir_api_domain = ./internal/api/domain/

docs_out_dir = ./docs

dir_build = ./bin
file_build = $(dir_build)/api

all: build-api

docs-gen:
	~/go/go1.22.4/bin/swag init --dir $(dir_api_main),$(dir_api_transport),$(dir_api_domain) --output $(docs_out_dir)

debug-run-api: docs-gen
	go run $(dir_api_main)/main.go

build-api:
	go build -o $(file_build) $(dir_api_main)/main.go

run-api:
	$(file_build)

clean:
	rm -rf $(dir_build)
	rm -rf $(docs_out_dir)
