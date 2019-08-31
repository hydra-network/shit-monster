package utils

import (
	"net/http"
	"strconv"
)

func GetUrlIntParam(r *http.Request, paramName string, defaultValue int) (val int)  {
	value := r.URL.Query().Get(paramName)

	if value == "" {
		val = defaultValue;
	} else {
		i, err := strconv.ParseInt(value, 0, 16);

		if (err == nil) {
			val = int(i);
		} else {
			val = 0;
		}
	}

	return val
}