package gotools

import "encoding/json"

type jsonOpt struct{}

// BuildToJson 将数据转成 JSON 字串
func (opt jsonOpt) BuildToJson(data interface{}) (string, error) {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}
