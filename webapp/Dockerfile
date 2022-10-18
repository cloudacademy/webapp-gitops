FROM golang:alpine AS builder

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/demo .

FROM scratch

COPY --from=builder /go/bin/demo /go/bin/demo

CMD ["/go/bin/demo"]