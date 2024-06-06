package errors

func AdditionalData(key string, values interface{}) ErrorOption {
	return func(e *Errors) {
		e.additionals[key] = values
	}
}
