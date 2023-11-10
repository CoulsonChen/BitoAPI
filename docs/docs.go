// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://CoulsonChen.github.io/"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Match/NewPerson": {
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new pet to the store",
                "parameters": [
                    {
                        "description": "Person",
                        "name": "Name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/web.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        },
        "/Match/QueryMatches/{person_id}": {
            "get": {
                "description": "Query matches for person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query matches for person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "person_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Top N Matches",
                        "name": "top_n",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        },
        "/Match/Remove/{person_id}": {
            "delete": {
                "description": "Remove person from sorted set",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove person from sorted set",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "person_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        },
        "/Match/{person_id}/{person_to_match_id}": {
            "put": {
                "description": "Match people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Match people",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "person_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Person-to-match ID",
                        "name": "person_to_match_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/web.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "constant.Gender": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-varnames": [
                "Male",
                "Female"
            ]
        },
        "model.Person": {
            "type": "object",
            "properties": {
                "dates": {
                    "type": "integer"
                },
                "gender": {
                    "$ref": "#/definitions/constant.Gender"
                },
                "height": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "web.APIError": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "API documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
