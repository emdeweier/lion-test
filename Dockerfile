FROM golang:alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go mod tidy
RUN go mod vendor

RUN go build -o server

EXPOSE 1804

ENTRYPOINT ["/app/server"]