package option

type contentType string

const (
	JSON contentType = "application/json"
	XML  contentType = "application/xml"
	PDF  contentType = "application/pdf"
)

type Options struct {
	BasePath string `json:"base_path"`
	Consumes contentType
}
