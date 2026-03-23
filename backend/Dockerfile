FROM golang:alpine AS build
WORKDIR /app
ARG VERSION=dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${VERSION}" -o server ./cmd/server

FROM gcr.io/distroless/static-debian12 AS runtime
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
ENTRYPOINT ["/app/server"]
