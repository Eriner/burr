FROM golang:alpine as build

WORKDIR /build
RUN apk add -U upx git musl-dev gcc make
COPY . .
RUN make build

FROM scratch
COPY --from=build /build/burr /burr
WORKDIR /data
ENTRYPOINT [ "/burr" ]
