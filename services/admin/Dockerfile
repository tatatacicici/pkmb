FROM alpine:3.18
RUN apk add --no-cache ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY admin-app /usr/local/bin/admin-app
USER appuser
EXPOSE 8001
ENTRYPOINT ["/usr/local/bin/admin-app"]
