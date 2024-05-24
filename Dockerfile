FROM golang:1.18 as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/http-forward

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /go/bin/http-forward /
CMD ["/http-forward"]
