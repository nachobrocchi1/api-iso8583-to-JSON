package iso8583

//FieldType tipos de campo para iso 8583
type FieldType int

//LengthType tipo de largo de campo iso8583
type LengthType int

//FieldTypes
const (
	//N tipo numerico
	N = 1
	//A tipo alfabetico
	A = 2
	//B tipo binario
	B = 3
	//AN tipo alfa numerico
	AN = 4
	//ANS Alfanumerico con caracteres especiales
	ANS = 5
	//Z Track data
	Z = 6
)

//LengthTypes
const (
	//FIXED Largo fijo
	FIXED = 0
	//LVAR largo (hasta 10)
	LVAR = 1
	//LLVAR Alfanumerico de largo dinamico (Hasta 100)
	LLVAR = 2
	//LLLVAR Alfanumerico de largo dinamico (Hasta 1000)
	LLLVAR = 3
)

const (
	//LVARBytes bytes que utilizar LVAR para indicar largo
	LVARBytes = 1
	//LLVARBytes bytes que utilizar LLVAR para indicar largo
	LLVARBytes = 2
	//LLLVARBytes bytes que utilizar LLLVAR para indicar largo
	LLLVARBytes = 3
)

// MTI
const (
	MTI_START_POS = 0
	MTI_LENGHT    = 4
)

const (
	BitmapLen            = 16
	FirstBitMapStartPos  = 4
	SecondBitMapStartPos = 20
	IsoDataStartPosition = 36
)

//FieldConfiguration definicion de configuracion de cada campo
type FieldConfiguration struct {
	Name       string
	FieldType  FieldType
	LengthType LengthType
	Length     int
}

var (
	fieldConfiguration = make(map[int]FieldConfiguration)
)

// GetIsoFieldsConfig obtiene la configuracion de campos ISO
func GetIsoFieldsConfig() map[int]FieldConfiguration {
	if len(fieldConfiguration) == 0 {
		initFieldConfig()
	}

	return fieldConfiguration
}

func initFieldConfig() {
	fieldConfiguration[2] = FieldConfiguration{"Primary account number (PAN)", AN, LLVAR, 19}
	fieldConfiguration[3] = FieldConfiguration{"Processing code", N, FIXED, 6}
	fieldConfiguration[4] = FieldConfiguration{"Amount, transaction", N, FIXED, 12}
	fieldConfiguration[5] = FieldConfiguration{"Amount, Settlement", N, FIXED, 12}
	fieldConfiguration[6] = FieldConfiguration{"Amount, cardholder billing", N, FIXED, 12}
	fieldConfiguration[7] = FieldConfiguration{"Transmission date & time", N, FIXED, 10}
	fieldConfiguration[8] = FieldConfiguration{"Amount, Cardholder billing fee", N, FIXED, 8}
	fieldConfiguration[9] = FieldConfiguration{"Conversion rate, Settlement", N, FIXED, 8}
	fieldConfiguration[10] = FieldConfiguration{"Conversion rate, cardholder billing", N, FIXED, 8}
	fieldConfiguration[11] = FieldConfiguration{"Systems trace audit number", N, FIXED, 6}
	fieldConfiguration[12] = FieldConfiguration{"Time, Local transaction", N, FIXED, 6}
	fieldConfiguration[13] = FieldConfiguration{"Date, Local transaction (MMdd)", N, FIXED, 4}
	fieldConfiguration[14] = FieldConfiguration{"Date, Expiration", N, FIXED, 4}
	fieldConfiguration[15] = FieldConfiguration{"Date, Settlement", N, FIXED, 4}
	fieldConfiguration[16] = FieldConfiguration{"Date, conversion", N, FIXED, 4}
	fieldConfiguration[17] = FieldConfiguration{"Date, capture", N, FIXED, 4}
	fieldConfiguration[18] = FieldConfiguration{"Merchant type", N, FIXED, 4}
	fieldConfiguration[19] = FieldConfiguration{"Acquiring institution country code", N, FIXED, 3}
	fieldConfiguration[20] = FieldConfiguration{"PAN Extended, country code", N, FIXED, 3}
	fieldConfiguration[21] = FieldConfiguration{"Forwarding institution. country code", N, FIXED, 3}
	fieldConfiguration[22] = FieldConfiguration{"Point of service entry mode", N, FIXED, 3}
	fieldConfiguration[23] = FieldConfiguration{"Application PAN number", N, FIXED, 3}
	fieldConfiguration[24] = FieldConfiguration{"Function code(ISO 8583:1993)/Network International identifier", N, FIXED, 3}
	fieldConfiguration[25] = FieldConfiguration{"Point of service condition code", N, FIXED, 2}
	fieldConfiguration[26] = FieldConfiguration{"Point of service capture code", N, FIXED, 2}
	fieldConfiguration[27] = FieldConfiguration{"Authorizing identification response length", N, FIXED, 1}
	fieldConfiguration[28] = FieldConfiguration{"Amount, transaction fee", N, FIXED, 8}
	fieldConfiguration[29] = FieldConfiguration{"Amount. settlement fee", N, FIXED, 8}
	fieldConfiguration[30] = FieldConfiguration{"Amount, transaction processing fee", N, FIXED, 8}
	fieldConfiguration[31] = FieldConfiguration{"Amount, settlement processing fee", N, FIXED, 8}
	fieldConfiguration[32] = FieldConfiguration{"Acquiring institution identification code", N, LLVAR, 11}
	fieldConfiguration[33] = FieldConfiguration{"Forwarding institution identification code", N, LLVAR, 11}
	fieldConfiguration[34] = FieldConfiguration{"Primary account number, extended", AN, LLLVAR, 167}
	fieldConfiguration[35] = FieldConfiguration{"Track 2 data", Z, LLVAR, 37}
	fieldConfiguration[36] = FieldConfiguration{"Track 3 data", N, LLLVAR, 104}
	fieldConfiguration[37] = FieldConfiguration{"Retrieval reference number", AN, FIXED, 12}
	fieldConfiguration[38] = FieldConfiguration{"Authorization identification response", AN, FIXED, 6}
	fieldConfiguration[39] = FieldConfiguration{"Response code", AN, FIXED, 2}
	fieldConfiguration[40] = FieldConfiguration{"Service restriction code", AN, FIXED, 3}
	fieldConfiguration[41] = FieldConfiguration{"Card acceptor terminal identification", ANS, FIXED, 8}
	fieldConfiguration[42] = FieldConfiguration{"Card acceptor identification code", ANS, FIXED, 15}
	fieldConfiguration[43] = FieldConfiguration{"Card acceptor name/location", ANS, FIXED, 40}
	fieldConfiguration[44] = FieldConfiguration{"Additional response data", AN, LLVAR, 25}
	fieldConfiguration[45] = FieldConfiguration{"Track 1 Data", AN, LLVAR, 76}
	fieldConfiguration[46] = FieldConfiguration{"Additional data - ISO", AN, LLLVAR, 999}
	fieldConfiguration[47] = FieldConfiguration{"Additional data - National", AN, LLLVAR, 999}
	fieldConfiguration[48] = FieldConfiguration{"Additional data - Private", ANS, LLLVAR, 999}
	fieldConfiguration[49] = FieldConfiguration{"Currency code, transaction", AN, FIXED, 3}
	fieldConfiguration[50] = FieldConfiguration{"Currency code, settlement", AN, FIXED, 3}
	fieldConfiguration[51] = FieldConfiguration{"Currency code, cardholder billing", AN, FIXED, 3}
	fieldConfiguration[52] = FieldConfiguration{"Personal Identification number data", AN, FIXED, 16}
	fieldConfiguration[53] = FieldConfiguration{"Security related control information", AN, FIXED, 16}
	fieldConfiguration[54] = FieldConfiguration{"Additional amounts", AN, LLLVAR, 120}
	fieldConfiguration[55] = FieldConfiguration{"Reserved ISO", ANS, LLLVAR, 999}
	fieldConfiguration[56] = FieldConfiguration{"Reserved ISO", ANS, LLLVAR, 999}
	fieldConfiguration[57] = FieldConfiguration{"Reserved National", ANS, LLLVAR, 999}
	fieldConfiguration[58] = FieldConfiguration{"Reserved National", ANS, LLLVAR, 999}
	fieldConfiguration[59] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[60] = FieldConfiguration{"Advice/reason code (private reserved)", AN, LLVAR, 7} //modificado
	fieldConfiguration[61] = FieldConfiguration{"Reserved Private", ANS, LLLVAR, 999}
	fieldConfiguration[62] = FieldConfiguration{"Reserved Private", ANS, LLLVAR, 999}
	fieldConfiguration[63] = FieldConfiguration{"Reserved Private", ANS, LLLVAR, 999}
	fieldConfiguration[64] = FieldConfiguration{"Message authentication code (MAC)", B, FIXED, 16}
	fieldConfiguration[65] = FieldConfiguration{"Bit map, tertiary", B, FIXED, 16}
	fieldConfiguration[66] = FieldConfiguration{"Settlement code", N, FIXED, 1}
	fieldConfiguration[67] = FieldConfiguration{"Extended payment code", N, FIXED, 2}
	fieldConfiguration[68] = FieldConfiguration{"Receiving institution country code", N, FIXED, 3}
	fieldConfiguration[69] = FieldConfiguration{"Settlement institution county code", N, FIXED, 3}
	fieldConfiguration[70] = FieldConfiguration{"Network management Information code", N, FIXED, 3}
	fieldConfiguration[71] = FieldConfiguration{"Message number", N, FIXED, 4}
	fieldConfiguration[72] = FieldConfiguration{"Data record (ISO 8583:1993)/n 4 Message number, last", ANS, LLLVAR, 999}
	fieldConfiguration[73] = FieldConfiguration{"Action date (YYMMDD)", N, FIXED, 6}
	fieldConfiguration[74] = FieldConfiguration{"Number of credits", N, FIXED, 10}
	fieldConfiguration[75] = FieldConfiguration{"Credits, reversal number", N, FIXED, 10}
	fieldConfiguration[76] = FieldConfiguration{"Number of debits", N, FIXED, 10}
	fieldConfiguration[77] = FieldConfiguration{"Debits, reversal number", N, FIXED, 10}
	fieldConfiguration[78] = FieldConfiguration{"Transfer number", N, FIXED, 10}
	fieldConfiguration[79] = FieldConfiguration{"Transfer, reversal number", N, FIXED, 10}
	fieldConfiguration[80] = FieldConfiguration{"Number of inquiries", N, FIXED, 10}
	fieldConfiguration[81] = FieldConfiguration{"Number of authorizations", N, FIXED, 10}
	fieldConfiguration[82] = FieldConfiguration{"Credits, processing fee amount", N, FIXED, 12}
	fieldConfiguration[83] = FieldConfiguration{"Credits, transaction fee amount", N, FIXED, 12}
	fieldConfiguration[84] = FieldConfiguration{"Debits, processing fee amount", N, FIXED, 12}
	fieldConfiguration[85] = FieldConfiguration{"Debits, transaction fee amount", N, FIXED, 12}
	fieldConfiguration[86] = FieldConfiguration{"Total amount of credits", N, FIXED, 15}
	fieldConfiguration[87] = FieldConfiguration{"Credits, reversal amount", N, FIXED, 15}
	fieldConfiguration[88] = FieldConfiguration{"Total amount of debits", N, FIXED, 15}
	fieldConfiguration[89] = FieldConfiguration{"Debits, reversal amount", N, FIXED, 15}
	fieldConfiguration[90] = FieldConfiguration{"Original data elements", N, FIXED, 42}
	fieldConfiguration[91] = FieldConfiguration{"File update code", AN, FIXED, 1}
	fieldConfiguration[92] = FieldConfiguration{"File security code", N, FIXED, 2}
	fieldConfiguration[93] = FieldConfiguration{"Response indicator", N, FIXED, 5}
	fieldConfiguration[94] = FieldConfiguration{"Service indicator", AN, FIXED, 7}
	fieldConfiguration[95] = FieldConfiguration{"Replacement amounts", AN, FIXED, 42}
	fieldConfiguration[96] = FieldConfiguration{"Message security code ", AN, FIXED, 8}
	fieldConfiguration[97] = FieldConfiguration{"Amount, net settlement", N, FIXED, 16}
	fieldConfiguration[98] = FieldConfiguration{"Payee", ANS, FIXED, 25}
	fieldConfiguration[99] = FieldConfiguration{"Settlement institution identification code", N, LLVAR, 11}
	fieldConfiguration[100] = FieldConfiguration{"Receiving institution identification code", N, LLVAR, 11}
	fieldConfiguration[101] = FieldConfiguration{"File name", ANS, FIXED, 17}
	fieldConfiguration[102] = FieldConfiguration{"Account identification 1", ANS, LLVAR, 28}
	fieldConfiguration[103] = FieldConfiguration{"Account identification 2", ANS, LLVAR, 28}
	fieldConfiguration[104] = FieldConfiguration{"Transaction description", ANS, LLLVAR, 100}
	fieldConfiguration[105] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[106] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[107] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[108] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[109] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[110] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[111] = FieldConfiguration{"Reserved for ISO use", ANS, LLLVAR, 999}
	fieldConfiguration[112] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[113] = FieldConfiguration{"Authorizing agent institution id code", N, LLVAR, 11}
	fieldConfiguration[114] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[115] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[116] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[117] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[118] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[119] = FieldConfiguration{"Reserved for national use", ANS, LLLVAR, 999}
	fieldConfiguration[120] = FieldConfiguration{"Reserved for private use", ANS, LLLVAR, 999}
	fieldConfiguration[121] = FieldConfiguration{"Reserved for private use", ANS, LLLVAR, 999}
	fieldConfiguration[122] = FieldConfiguration{"Reserved for private use", ANS, LLLVAR, 999}
	fieldConfiguration[123] = FieldConfiguration{"Reserved for private use", ANS, LLLVAR, 999}
	fieldConfiguration[124] = FieldConfiguration{"Info Text", ANS, LLLVAR, 255}
	fieldConfiguration[125] = FieldConfiguration{"Network management information", ANS, LLLVAR, 999}
	fieldConfiguration[126] = FieldConfiguration{"Issuer trace id", ANS, LLLVAR, 999}
	fieldConfiguration[127] = FieldConfiguration{"Reserved for private use", ANS, LLLVAR, 999}
	fieldConfiguration[128] = FieldConfiguration{"Message Authentication code", B, FIXED, 16}
}
