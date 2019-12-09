FROM centos:latest
LABEL maintainer = "holmescheng"
ENV PROJECT_DIR ./release/
ENV BINARY_NAME linux_linux

# dragon 运行环境变量，默认为生成环境，可以设置为debug/production分别对应不同的配置文件
ENV DRAGON debug

# copy release project to docker container, then just run binary file

#通过Dockerfile 构建镜像时需要注意，对时区的修改一定要放在yum upgrade后面，否则upgrade 后，会修改时区为UTC

#update system timezone
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

#update application timezone
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY "${PROJECT_DIR}" /data/release

WORKDIR /data/release
# expose default 1130 port
EXPOSE 1130
CMD "./${BINARY_NAME}"
