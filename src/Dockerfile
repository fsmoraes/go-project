FROM golang:alpine AS builder

COPY main.go .

RUN go build -o /go/bin/hello

FROM scratch

COPY --from=builder /go/bin/hello /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]

