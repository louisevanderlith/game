FROM golang:1.11 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/game .
COPY conf conf

EXPOSE 8100

ENTRYPOINT [ "./game" ]