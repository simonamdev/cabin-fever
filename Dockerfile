# Build the client
FROM node as clientbuilder

RUN mkdir /build

ADD ./cabinclient /build/
WORKDIR /build 

RUN yarn
RUN yarn run build

# Build the server
# We're using go embed to ensure the files will be in the same binary
FROM golang as serverbuilder

RUN mkdir /build 

ADD ./cabinserver /build/
COPY --from=clientbuilder /build/dist/* /build/static/

WORKDIR /build 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o cabinserver .

# Build the final image
FROM scratch

COPY --from=serverbuilder /build/cabinserver /app/

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./cabinserver"]