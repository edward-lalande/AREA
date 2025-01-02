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
                "description": "Send the code received by the frontend to get the Dropbox access-token of the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dropbox OAUTH2"
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
        "/callback": {
            "get": {
                "description": "Send the code received by Dropbox to the frontend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dropbox OAUTH2"
                ],
                "summary": "Send the code received by Dropbox to the frontend",
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
        "/oauth2": {
            "get": {
                "description": "Send the url to redirect to for the OAUTH2 Dropbox",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dropbox OAUTH2"
                ],
                "summary": "Send the url to redirect to for the OAUTH2 Dropbox",
                "responses": {
                    "200": {
                        "description": "the URL to redirect to for the OAUTH2 Dropbox",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reaction": {
            "post": {
                "description": "Register the reactions received by the message brocker with all informations nedded",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dropbox Area"
                ],
                "summary": "Register an received Reactions",
                "parameters": [
                    {
                        "description": "It must contains the AreaId and the reactions type",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DropBoxReactions"
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
                "description": "send all the reactions available on the Dropbox services as an object arrays with the names and the object needed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dropbox Area"
                ],
                "summary": "send all the reactions",
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
                    "Dropbox trigger"
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
        "models.DropBoxReactions": {
            "type": "object",
            "properties": {
                "area_id": {
                    "type": "string"
                },
                "filepath_share": {
                    "type": "string"
                },
                "from_path": {
                    "type": "string"
                },
                "reaction_type": {
                    "type": "integer"
                },
                "to_path": {
                    "type": "string"
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
