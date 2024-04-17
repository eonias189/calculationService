gen:
	protoc -I . --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative ./proto/*.proto

build-orchestrator:
	cd backend && docker build -t eonias189/calculation-service/orchestrator -f Dockerfile.orchestrator . && cd ..

build-agent:
	cd backend && docker build -t eonias189/calculation-service/agent -f Dockerfile.agent . && cd ..

build-api:
	cd backend && docker build -t eonias189/calculation-service/api -f Dockerfile.api . && cd ..

build:
	make build-agent
	make build-orchestrator
	make build-api

.PHONY: gen