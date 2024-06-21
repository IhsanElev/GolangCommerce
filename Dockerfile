# syntax=docker/dockerfile:1

# Define the Go version to be used
ARG GO_VERSION=1.21.5
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build

# Set the working directory inside the container
WORKDIR /src

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download -x

# Copy the rest of the source code to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -o /bin/server ./cmd/api

# Create a new stage for the final image
FROM alpine:latest AS final

# Install necessary packages
RUN apk --update add ca-certificates tzdata && update-ca-certificates

# Create a non-privileged user
ARG UID=10001
RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "${UID}" appuser
USER appuser

# Copy the built executable from the build stage
COPY --from=build /bin/server /bin/server

# Copy the config.yaml file to the final image
COPY cmd/api/config.yaml /src/cmd/api/config.yaml

# Set the working directory
WORKDIR /src

# Expose the application port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/bin/server"]
