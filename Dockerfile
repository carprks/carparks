# Build box
FROM golang:1.11.1 AS build

RUN mkdir -p /home/main
WORKDIR /home/main
ENV GO111MODULE=on
ADD . /home/main
RUN go get -d ./...

# Lint and Test
RUN go get -u golang.org/x/lint/golint
RUN golint -set_exit_status
RUN go test

# Build
ARG build
ARG version
ARG serviceName
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o carparks
RUN cp carparks /

# Final
FROM alpine
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN apk add curl
RUN rm -rf /var/cache/apk/*

# Move
COPY --from=build /carparks /home/

# Set TimeZone
ENV TZ=Europe/London

# Entrypoint
WORKDIR /home
ENV _SERVICENAME=carparks
RUN echo "#!/bin/bash" > ./entrypoint.sh
RUN echo "./carparks" >> ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["sh", "./entrypoint.sh"]

# Healthcheck
HEALTHCHECK --interval=5s --timeout=2s --retries=12 CMD curl --silent --fail localhost/probe || exit 1

# Expose Port
EXPOSE 80
