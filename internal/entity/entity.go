package entity

type Iso8583 struct {
	Mti      string         `json:"mti" validate:"required"`
	Fields   map[int]string `json:"fields"`
	Request  string         `json:"-"`
	Response string         `json:"-"`
}
