# Этап 1: Сборка
FROM golang:1.23.3-alpine AS temp
WORKDIR /temp

ARG CI_COMMIT_REF_NAME
ENV CI_COMMIT_REF_NAME=$CI_COMMIT_REF_NAME


COPY . .
RUN go mod tidy
RUN go build -o warhoop ./cmd/api/

RUN apk add --no-cache git jq
RUN chmod +x scripts/generate_gitinfo.sh
RUN ./scripts/generate_gitinfo.sh "$CI_COMMIT_REF_NAME"


FROM alpine

WORKDIR /app
COPY --from=temp /temp/warhoop /app/warhoop
COPY --from=temp /temp/templates /app/templates
COPY --from=temp /temp/config /app/config
COPY --from=temp /temp/static /app/static
ENTRYPOINT ["./warhoop"]
