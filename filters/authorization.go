package filters

import (
	// "fmt"
	"github.com/mantishK/gonotevanilla/views"
	"net/http"
)

type Authorize struct {
}

func (a *Authorize) Filter(writer http.ResponseWriter, req *http.Request) bool {
	authorized := true
	if !authorized {
		view := views.NewView(writer)
		result := make(map[string]interface{})
		result["error"] = "Authorization Error"
		result["response"] = "error"
		view.RenderJson(result)
	}
	return authorized
}
