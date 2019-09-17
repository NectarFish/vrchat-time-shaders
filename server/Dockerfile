#build stage
FROM golang:1.13-alpine AS builder
RUN mkdir -p /go/src/vrchat_time_shader_server
WORKDIR /go/src/vrchat_time_shader_server
COPY . /go/src/vrchat_time_shader_server/
RUN apk add --no-cache git

ENV GO111MODULE="on"
ENV CGO_ENABLED=0

RUN go build

#final stage
FROM alpine:latest
RUN mkdir -p /home/app
RUN mkdir -p /usr/local/go/lib/time
WORKDIR /home/app
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /go/src/vrchat_time_shader_server /home/app/
ENTRYPOINT /home/app/vrchat_time_shader_server
LABEL Name=VRChatTimeShaderServer Version=0.0.1
EXPOSE 4305