# Build the server
FROM golang:alpine as serverbuilder

RUN apk --update upgrade

RUN rm -rf /var/cache/apk/*

RUN mkdir /build 
ADD ./cabinserver /build/
WORKDIR /build 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o cabinserver .

# Build the client
FROM node:14-alpine as clientbuilder

RUN apk --update upgrade
RUN rm -rf /var/cache/apk/*

RUN mkdir /build 
ADD ./cabinclient /build/
WORKDIR /build 

RUN yarn
RUN yarn run build


# Build the final image
FROM scratch

COPY --from=serverbuilder /build/cabinserver /app/
COPY --from=clientbuilder /build/dist/* /app/static/

WORKDIR /app

EXPOSE 8080

RUN chmod +x ./cabinserver

ENTRYPOINT ["./cabinserver"]