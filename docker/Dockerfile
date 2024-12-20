# Этап 1: Сборка
FROM golang:1.23.3-alpine AS temp
WORKDIR /temp

COPY . .
RUN go mod tidy
RUN go build -o grimoire ./cmd/api/

FROM alpine

WORKDIR /app
COPY --from=temp /temp/grimoire /app/grimoire
COPY --from=temp /temp/templates /app/templates
COPY --from=temp /temp/config /app/config
ENTRYPOINT ["./grimoire"]
