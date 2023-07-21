# syntax=docker/dockerfile:1

FROM golang:1.20.5

WORKDIR /app

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/*.go ./
COPY api/routes ./routes

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

EXPOSE 8080

CMD [ "/build" ]