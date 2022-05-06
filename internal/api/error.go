package api

type Error struct {
	Message    string     `json:"message"`
	Extensions Extensions `json:"extensions"`
	Locations  []Location
}

func (e *Error) Error() string {
	return e.Message
}

type Extensions struct {
	Category string `json:"category"`
}

type Location struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
