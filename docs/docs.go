// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/books": {
            "get": {
                "description": "Get all books.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "get all books",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.BookListResponse"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new book.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "create a new book",
                "parameters": [
                    {
                        "description": "Create new book",
                        "name": "createbook",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateBook"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/schema.Book"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/books/{id}": {
            "get": {
                "description": "a book.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "get a book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Book"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "update a book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a book",
                        "name": "updatebook",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Book"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "delete a book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/token/new": {
            "post": {
                "description": "Create a new access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "create a new access token",
                "parameters": [
                    {
                        "description": "Request for token",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/schema.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get all users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.UserListResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create a new user",
                "parameters": [
                    {
                        "description": "Create new user",
                        "name": "createuser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "first_name, last_name, is_active, is_admin only",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a user",
                        "name": "updateuser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.Auth": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.Book": {
            "type": "object",
            "required": [
                "author",
                "meta",
                "status",
                "title",
                "user_id"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "maxLength": 255
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/schema.Meta"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "schema.BookListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Book"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "schema.CreateBook": {
            "type": "object",
            "required": [
                "author",
                "meta",
                "status",
                "title",
                "user_id"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "maxLength": 255
                },
                "meta": {
                    "$ref": "#/definitions/schema.Meta"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "schema.CreateUser": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 150
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 100
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 100
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 10
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 5
                }
            }
        },
        "schema.ErrorResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "schema.Meta": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 1
                }
            }
        },
        "schema.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "schema.UpdateUser": {
            "type": "object",
            "required": [
                "first_name",
                "last_name"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "maxLength": 100
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "schema.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.UserListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.User"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Fiber Go API",
	Description:      "Fiber go web framework based REST API boilerplate",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
