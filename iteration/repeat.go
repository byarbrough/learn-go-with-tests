package iteration

// Repeat char 5 times
func Repeat(c string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += c
	}
	return repeated
}
