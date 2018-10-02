# Scratch does not have root CA certificates, start with alpine image
FROM alpine

RUN apk add --update ca-certificates openssl
RUN apk update; apk add curl

WORKDIR /app
COPY build /app/
CMD ["/app/articleapp-api"]
