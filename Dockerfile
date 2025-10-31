FROM golang:1.24.5

RUN apt-get update && apt-get install -y gcc sqlite3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o GoAuthFlow main.go

EXPOSE 8080

CMD ["./GoAuthFlow"]
