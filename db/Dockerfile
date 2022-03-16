FROM golang:1.17.8-alpine3.15 AS build

WORKDIR /src
COPY ./ /src/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /outputs/db

FROM alpine:3.15
WORKDIR /app
COPY --from=build /outputs/db .
ENTRYPOINT ["/app/db"]
EXPOSE 8083