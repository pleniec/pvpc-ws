process.env.NODE_ENV ||= 'development'

config = require 'config'
#logger = new (require('./logger'))()
#redis = require('redis').createClient(config.get('redis.port'),
#                                      config.get('redis.host'))
#amqp = require('amqp').createConnection(config.get('amqp'))
services =
  logger: new (require('./logger'))()
  redis: require('redis').createClient(config.get('redis.port'), config.get('redis.host'))
  amqp: require('amqp').createConnection(config.get('amqp'))
server = require('http').Server(require('express')())
io = require('socket.io')(server)
connectionAuthenticator = new (require('./connectionAuthenticator'))(services.logger, io, services.redis)
componentManager = new (require('./componentManager'))(services)

services.redis.on 'ready', () ->
  services.amqp.on 'ready', () ->
    services.logger.info 'server started'
    connectionAuthenticator.authenticateIncomingConnections (connection) ->
      componentManager.bindComponents(connection)
    server.listen(config.get('express.port'))
