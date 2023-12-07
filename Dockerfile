FROM golang:1.19.0-alpine AS builder

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR /itsm-v1

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
RUN go mod vendor
RUN go mod verify

COPY . ./
RUN ls -la ./*

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /itsm-v1

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /itsm-v1

COPY --from=builder /itsm-v1 .

# create and set non-root USER
RUN addgroup -g 1001 docker && \
    adduser -S -u 1001 -G docker docker

RUN chown -R docker:docker /itsm-v1 && \
    chmod 755 /itsm-v1

USER docker:docker

EXPOSE 9193

ENTRYPOINT ["./itsm-v1"]
