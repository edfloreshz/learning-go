# syntax=docker/dockerfile:1

FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./*.go ./
COPY routes ./routes

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

EXPOSE 8080

CMD [ "/build" ]