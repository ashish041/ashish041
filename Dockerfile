# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as build
WORKDIR /app/discovergy
COPY go.mod ./
COPY util ./util/
COPY *.go ./
ADD nohup.out ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o www

FROM alpine:latest
WORKDIR /app/discovergy
ADD nohup.out ./
COPY --from=build /app/discovergy/www ./
EXPOSE 3333
CMD [ "nohup","/app/discovergy/www","&" ]
