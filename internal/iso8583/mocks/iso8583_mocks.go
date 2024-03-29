package iso8583_mocks

import "api-iso8583-to-JSON/internal/entity"

var (
	Iso8583Message1 = entity.Iso8583{
		Mti: "0200",
		Fields: map[int]string{
			2:  "4321123443211234",
			3:  "000000",
			4:  "000000012300",
			7:  "0304054133",
			11: "001205",
			14: "0205",
			18: "5399",
			22: "022",
			25: "00",
			35: "2312312332",
			37: "206305000014",
			41: "29110001",
			42: "1001001",
			49: "840",
		},
	}
	Iso8583Message1Bytes  = []byte("02007224448028C080001643211234432112340000000000000123000304054133001205020553990220010231231233220630500001429110001        1001001840")
	Iso8583Message1Bitmap = []int{
		0,
		1,
		1,
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		1,
		0,
		0,
		0,
		1,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}

	Iso8583Message2 = entity.Iso8583{
		Mti: "0200",
		Fields: map[int]string{
			2:  "4321123443211234",
			3:  "000000",
			4:  "000000012300",
			7:  "0304054133",
			11: "001205",
			14: "0205",
			18: "5399",
			22: "022",
			25: "00",
			35: "2312312332",
			37: "206305000014",
			41: "29110001",
			42: "1001001",
			49: "840",
			71: "0010",
			72: "0009",
			73: "220518",
		},
	}
	Iso8583Message2Bytes   = []byte("0200F224448028C0800003800000000000001643211234432112340000000000000123000304054133001205020553990220010231231233220630500001429110001        100100184000100009220518")
	Iso8583Message2Bitmap1 = []int{
		1,
		1,
		1,
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		1,
		0,
		0,
		0,
		1,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
	Iso8583Message2Bitmap2 = []int{
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		1,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}
)
