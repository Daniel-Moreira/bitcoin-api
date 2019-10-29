.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/sign src/interface/auth/sign.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/login src/interface/auth/login.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/xChangeHandler src/interface/bitcoin/xChangeHandler.go
  env GOOS=linux go build -ldflags="-s -w" -o bin/report src/interface/report/handler.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

offline: build
	sam local start-api --docker-network bitcoin-api-docker_local_network --skip-pull-image

deploy: clean build
	sls deploy --stage dev
