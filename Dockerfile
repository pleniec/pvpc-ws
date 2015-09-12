FROM node:4.0.0

COPY . /opt/pvpc-ws
WORKDIR /opt/pvpc-ws

ENV NODE_ENV vagrant

RUN npm install

EXPOSE 8080

CMD ["./node_modules/.bin/coffee", "index.coffee"]
