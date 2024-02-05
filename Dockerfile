FROM golang:1.19 as builder

# Install dependencies
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /go/src/go-backend-template
COPY . .

RUN go mod download && go mod verify

RUN make build

EXPOSE 8000

CMD ["./dist/app"]