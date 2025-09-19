FROM alpine:3.18
RUN apk add --no-cache ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY gateway-app /usr/local/bin/gateway-app
USER appuser
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/gateway-app"]

