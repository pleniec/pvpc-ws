class Notifier
  constructor: (@logger, @amqp) ->

  notify: (connection) ->
    @amqp.queue "notifications:user:#{connection.userId}", (queue) =>
      queue.subscribe (queuedMessage) ->
        connection.emit 'notification', JSON.parse(queuedMessage.data.toString('utf8'))

module.exports = Notifier
