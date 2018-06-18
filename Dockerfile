FROM golang:1.10 AS build
ADD src /go/src/Zendesk-Exporter/src
WORKDIR /go/src/Zendesk-Exporter/src
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o zendesk-exporter

FROM alpine
RUN apk --no-cache add ca-certificates && update-ca-certificates
WORKDIR /app
COPY --from=build /go/src/Zendesk-Exporter/src/zendesk-exporter /app/
ADD zendesk.yml /app/config/
ENTRYPOINT [ "/app/zendesk-exporter","--config.file=config/config.yml" ]