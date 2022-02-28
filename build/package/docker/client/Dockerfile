FROM golang:1.17-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/client

WORKDIR /dist
RUN cp /build/main .

# Build a small image
FROM alpine:3.15.0

LABEL maintainer="nico.rogalski@msg-david.de"
LABEL cng-hello-backend.project="CngHelloGrpc"
LABEL cng-hello-backend.version="1.0.0"

RUN echo "adding run user go to system" \
    && addgroup -S go -g 1000 \
    && adduser -S go -u 1000 -G go


COPY --from=builder /dist/main /

USER 1000

EXPOSE 8080

ENTRYPOINT ["/main"]