FROM golang:1.16-alpine AS build-env
COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM busybox
COPY --from=build-env /src/app /app/

WORKDIR /app
CMD [ "./app" ]
