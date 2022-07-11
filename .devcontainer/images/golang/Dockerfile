FROM golang:1.18.2

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

RUN apt update && apt install -y vim git \
    && echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest