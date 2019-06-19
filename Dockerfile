FROM golang:1.12.6

ARG GO111MODULE=on

# Module builds
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

# App builds
COPY . .
RUN go install -v ./...

CMD ["kubecd"]
