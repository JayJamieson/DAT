FROM golang:1.18-buster AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.release=`git rev-parse --short=8 HEAD`'" -o /bin/node *.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app

COPY --from=builder /bin/node ./

CMD ["./node"]