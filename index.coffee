config = new (require('./config'))()
services =
  logger: new (require('./logger'))()
  redis: require('redis').createClient(config.redis.port, config.redis.host)
  amqp: require('amqp').createConnection(config.amqp)
server = require('http').Server(require('express')())
io = require('socket.io')(server)
connectionAuthenticator = new (require('./connectionAuthenticator'))(services.logger, io, services.redis)
componentManager = new (require('./componentManager'))(services)

services.redis.on 'ready', () ->
  services.amqp.on 'ready', () ->
    services.logger.info 'server started'
    connectionAuthenticator.authenticateIncomingConnections (connection) ->
      componentManager.bindComponents(connection)
    server.listen(config.express.port)
