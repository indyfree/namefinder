FROM golang:latest
EXPOSE 8080
WORKDIR /app
ADD . /app
RUN go get gopkg.in/mgo.v2
RUN go build -o main .
CMD ["/app/main"]