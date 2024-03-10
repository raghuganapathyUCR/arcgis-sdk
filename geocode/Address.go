package geocode



// MultiLineAddress is a struct that represents a multi-line address, need to figure out how to use this on client
type SingleLineAdress struct {
	Value string `json:"singleLine"`
}

func (s SingleLineAdress) isAddress() {}
func (m MultiLineAddress) isAddress() {}

type Address interface {
	isAddress()
}
