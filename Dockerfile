FROM golang:1.19.3-alpine

WORKDIR /app

COPY . .

RUN go build -o backend_skyshi

EXPOSE 3030

CMD ./backend_skyshi