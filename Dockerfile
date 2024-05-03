FROM golang:1.21

#设置工作目录
WORKDIR /app

#复制代码到容器中的/app目录
COPY . .

ENTRYPOINT ["./main"]