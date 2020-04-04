package app

import (
	"github.com/g-wilson/runtime/rpcservice"
	"github.com/xeipuuv/gojsonschema"
)

func (a *App) RPCService() *rpcservice.Service {
	return rpcservice.NewService(a.Logger).
		AddMethod("greet", a.Greet, gojsonschema.NewStringLoader(`{
			"type": "object",
			"additionalProperties": false,
			"required": [ "name" ],
			"properties": {
				"name": {
					"type": "string",
					"minLength": 1
				}
			}
		}`))
}
