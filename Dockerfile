FROM golang:1.17.8-alpine3.15 AS build

ARG TARGET
ARG PORT
WORKDIR /src
COPY $TARGET /src/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /outputs/$TARGET

FROM alpine:3.15
WORKDIR /app
COPY --from=build /outputs/$TARGET .
ENTRYPOINT ["/app/$TARGET"]
EXPOSE $PORT