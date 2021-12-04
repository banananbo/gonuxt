FROM amazonlinux:latest
RUN yum install -y golang \
    go mod init app \
    go get github.com/labstack/ \
    go get github.com/valyala/fasthttp
realize
