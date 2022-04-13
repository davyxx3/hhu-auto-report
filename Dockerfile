FROM golang:bullseye

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

# 安装必要的环境
RUN apt update -y \
    && apt install tesseract-ocr -y\
    && apt install libtesseract-dev -y

WORKDIR /go/src/app

ADD . .

RUN go build -o hhu_auto_report .

CMD  ["./hhu_auto_report"]