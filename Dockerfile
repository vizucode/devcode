FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY ./binary/devcode .

CMD ["./devcode"]