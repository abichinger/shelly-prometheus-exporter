FROM golang:1.19-alpine as build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build shelly-prometheus-exporter.go


FROM alpine:3.17
COPY --from=build /go/src/app/shelly-prometheus-exporter /shelly-prometheus-exporter

ENTRYPOINT ["/shelly-prometheus-exporter"]
