FROM ubuntu:18.04

RUN sed -i 's/archive.ubuntu.com/mirror.kakao.com/g' /etc/apt/sources.list
RUN sed -i 's/kr.archive.ubuntu.com/mirror.kakao.com/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirror.kakao.com/g' /etc/apt/sources.list
RUN sed -i 's/extras.ubuntu.com/mirror.kakao.com/g' /etc/apt/sources.list

ENV DEFAULT_CODE all
ENV BIN_DIR /opt/bin
ENV PKG_DIR /tmp/pkg
ENV GOBIN=/root/go/bin

COPY pkg/* ${PKG_DIR}/


RUN apt-get update && apt-get install -y \
    $(cat ${PKG_DIR}/apt_packages.txt)

RUN python3 -m pip install --upgrade pip && \
    python3 -m pip install --upgrade -r ${PKG_DIR}/pip_requirements.txt

RUN unzip ${PKG_DIR}/protoc-3.6.1-linux-x86_64.zip -d /usr/local && \
    go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

WORKDIR ${BIN_DIR}
CMD ["python3", "build.py", "-h"]
