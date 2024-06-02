# Use the official Golang image as the base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY . /app

# Build the Go app
RUN go build -o streaming-key-server-manager main.go

# Command to run the executable
ENTRYPOINT [ "./streaming-key-server-manager" ]

EXPOSE 8000
