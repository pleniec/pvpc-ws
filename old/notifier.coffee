class Notifier
  constructor: (@services) ->

  onConnect: (connection) ->
    @services.amqp.queue "notifications:user:#{connection.userId}", (queue) =>
      queue.subscribe (queuedMessage) ->
        connection.emit 'notification', JSON.parse(queuedMessage.data.toString('utf8'))

  onDisconnect: (connection) ->

module.exports = Notifier
