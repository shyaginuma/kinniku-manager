FROM golang:1.17-alpine as dev

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update && apk --no-cache add git

WORKDIR /app
EXPOSE 8080

COPY ./go.* ./
RUN go mod download

COPY . ./

RUN go install -v ./cmd/server

CMD ["/go/bin/server"]
