FROM golang:1.21 AS build-stage
WORKDIR /app
# COPY src/go.mod src/go.sum ./
COPY src/go.mod  ./
RUN go mod download
COPY src/ /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /web
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=build-stage /web /app
COPY src/static /app/static
ENTRYPOINT ["/app/web"]