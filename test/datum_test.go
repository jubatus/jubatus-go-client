package jubatus_client

import (
	"testing"

	common "../lib/common"
)

func check_empty(t *testing.T, i []interface{}) {
	if len(i) != 0 {
		t.Errorf("invalid length")
	}
}

func TestAddNum(t *testing.T) {
	d := common.NewDatum()
	for i := 0; i < 10; i++ {
		d.AddNum("a", 12)
		if len(d.NumValues()) != (i + 1) {
			t.Errorf("invalid length of num values")
		}
		check_empty(t, d.StringValues())
		check_empty(t, d.BinaryValues())
	}
}

func TestAddString(t *testing.T) {
	d := common.NewDatum()
	for i := 0; i < 10; i++ {
		d.AddString("a", "foo")
		if len(d.StringValues()) != (i + 1) {
			t.Errorf("invalid length of string values")
		}
		check_empty(t, d.NumValues())
		check_empty(t, d.BinaryValues())
	}
}

func TestAddBinary(t *testing.T) {
	d := common.NewDatum()
	for i := 0; i < 10; i++ {
		d.AddBinary("a", []byte("bar"))
		if len(d.BinaryValues()) != (i + 1) {
			t.Errorf("invalid length of binary values")
		}
		check_empty(t, d.NumValues())
		check_empty(t, d.StringValues())
	}
}
