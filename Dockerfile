# Start from the alpine golang base imagesss
FROM docker.io/library/golang:alpine
ARG token
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY src .

# Build the Go app
RUN go build -o main .

RUN echo $token > token.txt

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"]
