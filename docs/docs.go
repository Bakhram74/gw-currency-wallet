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
        "/balance": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Handler for Getting balance from wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Balance"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/exchange/rates": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Exchange currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exchange"
                ],
                "summary": "Exchange currency",
                "parameters": [
                    {
                        "description": "FromCurrency, ToCurrency, Amount",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ExchangeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ExchangeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "name password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Create user with his wallet.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "register",
                "parameters": [
                    {
                        "description": "name password email",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/wallet/deposit": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deposit to users wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Deposit",
                "parameters": [
                    {
                        "description": "amount, currency",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.DepositResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/wallet/withdraw": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Withdraw from users wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Withdraw",
                "parameters": [
                    {
                        "description": "amount, currency",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.DepositResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Balance": {
            "type": "object",
            "properties": {
                "EUR": {
                    "type": "number"
                },
                "RUB": {
                    "type": "number"
                },
                "USD": {
                    "type": "number"
                }
            }
        },
        "entity.Currency": {
            "type": "string",
            "enum": [
                "USD",
                "RUB",
                "EUR"
            ],
            "x-enum-varnames": [
                "USD",
                "RUB",
                "EUR"
            ]
        },
        "entity.DepositResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "new_balance": {
                    "$ref": "#/definitions/entity.Balance"
                }
            }
        },
        "entity.ExchangeReq": {
            "type": "object",
            "required": [
                "amount",
                "from_currency",
                "to_currency"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "from_currency": {
                    "type": "string"
                },
                "to_currency": {
                    "type": "string"
                }
            }
        },
        "entity.ExchangeResponse": {
            "type": "object",
            "properties": {
                "exchanged_amount": {
                    "type": "number"
                },
                "message": {
                    "type": "string"
                },
                "new_balance": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                }
            }
        },
        "entity.LoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.RegisterReq": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
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
        },
        "entity.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "currency": {
                    "$ref": "#/definitions/entity.Currency"
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
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Wallet-exchanger",
	Description:      "API docs for Wallet-exchanger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
