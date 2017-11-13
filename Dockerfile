FROM golang:latest
EXPOSE 8080
WORKDIR /go/src/github.com/indyfree/namefinder
ADD . /go/src/github.com/indyfree/namefinder
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go build -o /namefinder .
CMD ["/namefinder"]
