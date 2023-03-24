package custom

import (
	"encoding/json"
	"fmt"
	"io"
)

type JSON map[string]interface{}

func (j JSON) MarshalGQL(w io.Writer) {
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		panic(err)
	}
}

func (j *JSON) UnmarshalGQL(v interface{}) error {
	m, ok := v.(JSON)
	if !ok {
		return fmt.Errorf("%T is not a map", v)
	}
	*j = m
	return nil
}
