package utils

func MaskAadhaar(number string) string {
	if len(number) < 4 {
		return "XXXX"
	}
	return "XXXX-XXXX-" + number[len(number)-4:]
}

func MaskPAN(number string) string {
	if len(number) < 4 {
		return "****"
	}
	return "****" + number[len(number)-4:]
}