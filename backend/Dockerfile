FROM golang:1.23-alpine AS build
WORKDIR /usr/src

COPY go.mod go.sum ./
RUN go mod download -x

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY configs ./configs

RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./out/service ./cmd/service/main.go


FROM alpine:3 AS run
WORKDIR /app

# Копируем бинарник
COPY --from=build /usr/src/out/service /app/service

# Копируем конфиг (из папки configs в сборке)
COPY --from=build /usr/src/configs/config.yaml /app/config.yaml

# Запускаем сервис с указанием пути к конфигу
CMD [ "/app/service", "-config", "/app/config.yaml" ]