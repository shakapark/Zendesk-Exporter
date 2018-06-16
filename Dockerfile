FROM golang:1.10 AS build
WORKDIR /go/src/Zendesk-Exporter
ADD src .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o zendesk-exporter

FROM alpine
WORKDIR /app
COPY --from=build /go/src/Zendesk-Exporter/zendesk-exporter /app/
ADD zendesk.yml /app/config/
ENTRYPOINT [ "/app/zendesk-exporter","--config.file=config/config.yml" ]