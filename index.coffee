process.env.NODE_ENV ||= 'development'

config = require 'config'
logger = new (require('./logger'))()
redis = require('redis').createClient(config.get('redis.port'),
                                      config.get('redis.host'))
amqp = require('amqp').createConnection(config.get('amqp'))

redis.on 'ready', () ->
  logger.info 'redis rdy'
  amqp.on 'ready', () ->
    logger.info 'amqp rdy'

###
var server = require('http').Server(require('express')());
var io = require('socket.io')(server);

var amqp = require('amqp').createConnection({
  host: 'localhost',
  login: 'rabbitmq',
  password: 'QdRtHV5R9irvJDTG'
});
var redis = require('redis').createClient();

redis.on('ready', function() {
  amqp.on('ready', function() {
    console.log('chat server ready');
    io.on('connection', function(socket) {
      console.log('new connection');
      redis.get('access_token:' + socket.handshake.query.accessToken, function(error, userId) {
        if(!userId) {
          console.log('invalid access token, disconnecting');
          socket.emit('errorMessage', {text: 'invalid access token'});
          socket.disconnect();
          return;
        }

        socket.on('message', function(message) {
          console.log('new message - ' + JSON.stringify(message));
          message.userId = userId;
          amqp.publish('chat:messages', JSON.stringify(message));
          redis.smembers('chat:conversation:' + message.conversation, function(error, userIds) {
            if(userIds.indexOf(userId) != -1) {
              userIds.forEach(function(userId) {
                amqp.publish('chat:user:' + userId, JSON.stringify(message));
              });
            }
            else {
              socket.emit('errorMessage', {conversation: message.conversation,
                                           text: 'you are not authorized to access this conversation'});
            }
          });
        });

        amqp.queue('chat:user:' + userId, function(queue) {
          queue.subscribe(function(queueMessage) {
            socket.emit('message', JSON.parse(queueMessage.data.toString('utf8')));
          });
        });
      });
    });

    server.listen(8080);
  });
});

###
###

{
  "name": "pvpc-chat",
  "version": "0.0.0",
  "description": "",
  "main": "app.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "repository": {
    "type": "git",
    "url": "https://github.com/pleniec/pvpc-chat.git"
  },
  "author": "",
  "bugs": {
    "url": "https://github.com/pleniec/pvpc-chat/issues"
  },
  "dependencies": {
    "socket.io": "~1.3.6",
    "express": "~4.13.1",
    "nodemon": "~1.3.7",
    "amqp": "~0.2.4",
    "redis": "~0.12.1",
    "restler": "~3.3.0"
  }
}


###
