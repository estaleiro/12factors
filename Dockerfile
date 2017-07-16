FROM golang:1.7 as builder
WORKDIR /go/src/github.com/rikatz/12factors
RUN go get -d -v github.com/bradfitz/gomemcache/memcache; go get -d -v github.com/bradleypeabody/gorilla-sessions-memcache; \
    go get -d -v github.com/gorilla/sessions 
COPY ./ /go/src/github.com/rikatz/12factors 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o 12factors .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/rikatz/12factors .
CMD ["./12factors"]  
