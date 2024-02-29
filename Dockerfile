# https://hub.docker.com/_/golang
FROM golang:1.22.0 as builder

LABEL author="Andrej988"

# Copy local code to the container image.
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o mqtt-forwarder -ldflags "-X main.version=${VERSION}" 

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
WORKDIR /app
COPY --from=builder /app/mqtt-forwarder /app/mqtt-forwarder

VOLUME /app/certs

# Run the web service on container startup.
CMD ["/app/mqtt-forwarder"]