build:
	@go build -o ./dist/api main.go
run: build
	@./dist/api
dev:
	@"\\wsl$\Ubuntu\home\helltion\Reddit-Clone\go run main.go"
docker-dev:
	@docker rm -f reddit-clone || true && docker build -t reddit-clone . && docker run --name reddit-clone -v $$(pwd):/app -p 5050:5050  reddit-clone
