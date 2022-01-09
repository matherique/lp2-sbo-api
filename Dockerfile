FROM golang:alpine AS builder

WORKDIR /app 

COPY go.mod .

RUN go mod download

COPY . .

RUN go build ./cmd/api 

FROM alpine

WORKDIR /app

COPY --from=builder /app/api .

ENTRYPOINT ["/app/api"]

