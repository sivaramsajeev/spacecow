FROM golang:1.16 AS builder
WORKDIR /app
COPY . ./
ENV TERM xterm-256color
RUN apt-get update -y \
        && apt-get install -y libncurses-dev \
        && export CGO_CFLAGS_ALLOW=".*" \
        && export CGO_LDFLAGS_ALLOW=".*" \ 
        && go mod download \
        && go build -o spacecow
CMD ["./spacecow"] 


