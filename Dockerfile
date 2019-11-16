FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git \
  && go build -o app.exe \
  && go get gopkg.in/urfave/cli.v2@master \
  && go get github.com/oxequa/realize



FROM alpine

WORKDIR /app

COPY --from=build /go/app/app.exe .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app.exe

CMD ["./app.exe"]
