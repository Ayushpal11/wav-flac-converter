# Dockerfile
FROM golang:1.23


RUN apt-get update && apt-get install -y ffmpeg


WORKDIR /app
COPY . .

RUN go build -o main .
CMD ["./main"]
