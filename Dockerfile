FROM qnib/alplain-golang

WORKDIR /usr/local/src/github.com/qnib/galloc
COPY main.go ./
RUN go build
