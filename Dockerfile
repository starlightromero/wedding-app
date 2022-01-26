FROM golang:1.17.6-alpine as build

RUN apk add ca-certificates=20211220-r0 --no-cache

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/wedding-app /usr/bin/wedding-app
COPY --from=build /app/public/views/index.html /public/views/index.html

ENTRYPOINT ["/usr/bin/wedding-app"]
