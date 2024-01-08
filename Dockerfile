FROM golang:1.21 AS build

ENV CGO_ENABLED 0

COPY . /service

WORKDIR /service/app
RUN go build -o /bin/server

FROM scratch as final

COPY --from=build /bin/server /bin/server
COPY --from=build /service/.env /.env

EXPOSE 8080

ENTRYPOINT ["/bin/server"]
