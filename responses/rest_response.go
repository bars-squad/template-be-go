package responses

import (
	"encoding/json"
	"net/http"
)

type REST interface {
	JSON(w http.ResponseWriter, response *ResponsesImpl)
}

func JSON(w http.ResponseWriter, r *ResponsesImpl) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}
