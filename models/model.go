package models

import (
	"github.com/quocbang/swagger-generator/models/definitions"
	"github.com/quocbang/swagger-generator/models/information"
	"github.com/quocbang/swagger-generator/models/option"
	"github.com/quocbang/swagger-generator/models/paths"
	"github.com/quocbang/swagger-generator/models/response"
	"github.com/quocbang/swagger-generator/models/schemes"
	"github.com/quocbang/swagger-generator/models/tags"
)

type Swagger struct {
	Version     string
	Info        information.Info
	Schemes     schemes.Schemes
	Options     option.Options
	Tag         tags.Tags
	Definitions definitions.Definitions
	Response    response.Response
	Path        paths.Path
}
