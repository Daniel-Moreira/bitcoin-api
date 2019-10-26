.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/sign src/interface/auth/sign.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/login src/interface/auth/login.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

offline: build
  npm i
	sls offline

deploy: clean build
	sls deploy --stage dev
