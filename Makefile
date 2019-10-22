.PHONY: build clean deploy

build:
	dep ensure -v
	go install bitcoin-api/src/customtypes
	env GOOS=linux go build -ldflags="-s -w" -o bin/sign src/interface/authentication/sign.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
