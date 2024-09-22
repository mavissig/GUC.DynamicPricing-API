package common

import "net/http"

func ParseErrToHttpStatus(s string) int {
	status := http.StatusBadRequest
	switch s {
	case "not found":
		status = http.StatusNotFound
	}

	return status
}
