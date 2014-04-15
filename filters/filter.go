package filters

import (
	"net/http"
)

type Filterable interface {
	Filter(http.ResponseWriter, *http.Request) bool
}
