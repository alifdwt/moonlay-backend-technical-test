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
            "name": "Alif Dewantara",
            "url": "https://github.com/alifdwt",
            "email": "aputradewantara@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/lists": {
            "get": {
                "description": "Get all lists from database. Use query parameter to limit and offset",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Get all lists",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit number of lists by page (default 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Include sublists in response",
                        "name": "withSublists",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search list by title or description",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.List"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Create list",
                "parameters": [
                    {
                        "description": "Create list",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/listsdto.ListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.List"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            }
        },
        "/lists/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Get list by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "List id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.List"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Update list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "List id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update list",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/listsdto.ListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.List"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Delete list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "List id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.List"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            }
        },
        "/sublists": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sublists"
                ],
                "summary": "Get all sublists",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit number of sublists by page (default 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search sublist by title or description",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Sublist"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sublists"
                ],
                "summary": "Create sublist",
                "parameters": [
                    {
                        "description": "Sublist request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sublistsdto.SublistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Sublist"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            }
        },
        "/sublists/list/{listId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sublists"
                ],
                "summary": "Get sublist by list id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "List id",
                        "name": "listId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit number of sublists by page (default 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search sublist by title or description",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Sublist"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            }
        },
        "/sublists/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sublists"
                ],
                "summary": "Get sublist by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sublist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Sublist"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sublists"
                ],
                "summary": "Update sublist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sublist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update sublist",
                        "name": "sublist",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sublistsdto.SublistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Sublist"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sublists"
                ],
                "summary": "Delete sublist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sublist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Sublist"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ErrorResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "dto.SuccessResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "listsdto.ListRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "My shopping list"
                },
                "title": {
                    "type": "string",
                    "example": "Shopping List"
                }
            }
        },
        "models.List": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "My shopping list"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "sublists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Sublist"
                    }
                },
                "title": {
                    "type": "string",
                    "example": "Weekly Shopping List"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                }
            }
        },
        "models.Sublist": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "Ultramilk 1 L"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "list": {
                    "$ref": "#/definitions/models.List"
                },
                "list_id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Get milk"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                }
            }
        },
        "sublistsdto.SublistRequest": {
            "type": "object",
            "required": [
                "list_id",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Ultramilk 1 L"
                },
                "list_id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Get milk"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Moonlay Academy - Backend Test (GOLANG)",
	Description:      "To do list backend handlers",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}