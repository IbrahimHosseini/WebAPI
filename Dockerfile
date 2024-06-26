# Start from the latest golang base image
FROM golang:1.22.1

# Add Maintainer Info
LABEL maintainer="Ibrahim Hosseini <sehosseini@me.com>"

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Disable Go Modules
ENV GO111MODULE=off

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]