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
    "description": "The Terse URL shortener.",
    "title": "Terse URL",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "shorten.micahparks.com",
  "basePath": "/",
  "paths": {
    "/api/alive": {
      "get": {
        "description": "Any non-200 response means the service is not alive.",
        "tags": [
          "system"
        ],
        "summary": "Used by Caddy or other reverse proxy to determine if the service is alive.",
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
        "description": "If only Terse data is deleted, the API user is responsible for cleaning up its Visit data before adding new Terse data under the same shortened URL.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Delete Terse and or Visit data for the given shortened URL.",
        "operationId": "terseDelete",
        "parameters": [
          {
            "description": "Indicate if Terse and or Visit data should be deleted.",
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
            "description": "The shortened URL whose data should be deleted.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The shortened URL's data was successfully deleted from the backend storage."
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
    "/api/export": {
      "get": {
        "description": "Depending on the underlying storage and amount of data, this may take a while.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Export all Terse and Visit data from the backend.",
        "operationId": "terseExport",
        "responses": {
          "200": {
            "description": "The export was successfully retrieved.",
            "schema": {
              "description": "All of the Terse and Visit data from the backend.",
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Export"
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
    "/api/export/{shortened}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Export Terse and Visit data for a single shortened URL.",
        "operationId": "terseExportOne",
        "parameters": [
          {
            "type": "string",
            "description": "The shortened URL to get the export for.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The export was successfully retrieved.",
            "schema": {
              "description": "The Terse and Visits data for a single shortened URL.",
              "$ref": "#/definitions/Export"
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
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Read the Terse data for a single shortened URL.",
        "operationId": "terseRead",
        "parameters": [
          {
            "type": "string",
            "description": "The shortened URL to get the Terse data for.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The Terse data was successfully retrieved.",
            "schema": {
              "description": "The Terse data for a single shortened URL.",
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
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the Visit data for a single shortened URL.",
        "operationId": "terseVisits",
        "parameters": [
          {
            "type": "string",
            "description": "The shortened URL to get the Visit data for.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The Visit data was successfully retrieved.",
            "schema": {
              "description": "The visit data for a single shortened URL.",
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
        "description": "\"insert\" will fail if the shortened URL already exists. \"update\" will fail if the shortened URL does not already exist. \"upsert\" will only fail if there is a failure interacting with the underlying storage. If no shortened URL is included in the given Terse data, one ill be generated randomly and returned in the response.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Perform a write operation on Terse data for a shortened URL.",
        "operationId": "terseWrite",
        "parameters": [
          {
            "description": "The Terse data, with an optional shortened URL. If no shortened URL is given, one will be generated randomly and returned in the response.",
            "name": "terse",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TerseOptionalShortened"
            }
          },
          {
            "enum": [
              "insert",
              "update",
              "upsert"
            ],
            "type": "string",
            "description": "The write operation to perform with the Terse data.",
            "name": "operation",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The write operation was successful.",
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
        "tags": [
          "public"
        ],
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
            "description": "An HTTP response that will redirect to the shortened URL's full URL.",
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
    "Export": {
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
  "tags": [
    {
      "description": "Endpoints to perform operations on Terse data.",
      "name": "api"
    },
    {
      "description": "Endpoints that are publicly accessible.",
      "name": "public"
    },
    {
      "description": "Endpoints required by the system that are not public facing and do not affect Terse data.",
      "name": "system"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "The Terse URL shortener.",
    "title": "Terse URL",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "shorten.micahparks.com",
  "basePath": "/",
  "paths": {
    "/api/alive": {
      "get": {
        "description": "Any non-200 response means the service is not alive.",
        "tags": [
          "system"
        ],
        "summary": "Used by Caddy or other reverse proxy to determine if the service is alive.",
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
        "description": "If only Terse data is deleted, the API user is responsible for cleaning up its Visit data before adding new Terse data under the same shortened URL.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Delete Terse and or Visit data for the given shortened URL.",
        "operationId": "terseDelete",
        "parameters": [
          {
            "description": "Indicate if Terse and or Visit data should be deleted.",
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
            "description": "The shortened URL whose data should be deleted.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The shortened URL's data was successfully deleted from the backend storage."
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
    "/api/export": {
      "get": {
        "description": "Depending on the underlying storage and amount of data, this may take a while.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Export all Terse and Visit data from the backend.",
        "operationId": "terseExport",
        "responses": {
          "200": {
            "description": "The export was successfully retrieved.",
            "schema": {
              "description": "All of the Terse and Visit data from the backend.",
              "type": "object",
              "additionalProperties": {
                "$ref": "#/definitions/Export"
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
    "/api/export/{shortened}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Export Terse and Visit data for a single shortened URL.",
        "operationId": "terseExportOne",
        "parameters": [
          {
            "type": "string",
            "description": "The shortened URL to get the export for.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The export was successfully retrieved.",
            "schema": {
              "description": "The Terse and Visits data for a single shortened URL.",
              "$ref": "#/definitions/Export"
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
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Read the Terse data for a single shortened URL.",
        "operationId": "terseRead",
        "parameters": [
          {
            "type": "string",
            "description": "The shortened URL to get the Terse data for.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The Terse data was successfully retrieved.",
            "schema": {
              "description": "The Terse data for a single shortened URL.",
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
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Get the Visit data for a single shortened URL.",
        "operationId": "terseVisits",
        "parameters": [
          {
            "type": "string",
            "description": "The shortened URL to get the Visit data for.",
            "name": "shortened",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The Visit data was successfully retrieved.",
            "schema": {
              "description": "The visit data for a single shortened URL.",
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
        "description": "\"insert\" will fail if the shortened URL already exists. \"update\" will fail if the shortened URL does not already exist. \"upsert\" will only fail if there is a failure interacting with the underlying storage. If no shortened URL is included in the given Terse data, one ill be generated randomly and returned in the response.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "summary": "Perform a write operation on Terse data for a shortened URL.",
        "operationId": "terseWrite",
        "parameters": [
          {
            "description": "The Terse data, with an optional shortened URL. If no shortened URL is given, one will be generated randomly and returned in the response.",
            "name": "terse",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TerseOptionalShortened"
            }
          },
          {
            "enum": [
              "insert",
              "update",
              "upsert"
            ],
            "type": "string",
            "description": "The write operation to perform with the Terse data.",
            "name": "operation",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The write operation was successful.",
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
        "tags": [
          "public"
        ],
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
            "description": "An HTTP response that will redirect to the shortened URL's full URL.",
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
    "Export": {
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
  "tags": [
    {
      "description": "Endpoints to perform operations on Terse data.",
      "name": "api"
    },
    {
      "description": "Endpoints that are publicly accessible.",
      "name": "public"
    },
    {
      "description": "Endpoints required by the system that are not public facing and do not affect Terse data.",
      "name": "system"
    }
  ]
}`))
}
