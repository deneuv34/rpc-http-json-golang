FROM golang:1.9

WORKDIR /go/src/rpc-http-json-golang

COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD [ "go-wrapper", "run" ]