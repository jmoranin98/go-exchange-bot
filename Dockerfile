FROM go:1.16-alpine
WORKDIR /app
RUN touch members.txt

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./docker-exchange-bot
CMD ["./docker-exchange-bot"]