FROM golang:alpine3.15 as builder

COPY . /go/api
WORKDIR /go/api/cmd/http_api
RUN go mod vendor
RUN go build -o main .


FROM alpine:3.15.0

COPY --from=builder /go/api/http_api /api/

EXPOSE 8001
CMD ["/api/http_api"]


