# Use a multi-stage build for a smaller image size
FROM golang:1.21.0-alpine AS builder

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git

# Set the working directory to /itsm-mvp
WORKDIR /itsm-mvp

# Copy only the go.mod and go.sum files to leverage caching
COPY go.mod .
COPY go.sum .

RUN go mod tidy

# Copy the entire project
COPY . .

# Change the working directory to /itsm-mvp/backend
WORKDIR /itsm-mvp/backend

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /itsm-mvp/backend

# Second stage for a smaller image size
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /itsm-v1

COPY --from=builder /itsm-mvp/backend .

# create and set non-root USER
RUN addgroup -g 1001 docker && \
    adduser -S -u 1001 -G docker docker

RUN chown -R docker:docker /itsm-v1 && \
    chmod 755 /itsm-v1

USER docker:docker

EXPOSE 9193

ENTRYPOINT ["./itsm-v1"]
