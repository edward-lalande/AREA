FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go build -o message-brocker .

EXPOSE 8084

CMD ["./message-brocker"]
