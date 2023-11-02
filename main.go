package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/quocbang/swagger-generator/models"
	"github.com/quocbang/swagger-generator/models/definitions"
	"github.com/quocbang/swagger-generator/models/information"
	"github.com/quocbang/swagger-generator/models/option"
	"github.com/quocbang/swagger-generator/models/paths"
	"github.com/quocbang/swagger-generator/models/response"
	"github.com/quocbang/swagger-generator/models/schemes"
	"github.com/quocbang/swagger-generator/models/tags"
)

func main() {
	swagger := models.Swagger{
		Version: "2.0",
		Info: information.Info{
			Version:     "1.1.1",
			Title:       "",
			Description: "This project RESTful API specification for Teqnological",
			Contact:     "contact@teqnological.asia",
		},
		Schemes: schemes.HTTPS,
		Options: option.Options{
			BasePath: "/api",
			Consumes: option.JSON,
		},
		Tag: tags.Tags{
			{
				Name:        "Account",
				Description: "account service",
			},
			{
				Name:        "Message",
				Description: "message service",
			},
		},
		Definitions: definitions.Definitions{},
		Response:    response.Response{},
		Path: paths.Path{
			Url: "/user/login",
			Setup: paths.PathSetup{
				Summary:     "login",
				Description: "login to server",
				Tags:        []string{"Account"},
				OperationID: "Login",
				Parameters: []paths.Parameters{
					{
						In:          "body",
						Name:        "body",
						Required:    true,
						Description: "login to server",
					},
				},
				Responses: paths.Responses{
					Default: paths.Default{
						Description: "OK",
					},
				},
			},
		},
	}

	swaggerByte, err := yaml.Marshal(swagger)
	if err != nil {
		log.Fatalf("failed to marshal swagger, error: %v", err)
	}

	err = os.WriteFile("swagger.yaml", swaggerByte, 0644)
	if err != nil {
		log.Fatalf("failed to write file, error: %v", err)
	}
}
