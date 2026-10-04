# Build Stage
FROM golang:1.25-alpine AS build-env

ARG MAIN_PATH

RUN apk add --no-cache bash git openssh make

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /src/main ${MAIN_PATH}

# Final Stage
FROM alpine:3.22

ENV APPUSER=appuser
ENV APPUID=1000
ENV APPGID=1000

RUN addgroup --gid ${APPGID} "${APPUSER}" \
    && adduser \
    --disabled-password \
    --gecos "" \
    --home "/app" \
    --ingroup "${APPUSER}" \
    --uid "${APPUID}" \
    "${APPUSER}"

RUN apk add --no-cache curl

WORKDIR /app

USER appuser

COPY --chown=appuser:appuser --from=build-env /src/main /app/
COPY --chown=appuser:appuser products.json /app/

ENTRYPOINT ["/app/main"]