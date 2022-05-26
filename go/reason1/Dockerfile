FROM golang:1.18-buster AS builder
WORKDIR /src

ARG ENV

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.environment=$ENV'" -o /bin/main main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app

COPY --from=builder /bin/main ./

CMD ["./main"]
