FROM golang:1.21 AS builder


WORKDIR /source

COPY main.go ./
COPY go.mod ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest AS runtime

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

RUN apk --no-cache add tzdata

WORKDIR /work

COPY --from=builder /source/main .


ENV TZ=Asia/Bangkok

USER nonroot

ENV PORT=8080

CMD ["./main"] 