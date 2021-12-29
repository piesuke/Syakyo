package iteration

const repeatCount = 5

func Repeat(word string) (repeatedWords string) {

	for i := 0; i < repeatCount; i++ {
		repeatedWords += word
	}
	return
}
