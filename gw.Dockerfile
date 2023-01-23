# Build
# docker build -f gw.Dockerfile -t alex/bookservgw:1.0.0 .
# Then load image to cluster
#  kind load docker-image alex/bookservgw:1.0.0
FROM golang:latest as builder

WORKDIR /builder

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /goapp ./cmd/gw/

## Deploy
FROM alpine

WORKDIR /

COPY --from=builder /goapp /goapp

EXPOSE 9000

ENTRYPOINT ["/goapp"]