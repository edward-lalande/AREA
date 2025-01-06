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
        "/access-token": {
            "post": {
                "description": "Send the code received by the frontend to get the discord access-token of the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord OAUTH2"
                ],
                "summary": "Get",
                "responses": {
                    "200": {
                        "description": "the code to redirect to",
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
                "description": "send all the actions available on the discord services as an object arrays with the names and the object needed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord Area"
                ],
                "summary": "send all the actions",
                "responses": {
                    "200": {
                        "description": "Response is the received data",
                        "schema": {
                            "$ref": "#/definitions/models.ReactionGet"
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
        "/callback": {
            "get": {
                "description": "Send the code received by discord to the frontend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord OAUTH2"
                ],
                "summary": "Send the code received by discord to the frontend",
                "responses": {
                    "200": {
                        "description": "the code to redirect to",
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
        "/oauth": {
            "get": {
                "description": "Send the url to redirect to for the OAUTH2 discord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord OAUTH2"
                ],
                "summary": "Send the url to redirect to for the OAUTH2 discord",
                "responses": {
                    "200": {
                        "description": "the URL to redirect to for the OAUTH2 discord",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reaction": {
            "post": {
                "description": "Register the Actions received by the message brocker with all informations nedded",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord Area"
                ],
                "summary": "Register an received Actions",
                "parameters": [
                    {
                        "description": "It must contains the AreaId and the reactions type",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DiscordActionReceive"
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
        "/reactions": {
            "get": {
                "description": "send all the reactions available on the discord services as an object arrays with the names and the object needed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord Area"
                ],
                "summary": "send all the reactions",
                "responses": {
                    "200": {
                        "description": "Response is the received data",
                        "schema": {
                            "$ref": "#/definitions/models.ReactionGet"
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
        "/register": {
            "post": {
                "description": "Register the token of the user to an associated token if exists in the user database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord OAUTH2"
                ],
                "summary": "Register a token",
                "parameters": [
                    {
                        "description": "It must contains the access token of discord and the user token if exists",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OauthInformationSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response of the reactions",
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
        "/trigger": {
            "post": {
                "description": "Actions triggerd the reactions and call the trigger route",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord trigger"
                ],
                "summary": "Trigger an Area",
                "parameters": [
                    {
                        "description": "It contains the Area Id to the reactions",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TriggerdModels"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response of the reactions",
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
        }
    },
    "definitions": {
        "models.DiscordActionReceive": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "integer"
                },
                "area_id": {
                    "type": "string"
                },
                "channel_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message_id": {
                    "type": "string"
                },
                "user_token": {
                    "type": "string"
                }
            }
        },
        "models.OauthInformationSignUp": {
            "type": "object",
            "required": [
                "access_token",
                "user_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "user_token": {
                    "type": "string"
                }
            }
        },
        "models.ReactionGet": {
            "type": "object",
            "properties": {
                "channel_id": {
                    "type": "string"
                },
                "guild_id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "reaction_id": {
                    "type": "integer"
                },
                "reaction_type": {
                    "type": "integer"
                }
            }
        },
        "models.ReactionReceiveData": {
            "type": "object",
            "properties": {
                "area_id": {
                    "type": "string"
                },
                "channel_id": {
                    "type": "string"
                },
                "guild_id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "reaction_type": {
                    "type": "integer"
                },
                "user_token": {
                    "type": "string"
                }
            }
        },
        "models.TriggerdModels": {
            "type": "object",
            "properties": {
                "area_id": {
                    "type": "string"
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
