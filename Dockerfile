FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o application

CMD ["./application"]