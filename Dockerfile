FROM centos:latest
LABEL maintainer = "Yang Cheng"
ENV PROJECT_DIR ./release/
ENV BINARY_NAME dragon

# copy release project to docker container, then just run binary file
WORKDIR /data/release
COPY PROJECT_DIR /data/release
CMD /data/release/"${BINARY_NAME}"
