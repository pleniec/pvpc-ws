fs = require 'fs'
ejs = require 'ejs'

class Config
  constructor: () ->
    configFilePath = "config/#{process.env.NODE_ENV || 'development'}.json"
    configTemplate = fs.readFileSync(configFilePath).toString()
    configJson = JSON.parse(ejs.render(configTemplate))
    for k, v of configJson
      this[k] = v

module.exports = Config
