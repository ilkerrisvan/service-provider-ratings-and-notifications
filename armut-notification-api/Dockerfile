FROM golang:1.20.3

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

COPY go.mod ./
RUN go mod download
COPY /. .

RUN swag init

EXPOSE 8080
CMD ["go", "run", "."]