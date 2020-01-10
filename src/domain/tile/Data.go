package tile

type Data struct {
	Value_ int  `json:"value"`
	Open_  bool `json:"open"`
	Flag_  bool `json:"flag"`
}

func (d *Data) Value() int {
	return d.Value_
}

func (d *Data) Open() bool {
	return d.Open_
}

func (d *Data) Flag() bool {
	return d.Flag_
}

func (d *Data) WithOpen() Data {
	return Data{d.Value_, true, d.Flag_}
}

func (d *Data) WithClose() Data {
	return Data{d.Value_, false, d.Flag_}
}

func (d *Data) WithValue(value int) Data {
	return Data{value, true, d.Flag_}
}
