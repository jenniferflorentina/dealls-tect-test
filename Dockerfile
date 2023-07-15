# Stage build: will build the go into single binary
FROM golang:1.17-buster AS build

WORKDIR /go/src/app
COPY . .

RUN make build-production

# Stage final: the final runner of the image
FROM gcr.io/distroless/base-debian10 AS final

COPY --from=build /go/src/app/server /
COPY --from=build /go/src/app/.env /

CMD ["/server"]
