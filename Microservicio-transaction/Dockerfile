FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 3001

CMD ["./main"]