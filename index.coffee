process.env.NODE_ENV ||= 'development'

config = require 'config'
logger = new (require('./logger'))()
redis = require('redis').createClient(config.get('redis.port'),
                                      config.get('redis.host'))
amqp = require('amqp').createConnection(config.get('amqp'))
server = require('http').Server(require('express')())
io = require('socket.io')(server)
core = new (require('./core'))(redis)
chat = new (require('./chat'))(logger, io, core, amqp)
notifications = new (require('./notifications'))(logger, io)

redis.on 'ready', () ->
  logger.info 'redis rdy'
  amqp.on 'ready', () ->
    logger.info 'amqp rdy'
    #chat.listen()
    notifications.notify()
    server.listen(config.get('express.port'))
