FROM golang:1.25-alpine

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

COPY . .

CMD ["go", "run", "."]