FROM golang:1.22-alpine AS api-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main ./
COPY --from=builder /app/.env ./
COPY --from=builder /app/web/build ./web/build


CMD ["/main"]
