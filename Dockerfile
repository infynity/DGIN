# 设置golang的alpine库
FROM hub.deepin.com/library/golang:alpine AS builder

# 设置GOPATH
ENV GOPATH=/opt/gopath
# 设置项目名称
ENV PROJECT_NAME=wangfeihu
# 设置项目地址路径
ENV PROJECT_PATH=${GOPATH}/src/${PROJECT_NAME}

# 设置用于存放生成的二进制文件的目录
ADD . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
&& apk add --update make \
&& rm -rf /var/cache/apk/* \
&& make build

# Relaese docker image
FROM alpine:latest
ENV GOPATH=/opt/gopath
ENV PROJECT_NAME=wangfeihu
ENV PROJECT_PATH=${GOPATH}/src/${PROJECT_NAME}
WORKDIR ${PROJECT_PATH}
COPY --from=builder ${PROJECT_PATH}/${PROJECT_NAME} ${PROJECT_NAME}

# 启动容器时运行的命令 这里注意使用项目名称
CMD ["./wangfeihu"]