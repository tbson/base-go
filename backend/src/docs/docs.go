// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/v1/config/variable/": {
            "get": {
                "description": "Get list of variables with filtering, sorting and paging",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Get list of variables",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search string",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Order by id, key",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by data type",
                        "name": "data_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/restlistutil.ListRestfulResult-schema_Variable"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "restlistutil.ListRestfulResult-schema_Variable": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "Resulting items",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Variable"
                    }
                },
                "page_size": {
                    "description": "Number of items per page",
                    "type": "integer"
                },
                "pages": {
                    "description": "Pages",
                    "allOf": [
                        {
                            "$ref": "#/definitions/restlistutil.Pages"
                        }
                    ]
                },
                "total": {
                    "description": "Total records before applying pagination",
                    "type": "integer"
                },
                "total_pages": {
                    "description": "Total pages after pagination",
                    "type": "integer"
                }
            }
        },
        "restlistutil.Pages": {
            "type": "object",
            "properties": {
                "next": {
                    "description": "Next page",
                    "type": "integer"
                },
                "prev": {
                    "description": "Previous page",
                    "type": "integer"
                }
            }
        },
        "schema.Variable": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "data_type": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "simplepm.io.io",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Simple PM API",
	Description:      "Simple PM API document.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
