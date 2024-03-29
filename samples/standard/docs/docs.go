// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "enum": [
                            "A",
                            "B",
                            "C"
                        ],
                        "type": "string",
                        "description": "Enum",
                        "name": "Enum",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Param",
                        "name": "Param",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "description": "Time",
                        "name": "Time",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sample.GetResponse"
                        }
                    }
                }
            }
        },
        "/create_table": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sample.PostCreateTableRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sample.PostCreateTableResponse"
                        }
                    }
                }
            }
        },
        "/create_user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sample.PostCreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sample.PostCreateUserResponse"
                        }
                    }
                }
            }
        },
        "/service/article": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "ID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.GetArticleResponse"
                        }
                    }
                }
            }
        },
        "/service/groups/groups": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.GetGroupsResponse"
                        }
                    }
                }
            }
        },
        "/service/static_page/static_page": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/static.GetStaticPageResponse"
                        }
                    }
                }
            }
        },
        "/service/user/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.GetResponse"
                        }
                    }
                }
            }
        },
        "/service/user/update_user_name": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PostUpdateUserNameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.PostUpdateUserNameResponse"
                        }
                    }
                }
            }
        },
        "/service/user/update_user_password": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PostUpdateUserPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.PostUpdateUserPasswordResponse"
                        }
                    }
                }
            }
        },
        "/service/user2/update_user_name": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user2.PostUpdateUserNameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user2.PostUpdateUserNameResponse"
                        }
                    }
                }
            }
        },
        "/service/user2/update_user_password": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user2.PostUpdateUserPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user2.PostUpdateUserPasswordResponse"
                        }
                    }
                }
            }
        },
        "/service/user2/{userID}/user_job_get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/_userID.GetUserJobGetResponse"
                        }
                    }
                }
            }
        },
        "/service/user2/{userID}/{JobID}/job": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "JobID",
                        "name": "JobID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/_JobID.PutJobRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "UserID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/_JobID.PutJobResponse"
                        }
                    }
                }
            }
        },
        "/service/user2/{user_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get user information by user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "search_request",
                        "name": "search_request",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user2.GetUserResponse"
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
                "summary": "Delete a user by user id",
                "parameters": [
                    {
                        "description": "request parameter",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user2.DeleteUserRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user2.DeleteUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "_JobID.PutJobRequest": {
            "type": "object",
            "properties": {
                "jobID": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "_JobID.PutJobResponse": {
            "type": "object",
            "properties": {
                "jobID": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "_userID.GetUserJobGetResponse": {
            "type": "object",
            "properties": {
                "jobName": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                }
            }
        },
        "common.Metadata": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "groups.Company": {
            "type": "object",
            "properties": {
                "metadata": {
                    "$ref": "#/definitions/common.Metadata"
                }
            }
        },
        "groups.Department": {
            "type": "object",
            "properties": {
                "metadata": {
                    "$ref": "#/definitions/common.Metadata"
                }
            }
        },
        "groups.GetGroupsResponse": {
            "type": "object",
            "properties": {
                "companies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.Company"
                    }
                },
                "departments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/groups.Department"
                    }
                }
            }
        },
        "sample.GetResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "sample.PostCreateTableRequest": {
            "type": "object",
            "properties": {
                "flag": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "map": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "sample.PostCreateTableResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                }
            }
        },
        "sample.PostCreateUserRequest": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sample.Role"
                    }
                }
            }
        },
        "sample.PostCreateUserResponse": {
            "type": "object",
            "properties": {
                "createdType": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "requestedAt": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "sample.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "recursionRoles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sample.Role"
                    }
                }
            }
        },
        "service.GetArticleResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "group": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "requestTime": {
                    "type": "string"
                }
            }
        },
        "static.GetStaticPageResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "user.GetResponse": {
            "type": "object"
        },
        "user.PostUpdateUserNameRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "user.PostUpdateUserNameResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "user.PostUpdateUserPasswordRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "passwordConfirm": {
                    "type": "string"
                }
            }
        },
        "user.PostUpdateUserPasswordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "user2.DeleteUserRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "user2.DeleteUserResponse": {
            "type": "object"
        },
        "user2.GetUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                },
                "searchRequest": {
                    "type": "string"
                }
            }
        },
        "user2.PostUpdateUserNameRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "user2.PostUpdateUserNameResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "user2.PostUpdateUserPasswordRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "passwordConfirm": {
                    "type": "string"
                }
            }
        },
        "user2.PostUpdateUserPasswordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "requestTime": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
