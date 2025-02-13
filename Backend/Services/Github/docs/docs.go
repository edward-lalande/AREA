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
                "description": "Send the code received by the frontend to get the Github access-token of the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github OAUTH2"
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
                    "Github Area"
                ],
                "summary": "Register an received Actions",
                "parameters": [
                    {
                        "description": "It must contains the AreaId and the reactions type",
                        "name": "routes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GithubAction"
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
                "description": "send all the Actions available on the Github services as an object arrays with the names and the object needed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github Area"
                ],
                "summary": "send all the Actions",
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
        "/callback": {
            "get": {
                "description": "Send the code received by Github to the frontend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github OAUTH2"
                ],
                "summary": "Send the code received by Github to the frontend",
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
                "description": "Send the url to redirect to for the OAUTH2 Github",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github OAUTH2"
                ],
                "summary": "Send the url to redirect to for the OAUTH2 Github",
                "responses": {
                    "200": {
                        "description": "the URL to redirect to for the OAUTH2 Github",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/trigger": {
            "post": {
                "description": "Sends a trigger based on the provided Area ID",
                "tags": [
                    "Github Trigger"
                ],
                "summary": "Sends a trigger to the message broker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Area ID",
                        "name": "areaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Trigger sent successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Failed to encode the payload",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/webhook/commit_comment": {
            "post": {
                "description": "Handles incoming webhook commit comment events",
                "tags": [
                    "Github Webhook"
                ],
                "summary": "Processes GitHub commit comment events",
                "parameters": [
                    {
                        "description": "Webhook commit comment data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WebhooksCommitComment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Commit comment processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON payload",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/webhook/push": {
            "post": {
                "description": "Handles incoming webhook push events and triggers actions",
                "tags": [
                    "Github Webhook"
                ],
                "summary": "Processes GitHub push events",
                "parameters": [
                    {
                        "description": "Webhook push data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WebhookPush"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Webhook processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON payload",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/webhook/push/check": {
            "post": {
                "description": "Matches the incoming webhook push data with user-defined GitHub actions",
                "tags": [
                    "Github Webhook"
                ],
                "summary": "Checks webhook push data",
                "parameters": [
                    {
                        "description": "User GitHub Action",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GithubAction"
                        }
                    },
                    {
                        "description": "Webhook push data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WebhookPush"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Trigger processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Failed to process the trigger",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "reactions": {
                    "$ref": "#/definitions/models.Reactions"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.Commit": {
            "type": "object",
            "properties": {
                "added": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "modified": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "removed": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.GithubAction": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "integer"
                },
                "area_id": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "pusher": {
                    "type": "string"
                },
                "user_token": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.Pusher": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Reactions": {
            "type": "object",
            "properties": {
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                }
            }
        },
        "models.WebhookPush": {
            "type": "object",
            "properties": {
                "commits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Commit"
                    }
                },
                "pusher": {
                    "$ref": "#/definitions/models.Pusher"
                }
            }
        },
        "models.WebhooksCommitComment": {
            "type": "object",
            "properties": {
                "comment": {
                    "$ref": "#/definitions/models.Comment"
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
