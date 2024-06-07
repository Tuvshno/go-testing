package iteration

func Repeat(message string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + message
	}

	return repeated
}
