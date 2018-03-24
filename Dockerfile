FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/efreddo/todolist/
WORKDIR /go/src/github.com/efreddo/todolist

#RUN go get github.com/julienschmidt/httprouter
#RUN go get github.com/lib/pq
#RUN go install github.com/efreddo/todolist/main

CMD go run /go/src/github.com/efreddo/todolist/main/router.go
# Document that the service listens on port 8080.
EXPOSE 8080
