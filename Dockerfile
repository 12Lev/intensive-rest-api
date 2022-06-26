FROM scratch
COPY --from=alpine:latest /tmp /tmp
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY dist/lin app
COPY docs docs
EXPOSE 80
EXPOSE 8000
EXPOSE 443
ENTRYPOINT ["./app"]
