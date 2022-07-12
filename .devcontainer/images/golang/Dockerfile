FROM golang:1.18.2

ENV ROOT=/go/src/app
COPY . /go/src/app
WORKDIR ${ROOT}

RUN apt update && apt install -y vim git \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go mod tidy \
    && useradd -m -d /home/developer -s /bin/bash developer

USER developer

RUN echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc
