FROM golang:1.19
##-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd

##FROM alpine

##COPY --from=builder /app/main /main

EXPOSE 8071
EXPOSE 8081
EXPOSE 8091
EXPOSE 44973
EXPOSE 5672
EXPOSE 9092

CMD ["/main"]