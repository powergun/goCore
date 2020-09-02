package typings

import "testing"

type Anything map[string]interface{}

func TestConvertBoolToString(t *testing.T) {
	m := make(Anything)
	m["str"] = ""
	if b, OK := m["str"].(bool); OK {
		t.Log(b)
	} else {
		t.Log("failed to convert")
	}
}
