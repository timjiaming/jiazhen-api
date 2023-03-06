package common

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"net/http"
	"strconv"
)

func NewHTTPClient() *http.Client {
	return &http.Client{}
}
func StructToStringMap(value any) (map[string]string, error) {
	dict, err := convertor.StructToMap(value)
	if err != nil {
		return nil, err
	}
	retval := make(map[string]string)
	for k, v := range dict {
		var strVal string
		switch val := v.(type) {
		case string:
			strVal = val
		case int:
			strVal = strconv.Itoa(val)
		case float64:
			strVal = strconv.FormatFloat(val, 'f', -1, 64)
		default:
			strVal = fmt.Sprintf("%v", v)
		}
		retval[k] = strVal
	}
	return retval, nil
}
