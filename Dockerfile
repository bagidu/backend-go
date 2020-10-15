# Builder
FROM golang:alpine AS builder
RUN apk add --update git

ADD . /src

# Build 
RUN cd /src && go mod download && go build -o server ./http

# App Image
FROM alpine
RUN apk add --no-cache tzdata
COPY --from=builder /src/server /app/server
ENV TZ="Asia/Jakarta"
CMD /app/bagidu
