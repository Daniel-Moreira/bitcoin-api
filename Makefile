.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/sign src/interface/authentication/sign.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

offline: build
	sam local start-api

deploy: clean build
	sls deploy --stage dev
