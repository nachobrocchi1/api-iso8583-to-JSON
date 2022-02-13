package iso8583

import (
	"fmt"
	"regexp"
	"strconv"
)

// IsoBitmap - interface for encode and decode
type IsoBitmap interface {
	Decode(inputHex []byte) ([]int, error)
	Encode(inputBitMap []int) ([]byte, error)
}

type isoBitmap struct{}

// NewEncodeDecode - constructor
func NewIso8583BitmapEncodeDecode() IsoBitmap {
	return &isoBitmap{}
}

var (
	re = regexp.MustCompile("^[0-9A-Fa-f]{16}$")
)

func (i *isoBitmap) Decode(inputHex []byte) ([]int, error) {
	if !re.Match(inputHex) {
		return nil, iso8583Error("Invalid hex")
	}

	var bitMap []int

	hexUint, _ := strconv.ParseUint(string(inputHex), 16, 64)

	for _, value := range strconv.FormatUint(hexUint, 2) {
		bit, _ := strconv.Atoi(string(value))
		bitMap = append(bitMap, bit)

	}

	for len(bitMap) < 64 {
		bitMap = append([]int{0}, bitMap...)
	}

	return bitMap, nil
}

func (i *isoBitmap) Encode(inputBitMap []int) ([]byte, error) {
	bitmapHex := make([]byte, 0)
	if len(inputBitMap) != 64 {
		return nil, iso8583Error("Invalid bitmap lenght")
	}

	for i := 0; i < 64; i = i + 4 {
		bin := fmt.Sprintf("%d%d%d%d", inputBitMap[i], inputBitMap[i+1], inputBitMap[i+2], inputBitMap[i+3])
		hex := parseBinToHex(bin)
		bitmapHex = append(bitmapHex, []byte(hex)...)
	}

	return bitmapHex, nil
}

func parseBinToHex(bin string) string {
	ui, _ := strconv.ParseUint(bin, 2, 64) // (num,base,bitsize)
	return fmt.Sprintf("%X", ui)           // %X for hex b16
}
