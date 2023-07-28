FROM golang:alpine as builder

RUN mkdir app

WORKDIR /app

COPY . .

RUN go build -o main .

#stage 2

FROM alpine

RUN mkdir build_app

COPY --from=builder /app/main /build_app/

COPY --from=builder /app/static /build_app/static/

WORKDIR /build_app

EXPOSE 5000

CMD ["./main"]
