package iteration

const repeatedCout = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCout; i++ {
		repeated += character
	}
	return repeated
}
