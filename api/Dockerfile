FROM golang:alpine as builder

RUN mkdir app

WORKDIR /app

COPY . .

RUN go build -o main .

#stage 2

FROM alpine

RUN mkdir build_app

COPY --from=builder /app/main /build_app/

WORKDIR /build_app

EXPOSE 3000

CMD ["./main"]