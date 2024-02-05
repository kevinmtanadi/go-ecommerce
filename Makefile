.PHONY: build
build: docs
	go build -o ./dist/app ./src && go build -o ./dist/app.exe ./src

.PHONY: start
start:
	./dist/app

.PHONY: deploy
deploy: docs
	go build -o ./dist/app ./src && go build -o ./dist/app.exe ./src
	./dist/app
	
.PHONY: docs
docs:
	command -v swag >/dev/null 2>&1 || { go install github.com/swaggo/swag/cmd/swag@v1.8.12; }
	PATH=$$PATH:"$$GOROOT/bin":"$$GOPATH/bin" swag init -d ./src -o ./docs