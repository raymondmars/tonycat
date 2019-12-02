package config

import "tonycat/utils"

var BOT_TOKEN = utils.GetEvnWithDefaultVal("BOT_TOKEN", "")
var WEBHOOK_URL = utils.GetEvnWithDefaultVal("WEBHOOK_URL", "")
var BOT_USER_NAME = utils.GetEvnWithDefaultVal("BOT_USER_NAME", "")
var TURING_APIKEY = utils.GetEvnWithDefaultVal("TURING_APIKEY", "")
var TURING_USERID = utils.GetEvnWithDefaultVal("TURING_USERID", "")
var TURING_OPEN_API = utils.GetEvnWithDefaultVal("TURING_OPEN_API", "http://openapi.tuling123.com/openapi/api/v2")
