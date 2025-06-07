# syntax=docker/dockerfile:1

##############################
# Development build with Air #
##############################
FROM golang:1.24-alpine AS dev

RUN apk add --no-cache git curl

# Install Air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

CMD ["air"]

####################
# Production build #
####################
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o mendel-server ./cmd/mendel-server/
RUN CGO_ENABLED=0 go build -o db-migratre ./cmd/db-migrate/

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app

COPY --from=builder /app/mendel-server .
COPY --from=builder /app/db-migrate .

USER appuser
EXPOSE 8080
ENTRYPOINT ["./mendel-server"]