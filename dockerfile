#my-image
# Use a specific base image for Go
FROM golang:1.20-alpine
# Add metadata labels
LABEL maintainer.one="azraji <azraji30@example.com>"
LABEL maintainer.two="ychatoua <yousra.ch.etudiante@gmail.com>"
LABEL description="Ascii art web"
LABEL version="1.0.0"
# Set the working directory inside the container
WORKDIR /app
# Copy and download dependency files
COPY go.mod  ./
RUN go mod download
# Copy the application source code
COPY . .
WORKDIR /app/main
# Build the application binary
RUN go build -o main .
#add bash 
RUN apk update && apk add bash
COPY main/templates/ /app/main/templates/
# Expose the port the application will use
#EXPOSE 8080
# Command to run the application
CMD ["./main"]