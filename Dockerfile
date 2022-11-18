FROM golang

WORKDIR /app

COPY . .

RUN go build -o backend_skyshi

EXPOSE 3030

CMD ./backend_skyshi