# Build stage
FROM golang:alpine AS dev

WORKDIR /app

RUN apk add --no-cache --update ca-certificates git fish
RUN update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /entrypoint main.go

# FROM scratch
FROM alpine
COPY --from=dev /entrypoint /entrypoint
COPY --from=dev /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/entrypoint"]
