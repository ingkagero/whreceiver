############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk --no-cache add git
WORKDIR /src
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /whreceiver -a -ldflags '-extldflags "-static"' .
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /whreceiver /whreceiver
# Run the hello binary.
ENTRYPOINT ["/whreceiver"]

