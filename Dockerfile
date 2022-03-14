FROM golang:1.17.8-alpine3.15 AS build

ARG target
ARG port
ENV TARGET=${target}
ENV PORT=${port}
WORKDIR /src
COPY ./$TARGET /src/

RUN echo $TARGET $PORT
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /outputs/gateway

FROM alpine:3.15
WORKDIR /app
COPY --from=build /outputs/$TARGET .
ENTRYPOINT ["/app/$TARGET"]
EXPOSE $PORT