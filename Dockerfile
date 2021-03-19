FROM golang:alpine as serverbuilder

RUN apk --update upgrade

RUN rm -rf /var/cache/apk/*

RUN mkdir /build 
ADD ./cabinserver /build/
WORKDIR /build 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server .

FROM scratch

COPY --from=serverbuilder /build/server /app/

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./server"]