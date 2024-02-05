// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/template": {
            "get": {
                "description": "Template function",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Template function",
                "parameters": [
                    {
                        "enum": [
                            "application/json"
                        ],
                        "type": "string",
                        "description": "content type request",
                        "name": "Content-Type",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "data",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/user/register": {
            "post": {
                "description": "Register creates a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register creates a new user",
                "parameters": [
                    {
                        "enum": [
                            "application/json"
                        ],
                        "type": "string",
                        "description": "content type request",
                        "name": "Content-Type",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Register creates a new user",
                        "name": "qs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.RegisterReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/verify": {
            "put": {
                "description": "Verify a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Verify a user",
                "parameters": [
                    {
                        "enum": [
                            "application/json"
                        ],
                        "type": "string",
                        "description": "content type request",
                        "name": "Content-Type",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dtos.RegisterReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
