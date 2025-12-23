FROM golang:1.24-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/maphy9/url-shortener-svc

COPY go.mod go.sum ./
RUN go mod download

# COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/url-shortener-svc /go/src/github.com/maphy9/url-shortener-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/url-shortener-svc /usr/local/bin/url-shortener-svc
COPY --from=buildbase /go/src/github.com/maphy9/url-shortener-svc/config.yaml /usr/local/config/config.yaml

ENV KV_VIPER_FILE=/usr/local/config/config.yaml

RUN apk add --no-cache ca-certificates

ENTRYPOINT ["url-shortener-svc"]
CMD [ "run", "service" ]