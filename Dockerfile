
FROM golang:1.15 AS build

WORKDIR /go/src/tasko

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

# RUN mkdir ./static
# COPY ./static ./static

COPY --from=build /go/src/tasko/app .

EXPOSE 3000

CMD ["./app"]