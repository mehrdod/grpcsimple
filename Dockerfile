FROM golang:1.16-buster

COPY ./ ./

ENV GOPATH ""

# build go app
RUN go mod download
RUN go build -o grpcsimple ./cmd/app/main.go

CMD ["./grpcsimple"]