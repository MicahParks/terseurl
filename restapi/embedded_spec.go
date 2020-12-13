// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Terse URL",
    "version": "0.0.1"
  },
  "host": "shorten.micahparks.com",
  "basePath": "/",
  "paths": {
    "/api/alive": {
      "get": {
        "description": "For Caddy to determine if the upstream provider (this executable) is alive.",
        "operationId": "alive",
        "responses": {
          "200": {
            "description": "Service is alive."
          }
        }
      }
    },
    "/api/delete/{shortened}": {
      "delete": {
        "security": [
          {
            "JWT": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Delete the given shortened URL from the backend storage, cause the shortened URL to immediately expire.",
        "operationId": "terseDelete",
        "parameters": [
          {
            "name": "delete",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "terse": {
                  "type": "boolean",
                  "default": true
                },
                "visits": {
                  "type": "boolean",
                  "default": true
                }
              }
            }
          },
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The shortened URL was successfully deleted from the backend storage."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/dump": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseDump",
        "responses": {
          "200": {
            "schema": {
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Dump"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/dump/{shortened}": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseDumpShortened",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/Dump"
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/read/{shortened}": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseRead",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/Terse"
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/visits/{shortened}": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseVisits",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "description": "The shortened URL affected.",
              "type": "array",
              "items": {
                "$ref": "#/definitions/Visit"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/write/{operation}": {
      "post": {
        "security": [
          {
            "JWT": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseWrite",
        "parameters": [
          {
            "name": "terse",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TerseOptionalShortened"
            }
          },
          {
            "enum": [
              "create",
              "update",
              "upsert"
            ],
            "type": "string",
            "name": "operation",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "description": "The shortened URL affected.",
              "type": "string"
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/{shortened}": {
      "get": {
        "description": "Use the shortened URL. It will redirect to the full URL if it has not expired.",
        "operationId": "terseRedirect",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "302": {
            "description": "The shortened URL to visit that will redirect to the given full URL.",
            "headers": {
              "Location": {
                "type": "string",
                "description": "The full URL that the redirect leads to."
              }
            }
          },
          "404": {
            "description": "The shortened URL expired or never existed."
          }
        }
      }
    }
  },
  "definitions": {
    "Dump": {
      "type": "object",
      "properties": {
        "terse": {
          "$ref": "#/definitions/Terse"
        },
        "visits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Visit"
          }
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "JWTInfo": {
      "properties": {
        "email": {
          "type": "string"
        },
        "groups": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "MediaPreview": {
      "required": [
        "imageURL",
        "title",
        "type",
        "canonicalURL"
      ],
      "properties": {
        "audioURL": {
          "type": "string"
        },
        "canonicalURL": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "determiner": {
          "type": "string"
        },
        "imageURL": {
          "type": "string"
        },
        "locale": {
          "type": "string"
        },
        "localeAlt": {
          "type": "string"
        },
        "siteName": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "twitter": {
          "type": "object",
          "$ref": "#/definitions/Twitter"
        },
        "type": {
          "type": "string"
        },
        "videoURL": {
          "type": "string"
        }
      }
    },
    "Terse": {
      "required": [
        "originalURL",
        "shortenedURL"
      ],
      "properties": {
        "deleteAt": {
          "type": "string",
          "format": "date-time"
        },
        "mediaPreview": {
          "$ref": "#/definitions/MediaPreview"
        },
        "originalURL": {
          "type": "string"
        },
        "shortenedURL": {
          "type": "string"
        }
      }
    },
    "TerseOptionalShortened": {
      "required": [
        "originalURL"
      ],
      "properties": {
        "deleteAt": {
          "type": "string",
          "format": "date-time"
        },
        "mediaPreview": {
          "$ref": "#/definitions/MediaPreview"
        },
        "originalURL": {
          "type": "string"
        },
        "shortenedURL": {
          "type": "string"
        }
      }
    },
    "Twitter": {
      "properties": {
        "card": {
          "type": "string",
          "enum": [
            "app",
            "player",
            "summary",
            "summary_large_image"
          ]
        },
        "creator": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "imageURL": {
          "type": "string"
        },
        "site": {
          "type": "string"
        },
        "siteID": {
          "type": "string"
        },
        "streamURL": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Visit": {
      "required": [
        "accessed",
        "ip"
      ],
      "properties": {
        "accessed": {
          "type": "string",
          "format": "date-time"
        },
        "headers": {
          "type": "object",
          "additionalProperties": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        },
        "ip": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWT": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Terse URL",
    "version": "0.0.1"
  },
  "host": "shorten.micahparks.com",
  "basePath": "/",
  "paths": {
    "/api/alive": {
      "get": {
        "description": "For Caddy to determine if the upstream provider (this executable) is alive.",
        "operationId": "alive",
        "responses": {
          "200": {
            "description": "Service is alive."
          }
        }
      }
    },
    "/api/delete/{shortened}": {
      "delete": {
        "security": [
          {
            "JWT": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Delete the given shortened URL from the backend storage, cause the shortened URL to immediately expire.",
        "operationId": "terseDelete",
        "parameters": [
          {
            "name": "delete",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "terse": {
                  "type": "boolean",
                  "default": true
                },
                "visits": {
                  "type": "boolean",
                  "default": true
                }
              }
            }
          },
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The shortened URL was successfully deleted from the backend storage."
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/dump": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseDump",
        "responses": {
          "200": {
            "schema": {
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Dump"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/dump/{shortened}": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseDumpShortened",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/Dump"
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/read/{shortened}": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseRead",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/Terse"
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/visits/{shortened}": {
      "get": {
        "security": [
          {
            "JWT": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseVisits",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "description": "The shortened URL affected.",
              "type": "array",
              "items": {
                "$ref": "#/definitions/Visit"
              }
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/write/{operation}": {
      "post": {
        "security": [
          {
            "JWT": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "terseWrite",
        "parameters": [
          {
            "name": "terse",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TerseOptionalShortened"
            }
          },
          {
            "enum": [
              "create",
              "update",
              "upsert"
            ],
            "type": "string",
            "name": "operation",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "description": "The shortened URL affected.",
              "type": "string"
            }
          },
          "default": {
            "description": "Unexpected error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/{shortened}": {
      "get": {
        "description": "Use the shortened URL. It will redirect to the full URL if it has not expired.",
        "operationId": "terseRedirect",
        "parameters": [
          {
            "type": "string",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "302": {
            "description": "The shortened URL to visit that will redirect to the given full URL.",
            "headers": {
              "Location": {
                "type": "string",
                "description": "The full URL that the redirect leads to."
              }
            }
          },
          "404": {
            "description": "The shortened URL expired or never existed."
          }
        }
      }
    }
  },
  "definitions": {
    "Dump": {
      "type": "object",
      "properties": {
        "terse": {
          "$ref": "#/definitions/Terse"
        },
        "visits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Visit"
          }
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "JWTInfo": {
      "properties": {
        "email": {
          "type": "string"
        },
        "groups": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "MediaPreview": {
      "required": [
        "imageURL",
        "title",
        "type",
        "canonicalURL"
      ],
      "properties": {
        "audioURL": {
          "type": "string"
        },
        "canonicalURL": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "determiner": {
          "type": "string"
        },
        "imageURL": {
          "type": "string"
        },
        "locale": {
          "type": "string"
        },
        "localeAlt": {
          "type": "string"
        },
        "siteName": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "twitter": {
          "type": "object",
          "$ref": "#/definitions/Twitter"
        },
        "type": {
          "type": "string"
        },
        "videoURL": {
          "type": "string"
        }
      }
    },
    "Terse": {
      "required": [
        "originalURL",
        "shortenedURL"
      ],
      "properties": {
        "deleteAt": {
          "type": "string",
          "format": "date-time"
        },
        "mediaPreview": {
          "$ref": "#/definitions/MediaPreview"
        },
        "originalURL": {
          "type": "string"
        },
        "shortenedURL": {
          "type": "string"
        }
      }
    },
    "TerseOptionalShortened": {
      "required": [
        "originalURL"
      ],
      "properties": {
        "deleteAt": {
          "type": "string",
          "format": "date-time"
        },
        "mediaPreview": {
          "$ref": "#/definitions/MediaPreview"
        },
        "originalURL": {
          "type": "string"
        },
        "shortenedURL": {
          "type": "string"
        }
      }
    },
    "Twitter": {
      "properties": {
        "card": {
          "type": "string",
          "enum": [
            "app",
            "player",
            "summary",
            "summary_large_image"
          ]
        },
        "creator": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "imageURL": {
          "type": "string"
        },
        "site": {
          "type": "string"
        },
        "siteID": {
          "type": "string"
        },
        "streamURL": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Visit": {
      "required": [
        "accessed",
        "ip"
      ],
      "properties": {
        "accessed": {
          "type": "string",
          "format": "date-time"
        },
        "headers": {
          "type": "object",
          "additionalProperties": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        },
        "ip": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWT": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
