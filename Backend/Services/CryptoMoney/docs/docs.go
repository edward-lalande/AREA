// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/action": {
            "post": {
                "description": "Register the Actions received by the message brocker with all informations nedded",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CryptoMoney Services Area"
                ],
                "summary": "Register an received Actions",
                "parameters": [
                    {
                        "description": "It must contains the AreaId and the reactions type",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Actions"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response is the received data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request it contains the error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error it contains the error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/actions": {
            "get": {
                "description": "Get Actions from CryptoMoney Services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actions CryptoMoney Services"
                ],
                "summary": "Get Actions from CryptoMoney Services",
                "responses": {
                    "200": {
                        "description": "Reactions name with parameters of it as object",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Actions": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "integer"
                },
                "area_id": {
                    "type": "string"
                },
                "devise": {
                    "type": "string"
                },
                "symbole": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
