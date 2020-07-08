package handles

import (
	"net/http"
)

func GetLevels(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
}
