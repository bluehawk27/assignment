FROM golang:latest
FROM redis:alpine

# Put the binary into the container
COPY assignmentd /assignmentd
RUN mkdir /etc/assignmentd
# RUN glide get
ADD VERSION /etc/assignmentd
COPY redis-local.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
# RUN go install


# Runs the binary when someone uses it.
ENTRYPOINT ["/assignementd", "serve"]
EXPOSE 8082
