package paths

type method string

const (
	MethodGet     method = "GET"
	MethodHead    method = "HEAD"
	MethodPost    method = "POST"
	MethodPut     method = "PUT"
	MethodPatch   method = "PATCH"
	MethodDelete  method = "DELETE"
	MethodConnect method = "CONNECT"
	MethodOptions method = "OPTIONS"
	MethodTrace   method = "TRACE"
)

type Path struct {
	Url   string
	Setup PathSetup
}

type PathSetup struct {
	Summary     string       `json:"summary"`
	Description string       `json:"description"`
	Tags        []string     `json:"tags"`
	OperationID string       `json:"operationId"`
	Security    []any        `json:"security"`
	Parameters  []Parameters `json:"parameters"`
	Responses   Responses    `json:"responses"`
}

type Parameters struct {
	In          string `json:"in"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Description string `json:"descriptions"`
	Schema      Schema `json:"schema"`
}

type Schema struct {
	Ref string `json:"$ref"` // ref to definitions struct
}

type Default struct {
	Description string `json:"description"`
	Schema      Schema `json:"schema"`
}

type Responses struct {
	Default Default `json:"default"`
}
