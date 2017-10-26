FROM golang:1.8-alpine
RUN apk add --no-cache git
ADD . /go/src/bot
RUN cd /go/src/bot && go get && go install bot

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/bin/bot .
ENV PORT 9000
CMD ["./bot"]
