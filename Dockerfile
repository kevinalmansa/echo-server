#################################################
# Build Executable Binary
#################################################
FROM golang:1.17.5 AS builder

ENV GOCACHE /go/.cache
WORKDIR /go/src/echo-server
COPY . .
RUN go get -d -v ./...
RUN make static

#################################################
# Build Docker Image
#################################################
FROM scratch

USER nobody:nogroup

COPY --from=builder /go/src/echo-server/echo-server /app/echo-server
COPY --from=builder /go/src/echo-server/resources/passwd /etc/passwd
COPY --from=builder /go/src/echo-server/resources/group /etc/group

ENV PATH /app:$PATH

WORKDIR /app
ENTRYPOINT ["/app/echo-server"]
CMD ["/app/echo-server"]
