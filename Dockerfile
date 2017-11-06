FROM golang:latest
EXPOSE 8080
WORKDIR /app
ADD . /app
RUN go build -o main .
CMD ["/app/main"]
