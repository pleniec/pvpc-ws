class ConnectionAuthenticator
  constructor: (@logger, @io, @redis) ->

  authenticateIncomingConnections: (callback) ->
    @io.on 'connection', (connection) ->
      @redis.get "access_token:#{connection.handshake.query.accessToken}", (error, userId) =>
        if !userId
          @logger.info 'invalid access token, disconnecting'
          connection.emit 'exception',
            text: 'invalid access token'
          connection.disconnect()
        else
          @connection.info 'valid access token'
          connection.userId = userId
          callback(connection)

module.exports = ConnectionAuthenticator
