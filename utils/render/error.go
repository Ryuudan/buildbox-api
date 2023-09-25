package render

import (
	"encoding/json"
	"net/http"
	"time"
)

func Error(w http.ResponseWriter, r *http.Request, entity string, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Create a buffer to hold the JSON encoding.
	var buf []byte
	data := map[string]interface{}{
		"error": map[string]interface{}{
			"code":       code,
			"entity":     entity,
			"message":    message,
			"path":       r.URL.Path,
			"time_stamp": time.Now(),
		},
	}

	// Encode the data into the buffer.
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the buffer's content to the response writer.
	if _, err := w.Write(buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
