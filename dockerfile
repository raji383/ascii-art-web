# to Download the latest stable version of golang-- alpine is a Lightweight (khfifa) version of linux with a small size  5 MB
FROM golang:alpine

WORKDIR /app

# Install the Bash shell in our image
RUN apk add  bash

COPY . .

RUN go build -o main ./cmd

LABEL maintainer.one="azraji <azraji30@example.com>"
LABEL maintainer.two="ychatoua <yousra.ch.etudiante@gmail.com>"
LABEL description="Ascii art web"
LABEL version="1.0.0"

EXPOSE 8080

CMD ["./main"]