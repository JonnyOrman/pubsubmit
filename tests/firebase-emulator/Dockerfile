ARG PROJECT

FROM node:16-alpine3.15

RUN apk --no-cache add openjdk16-jre bash openssl

RUN npm config set user roo
RUN npm i -g firebase-tools@11.16.0 && \
    firebase -V

WORKDIR firebase

COPY . .

CMD firebase emulators:start --project=${PROJECT} --only=pubsub