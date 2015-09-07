class Core
  constructor: (redis) ->
    @redis = redis

  authenticate: (accessToken, callback) ->
    @redis.get 'access_token:' + accessToken, (error, userId) ->
      callback(userId)

module.exports = Core
