FROM golang:1.21.6 as build

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 go build -o ./procsignal ./main.go

FROM alpine:3.17

COPY --from=build /go/src/app/procsignal /procsignal
ENTRYPOINT ["/procsignal"]
