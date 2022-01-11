# syntax=docker/dockerfile:1

FROM alpine:latest

WORKDIR /app

COPY otelcol-linux config.yaml uuid ./

CMD ["./otelcol-linux", "--config", "config.yaml"]