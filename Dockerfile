# Use an official Go image as the base
FROM golang:1.20-alpine AS builder

# Set up dependencies first to maximize caching
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code and build the binary
COPY . ./
RUN go build -o main .

# Debug step to list the files
RUN ls -l /app

# Use a lightweight base image for production
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .

# Command to run the app
CMD ["./main"]
