FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
