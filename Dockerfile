FROM golang:alpine

RUN mkdir /app

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY ./binary/devcode .

RUN chmod +x devcode

CMD ["./devcode"]