# Build
FROM golang:alpine AS builder
RUN mkdir /src
ADD . /src
WORKDIR /src/
RUN go build -o ./cmd/ehko/main.go ehko

# Final
FROM alpine
WORKDIR /app
COPY --from=builder /src/ehko /app/
ENTRYPOINT ["./ehko"]