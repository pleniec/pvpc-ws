Notifier = require './notifier'
Chat = require './chat'

class ComponentManager
  constructor: (@services) ->
    @componentClasses = [Notifier, Chat]

  bindComponents: (connection) ->
    components = (new componentClass(@services) for componentClass in @componentClasses)
    for component in components
      component.onConnect(connection)
    connection.on 'disconnect', () ->
      for component in components
        component.onDisconnect(connection)

module.exports = ComponentManager
