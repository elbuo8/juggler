package utils

import (
	"encoding/json"
	"net/http"
)

// Idea is to sanitize before it even reaches the other middlewares

func ParseJSONFromReq(r *http.Request, m interface{}) error {
	return json.NewDecoder(r.Body).Decode(m)
}
