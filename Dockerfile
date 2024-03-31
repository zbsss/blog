# TODO: include a separate stage to run `hugo` to build the static content
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /blog ./cmd/web/

FROM alpine:latest

COPY --from=build /blog /app/blog

CMD ["/app/blog"]

EXPOSE 8080
