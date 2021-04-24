FROM golang:1.15.11-alpine3.13
RUN apk add git
RUN go get -u github.com/gorilla/mux
COPY code.go /go/code.go
RUN go build code.go
EXPOSE 8080
CMD /go/code.go
