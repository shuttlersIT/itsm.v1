FROM golang:1.21.0-alpine AS builder

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR /itsm-mvp/backend

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./
RUN ls -la ./*

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /itsm-mvp/backend

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /itsm-mvp

COPY --from=builder /itsm-mvp .

# create and set non-root USER
RUN addgroup -g 1001 docker && \
    adduser -S -u 1001 -G docker docker

RUN chown -R docker:docker /itsm-mvp && \
    chmod 755 /itsm-mvp

USER docker:docker

EXPOSE 9193

ENTRYPOINT ["./backend"]
