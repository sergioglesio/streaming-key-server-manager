# Use the official Golang image as the base image
FROM golang:1.22

# Copy nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Create ads directory and copy the script and ads
RUN mkdir -p /hls/ads/app
COPY ads/ /hls/ads/app

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY . .

COPY insert-ads.sh /hls/ads/app/insert-ads.sh

# Build the Go app
RUN go build -o streaming-key-server-manager main.go

# Command to run the executable
ENTRYPOINT [ "./streaming-key-server-manager" ]

EXPOSE 8000
