FROM golang:1.17.8-alpine3.15 AS build

ARG TARGET
ARG PORT
ENV target=${TARGET}
ENV port=${PORT}
WORKDIR /src
COPY ${target} /src/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /outputs/${target}

FROM alpine:3.15
WORKDIR /app
COPY --from=build /outputs/${target} .
ENTRYPOINT ["/app/${target}"]
EXPOSE $port