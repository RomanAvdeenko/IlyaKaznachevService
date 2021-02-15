FROM golang:1.15-alpine

ENV WorkDir=/opt/code
WORKDIR ${WorkDir}

COPY ./ ${WorkDir}

RUN apk update && apk upgrade && apk add --no-cache git
RUN go mod download
RUN go build -o ./bin/workshop ./cmd/workshop/main.go

ENTRYPOINT [ "./bin/workshop" ]