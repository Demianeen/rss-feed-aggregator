package utils

import (
	"net/url"
	"strconv"
)

// Generic function to get query parameter with a default value and convert it to an integer
func GetQueryParamInt(v url.Values, paramName string, defaultValue int) (int, error) {
	paramStr := v.Get(paramName)
	if paramStr == "" {
		return defaultValue, nil
	}

	paramInt, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
