package iterations

func Repeat(chunk string, limit int) string {

	var result string

	for i := 0; i < limit; i++ {
		result += chunk
	}

	return result

}
