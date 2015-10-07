winston = require 'winston'
fs = require 'fs'

class Logger
  constructor: () ->
    if !fs.existsSync 'logs'
      fs.mkdirSync 'logs'
    @logger = new winston.Logger
      transports: [
        new winston.transports.Console,
        new winston.transports.File
          filename: "logs/#{process.env.NODE_ENV}.log"
      ]

  info: (message) ->
    @logger.info(message)

  warn: (message) ->
    @logger.warn(message)

  error: (message) ->
    @logger.error(message)

module.exports = Logger
