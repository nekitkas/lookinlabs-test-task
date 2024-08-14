FROM node:20-alpine AS web-builder

WORKDIR /app

COPY web/package.json ./

RUN npm install

COPY web/ ./

RUN npm run build

FROM golang:1.22-alpine AS api-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY --from=web-builder /app/build ./web/build

RUN go build -o main .

FROM alpine:latest

COPY --from=api-builder /app/main ./
COPY --from=api-builder /app/.env ./
COPY --from=web-builder /app/build ./web/build

EXPOSE $APP_PORT

CMD ["/main"]
