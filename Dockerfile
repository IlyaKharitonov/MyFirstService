FROM golang:1.13.8

WORKDIR /app

ADD . /app


EXPOSE 8080

CMD ["./MyFirstService"]
