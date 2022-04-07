FROM golang:bullseye

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

RUN cd /etc/apt && > sources.list \
    && echo "deb http://mirrors.aliyun.com/debian  stable main contrib non-free" >> sources.list \
    && echo "deb http://mirrors.aliyun.com/debian  stable-updates main contrib non-free" >> sources.list \
    && apt update && apt upgrade \
    && apt install tesseract-ocr -y\
    && apt install libtesseract-dev -y

WORKDIR /go/src/app

ADD . .

RUN go build -o hhu_auto_report .

CMD  ["./hhu_auto_report"]