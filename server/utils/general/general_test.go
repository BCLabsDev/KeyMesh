package general

import "testing"

func TestCreateUID(t *testing.T) {
	a, err := GenerateUID()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(a)
}
