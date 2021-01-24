package padding

const (
	CHARACTER rune = 2 // Unicode U+0002: start-of-text
)

//--- FUNCTIONS

// Apply ...
func Apply(data []byte) []byte {
	if len(data)%2 == 0 {
		return data
	}
	padded := []byte{byte(CHARACTER)}
	padded = append(padded, data...)
	return padded
}

// Unapply ...
func Unapply(padded []byte) []byte {
	if len(padded) == 0 {
		return nil
	}
	if padded[0] != byte(CHARACTER) {
		return padded
	}
	return padded[1:]
}
