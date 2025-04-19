# Dockerfile for building and running the infrasnap-agent
FROM golang:1.18 AS builder


WORKDIR /app


COPY . .

# Fetch dependencies and tidy up go.mod
RUN go mod tidy

# Compile the Go binary
RUN go build -o infrasnap-agent main.go


FROM debian:bullseye-slim

# Set working directory in the final container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/infrasnap-agent .

# Expose the port your app listens on
EXPOSE 3000

# Set the default command to run your agent
ENTRYPOINT ["./infrasnap-agent"]
