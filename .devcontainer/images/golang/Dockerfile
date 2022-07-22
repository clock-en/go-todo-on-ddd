FROM golang:1.18.2

ENV ROOT=/go/src/app
COPY . /go/src/app
WORKDIR ${ROOT}

RUN apt update && apt install -y vim git \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v github.com/99designs/gqlgen@v0.17.12 \
    && go mod tidy

RUN echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc
