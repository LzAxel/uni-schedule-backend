FROM golang:1.23-alpine as build

WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .
RUN go build -o /app/builded cmd/main.go

FROM alpine:latest
COPY --from=build /app/builded /app/builded
COPY ./schema /schema

EXPOSE 8080
ENTRYPOINT ["/app/builded"]
