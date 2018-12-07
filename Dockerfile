# FROM golang:latest as builder

# WORKDIR /go/src/example
# RUN export HTTPS_PROXY=http://87.254.212.120:8080/; \
#     go get -v k8s.io/apimachinery/pkg/api/errors; \
#     go get -v k8s.io/apimachinery/pkg/apis/meta/v1; \
#     go get -v k8s.io/client-go/kubernetes; \
#     go get -v k8s.io/client-go/rest

# ADD . /go/src/example/
# RUN go build -o example;

FROM alpine:latest

#WORKDIR /app
#COPY --from=builder /go/src/example/example /app
COPY example /
CMD ["/example"]
