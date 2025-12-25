package util

const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ToBase62(n int64) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result += string(characters[n%62])
		n /= 62
	}
	return result
}
