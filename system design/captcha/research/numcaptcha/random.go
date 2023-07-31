package numcaptcha

func randomID() string {
	return "random-string"
}

// randomBytesMod returns a byte slice of the given length, where each byte is
// a random number module mod.
func randomBytesMod(length int, mod byte) (b []byte) {
	// if length == 0 {
	// 	return nil
	// }
	// return
	return []byte("randomBytesMod")
}

// RandomDigits returns a bytes slice of the given length containing
// pseudorandom number in range 0-9. The slice can be used as  a captcha
// solution
func RandomDigits(length int) []byte {
	return randomBytesMod(length, 10)
}
