# Use an official Golang runtime as a parent image
FROM golang:1.21.1

# Set the working directory inside the container
WORKDIR /cmd

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o build_box_service ./cmd

# Expose the port the application will run on
EXPOSE 9090

# Define the command to run your application
CMD ["./build_box_service"]
