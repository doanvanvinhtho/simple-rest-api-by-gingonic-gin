package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	r1 := Response{
		Code:    200,
		Message: "200",
		Data:    nil,
	}
	jsonValue1, err1 := json.Marshal(r1)
	assert.Nil(t, err1)
	assert.Equal(t, `{"code":200,"message":"200"}`, string(jsonValue1))

	r2 := Response{
		Code:    404,
		Message: "",
		Data:    nil,
	}
	jsonValue2, err2 := json.Marshal(r2)
	assert.Nil(t, err2)
	assert.Equal(t, `{"code":404}`, string(jsonValue2))
}
