FROM golang:alpine AS build-env
WORKDIR /lcm
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /lcm/go.mod
COPY go.sum /lcm/go.sum
RUN go mod download
COPY . /lcm
RUN CGO_ENABLED=0 GOOS=linux go build -o build/lcm ./lcm


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /lcm/build/lcm /
ENTRYPOINT ["/lcm"]
CMD ["up", "--grpc-port=80"]
