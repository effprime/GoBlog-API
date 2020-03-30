### Build stage

FROM golang:alpine as builder

RUN apk add git
RUN go get -v -u github.com/gorilla/mux github.com/lib/pq
COPY ./goblog /usr/local/go/src/goblog
RUN go build -o / /usr/local/go/src/goblog/cmd/blog/main.go

### Final stage

FROM golang:alpine as production

COPY --from=builder /main /app/main
WORKDIR /app/
CMD ./main