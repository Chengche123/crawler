package log

import (
	"encoding/json"
	"fmt"
)

func LogJson(v interface{}) {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
