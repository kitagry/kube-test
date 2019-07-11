FROM golang:1.12 AS builder

WORKDIR /go/api

ENV GO111MODULE=on

COPY go.mod go.sum /go/api/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .


FROM alpine
COPY --from=builder /go/api/app /
ENTRYPOINT ["/app"]
