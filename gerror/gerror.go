package gerror

func FailOnError(e error) {
	if e != nil {
		panic(e)
	}
}
