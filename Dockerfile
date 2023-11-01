# Build the server from source.
FROM golang:1.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /cloudrun

# Run the server.
FROM alpine:latest
WORKDIR /
COPY --from=build /go/bin/weaver /weaver
COPY --from=build /cloudrun /cloudrun
COPY weaver.toml .
EXPOSE 8080
ENTRYPOINT ["/weaver", "multi", "deploy", "weaver.toml"]
