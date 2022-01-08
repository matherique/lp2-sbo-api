FROM golang:alpine AS builder

WORKDIR /app 

COPY . .

RUN go mod download

RUN go build ./cmd/api 

FROM alpine
WORKDIR /app

COPY --from=builder /app/api .

ENTRYPOINT ["/app/api"]

