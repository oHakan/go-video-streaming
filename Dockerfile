FROM golang:1.23-alpine

RUN apk update && apk add --no-cache ffmpeg git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./src/cmd/app ./src/cmd/app

EXPOSE 9000

CMD ["./src/cmd/app/app"]