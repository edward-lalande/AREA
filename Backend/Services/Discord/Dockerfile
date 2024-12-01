FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go build -o discord .

EXPOSE 8083

CMD ["./discord"]
