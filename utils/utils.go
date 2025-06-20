package utils

func IsLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
