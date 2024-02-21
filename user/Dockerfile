FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd

FROM alpine

COPY --from=builder /app/main /main

EXPOSE 8073
EXPOSE 8083
EXPOSE 8093
EXPOSE 44971

CMD ["/main"]