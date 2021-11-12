package encode

var (
	alphabetLowerCase = "abcdefghijklmnopqrstuvwxyz"
	alphabetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shift             = 32
)

func EncodeString(str string) string {
	str = rotate32(str, shift)
	return str
}

func DecodeString(str string) string {
	str = rotate32(str, -shift)
	return str
}

func rotate32(str string, shift int) string {
	byteSlice := []byte(str)
	shift = shift % 26

	for i, charVal := range byteSlice {
		if charVal >= 'a' && charVal <= 'z' {
			byteSlice[i] = alphabetLowerCase[(int((26+(charVal-'a')))+shift)%26]
		} else if charVal >= 'A' && charVal <= 'Z' {
			byteSlice[i] = alphabetUpperCase[(int((26+(charVal-'A')))+shift)%26]
		}
	}
	return string(byteSlice)
}
