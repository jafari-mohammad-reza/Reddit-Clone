build:
	@go build -o ./dist/api main.go
run: build
	@./dist/api
dev:
	@go /bin/reflex -s -r '\.go' -R '^vendor/.' -R '^_.*' go run main.go
docker-dev:
	@docker rm -f reddit-clone || true && docker build -t reddit-clone . && docker run --name reddit-clone -v $$(pwd):/app -p 5050:5050  reddit-clone
