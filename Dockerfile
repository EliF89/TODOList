FROM golang

ADD . /go/src/github.com/efreddo/v1/todolist/
WORKDIR /go/src/github.com/efreddo/v1/todolist

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/lib/pq
RUN go install github.com/efreddo/v1/todolist/server

CMD go run /go/src/github.com/efreddo/v1/todolist/server/server.go

EXPOSE 8080
