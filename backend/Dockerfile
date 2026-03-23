FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM gcr.io/distroless/static-debian12 AS runtime
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
ENTRYPOINT ["/app/server"]
