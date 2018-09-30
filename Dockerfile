# Scratch does not have root CA certificates, start with alpine image
FROM scratch

# Get the certificates
RUN apk add --update ca-certificates openssl

RUN echo "@edge http://nl.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "@community http://nl.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \

WORKDIR /
COPY build /
CMD ["./articleapp"]
