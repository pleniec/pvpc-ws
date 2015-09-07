class Core
  constructor: (redis) ->
    @redis = redis

  authenticate: (accessToken, callback) ->
    @redis.get 'access_token:' + accessToken, (error, userId) ->
      callback(userId)


###
      redis.get('access_token:' + socket.handshake.query.accessToken, function(error, userId) {
        if(!userId) {
          console.log('invalid access token, disconnecting');
          socket.emit('errorMessage', {text: 'invalid access token'});
          socket.disconnect();
          return;
        }
###