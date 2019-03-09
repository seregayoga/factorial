FROM golang:1.12

WORKDIR /app
COPY . .

RUN go build -v ./cmd/factorial

ENV FACTORIAL_HOST 0.0.0.0
ENV FACTORIAL_PORT 5000

CMD ["/app/factorial"]
