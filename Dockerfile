FROM golang:1.16-alpine AS build-env
COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM busybox
COPY --from=build-env /src/app /app/
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

WORKDIR /app
CMD [ "./app" ]
