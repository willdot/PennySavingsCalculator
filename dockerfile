# Specify the base image needed for the Go application
FROM golang:1.12

# Create an /app directory within the image that will hold the application
# source files
RUN mkdir /app

# Copy everything in the root directory into the /app directory
ADD . /app

# Specify that all other commands will now come from within the /app directory
WORKDIR /app

# Go get dependancies
RUN go get -d -v ./...

# Run go build to compile the binary executable of the Go program
RUN go build -o main .

# Start command that will run the program
CMD ["/app/main"]