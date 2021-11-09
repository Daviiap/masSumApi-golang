FROM golang:1.16-alpine AS BUILD_IMAGE

WORKDIR /usr/app

COPY . /usr/app/

RUN go build -o ./build/build server.go

####################################################
FROM golang:1.16-alpine

LABEL maintainer="Daviiap"
LABEL author="Daviiap"

RUN apk add --no-cache curl && apk add --no-cache --upgrade bash && rm -rf /var/cache/apk/*

WORKDIR /usr/app

COPY --from=BUILD_IMAGE /usr/app/build/build ./
COPY --from=BUILD_IMAGE /usr/app/entrypoint.sh ./

EXPOSE 8080

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]

CMD ["start"]
