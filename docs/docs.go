// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://belajarqywok.github.io/",
        "contact": {
            "name": "Qywok",
            "url": "https://belajarqywok.github.io/",
            "email": "belajarqywok@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/join": {
            "post": {
                "description": "Join Community Request",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "join"
                ],
                "summary": "Join Community Request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JoinRequest"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.JoinRequest": {
            "type": "object",
            "properties": {
                "github_username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Join Community",
	Description:      "Join Community REST API Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
