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
        "/actions": {
            "get": {
                "description": "Get actions from all services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area api-gateway"
                ],
                "summary": "Get actions from all services",
                "responses": {
                    "200": {
                        "description": "Actions name with parameters of it as object",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
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
        "/area": {
            "post": {
                "description": "Create a new combination of action and reaction (Area) for a users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area api-gateway"
                ],
                "summary": "Create a new actions-reactions or actions-multiple reactions",
                "parameters": [
                    {
                        "description": "Data for all actions-reactions",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PayloadItem"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response of all Services with the details of the executions",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
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
        "/discord": {
            "get": {
                "description": "Get data from discord like ping, access-token...",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord api-gateway"
                ],
                "summary": "Get Data from discord services",
                "parameters": [
                    {
                        "description": "routes you would like to access to Discord",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DiscordGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response of Discord",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Post data to discord for oauth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Discord api-gateway"
                ],
                "summary": "Post Data to discord services",
                "parameters": [
                    {
                        "description": "routes you would like to access to Discord, code of the user and token of him if already exists",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DiscordPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response of Discord",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
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
                "description": "Get reactions from all services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area api-gateway"
                ],
                "summary": "Get reactions from all services",
                "responses": {
                    "200": {
                        "description": "Reactions name with parameters of it as object",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
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
        "/services": {
            "get": {
                "description": "Get services up with name to display, routes to call it to the api-gateway, url and color to display",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area api-gateway"
                ],
                "summary": "Get all services up",
                "responses": {
                    "200": {
                        "description": "services up with name to display, routes to call it to the api-gateway, url and color to display",
                        "schema": {
                            "$ref": "#/definitions/routes.serviceList"
                        }
                    },
                    "500": {
                        "description": "Internal error",
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
        "/user": {
            "get": {
                "description": "Routes get user services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User api-gateway"
                ],
                "summary": "Get to the user services",
                "parameters": [
                    {
                        "description": "routes wanted to the user services",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UsersGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User services responses",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Routes to add a new user to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User api-gateway"
                ],
                "summary": "Post a new users to the user database without OAUTH2 or login",
                "parameters": [
                    {
                        "description": "user information to login or sign-up",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserInformation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
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
        "models.DiscordGet": {
            "type": "object",
            "properties": {
                "routes": {
                    "type": "string"
                }
            }
        },
        "models.DiscordPost": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "routes": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.PayloadItem": {
            "type": "object"
        },
        "models.UserInformation": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "routes": {
                    "type": "string"
                }
            }
        },
        "models.UsersGet": {
            "type": "object",
            "properties": {
                "routes": {
                    "type": "string"
                }
            }
        },
        "routes.serviceList": {
            "type": "object",
            "properties": {
                "call_to_api_gateway": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
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
