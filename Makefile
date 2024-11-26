.PHONY: build
build:
	go build -o ./tmp/main .

.PHONY: build_docker
build_docker:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o ./tmp/main .

.PHONY: run_docker
run_docker:
	docker compose up --build

.PHONY: stop_docker
stop_docker:
	docker compose down --rmi all
