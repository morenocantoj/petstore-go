package errors

// Check : If the error is not null stops the execution
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
