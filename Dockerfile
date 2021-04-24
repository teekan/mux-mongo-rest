FROM golang:1.15.11-alpine3.13
RUN apk add git
RUN go get -u github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
COPY code.go /go/code.go
RUN go build code.go
CMD /go/code
