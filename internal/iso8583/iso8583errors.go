package iso8583

import "fmt"

const (
	InvalidIso8583Header = iso8583Error("Invalid Iso8583 Header")
	InvalidIso8583Input  = iso8583Error("Invalid Iso8583 Input")
	InvalidIso8583       = iso8583Error("Invalid Iso8583")
)

type iso8583Error string

func (e iso8583Error) Error() string {
	return fmt.Sprintf("ISO 8583 Error : %s", string(e))
}
