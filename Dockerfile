FROM golang:1.12.6

ARG GO111MODULE=on

WORKDIR /go/src/app
COPY . .
RUN go install -v ./...

CMD ["kubecd"]
