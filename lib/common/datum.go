package jubatus_client

const (
	string_values = iota
	num_values
	binary_values
)

type Datum [3][]interface{}

func (d *Datum) String() string {
	ret := ""
	ret += "[{string_values: ["
	return ret
}

func NewDatum() Datum {
	return Datum{}
}

// string

func (d *Datum) StringValues() []interface{} {
	return d[string_values]
}

func (d *Datum) AddString(key string, value string) *Datum {
	if d[string_values] == nil {
		d[string_values] = make([]interface{}, 0)
	}
	d[string_values] = append(d[string_values],
		[2]interface{}{key, value})
	return d
}

// num

func (d *Datum) NumValues() []interface{} {
	return d[num_values]
}

func (d *Datum) AddNum(key string, value float64) *Datum {
	if d[num_values] == nil {
		d[num_values] = make([]interface{}, 0)
	}
	d[num_values] = append(d[num_values],
		[2]interface{}{key, value})
	return d
}

// binary

func (d *Datum) BinaryValues() []interface{} {
	return d[binary_values]
}

func (d *Datum) AddBinary(key string, value []byte) *Datum {
	if d[binary_values] == nil {
		d[binary_values] = make([]interface{}, 0)
	}
	d[binary_values] = append(d[binary_values],
		[2]interface{}{key, value})
	return d
}
