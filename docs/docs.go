// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Afrizalmy",
            "email": "afrizalmy1@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/license/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comment": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get all Comment user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Get all Comment user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Comment"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create Comment user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Create Comment user",
                "parameters": [
                    {
                        "description": "Create Comment user",
                        "name": "requestCreate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestComments"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.RequestComments"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            }
        },
        "/comment/{commentId}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get Comment user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Get Comment user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update Comment user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Update Comment user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Comment user",
                        "name": "requestUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestComments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete Comment user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Delete Comment user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/photo": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get all photo user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Get all photo user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Photo"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create photo user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Create photo user",
                "parameters": [
                    {
                        "description": "Create Photo user",
                        "name": "requestCreate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestPhoto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            }
        },
        "/photo/{photoId}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get photo user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Get photo user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update photo user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Update photo user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Photo user",
                        "name": "requestUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestPhoto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete photo user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Delete photo user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/socialmedia": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get all social media user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Get all social media user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SocialMedia"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create social media user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Create social media user",
                "parameters": [
                    {
                        "description": "Create social media user",
                        "name": "requestCreate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestSocialMedia"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.RequestSocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            }
        },
        "/socialmedia/{socialMediaId}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get social media user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Get social media user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update social media user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Update social media user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Social Media user",
                        "name": "requestUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestSocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete social media user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Delete social media user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "requestLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register user",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Register user",
                        "name": "requestRegister",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.RestErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.RestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo": {
                    "$ref": "#/definitions/models.Photo"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.RequestComments": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                }
            }
        },
        "models.RequestPhoto": {
            "type": "object",
            "required": [
                "photo_url",
                "title"
            ],
            "properties": {
                "caption": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.RequestSocialMedia": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "socialmedias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMedia"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UserRegister": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
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
    },
    "securityDefinitions": {
        "JWT": {
            "description": "description: Enter the token with the ` + "`" + `Bearer: ` + "`" + ` prefix, e.g. \"Bearer abcde12345\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "mygram-production-80cd.up.railway.app",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger MyGram API",
	Description:      "This is a sample service for managing MyGram server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
