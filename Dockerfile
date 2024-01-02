FROM golang:1.21-alpine AS build_base
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd cmd/gecko-tracker && \
    go build -o /build/out/my-app .

# Start fresh from a smaller image
FROM alpine:3.19.0
RUN apk add ca-certificates
COPY --from=build_base /build/out/my-app /app/gecko-tracker
COPY cmd/gecko-tracker/assets /app/assets
WORKDIR /app
CMD ["/app/gecko-tracker"]
