process.env.NODE_ENV ||= 'development'

config = require 'config'
logger = new (require('./logger'))()
redis = require('redis').createClient(config.get('redis.port'),
                                      config.get('redis.host'))
amqp = require('amqp').createConnection(config.get('amqp'))
server = require('http').Server(require('express')())
io = require('socket.io')(server)
chat = new (require('./chat'))(logger, io, amqp)
notifier = new (require('./notifier'))(logger)
connectionAuthenticator = new (require('./connectionAuthenticator'))(logger, io, redis)

redis.on 'ready', () ->
  amqp.on 'ready', () ->
    logger.info 'server started'
    connectionAuthenticator.authenticateIncomingConnections (connection) ->
      #chat.listen()
      notifier.notify(connection)
    server.listen(config.get('express.port'))
