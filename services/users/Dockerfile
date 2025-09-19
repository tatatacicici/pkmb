FROM alpine:3.18
RUN apk add --no-cache ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# nama file binary harus sama dengan hasil build di host
COPY users-app /usr/local/bin/users-app
USER appuser
EXPOSE 8002
ENTRYPOINT ["/usr/local/bin/users-app"]

