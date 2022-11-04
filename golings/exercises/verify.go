package exercises

// Verify checks if all the exercises compile and pass the tests
func Verify() (<-chan Result, error) {
	verified := make(chan Result)

	exs, err := List()
	if err != nil {
		return verified, err
	}

	go func() {
		for _, ex := range exs {
			result, _ := Run(ex.Name)
			verified <- result
		}
		close(verified)
	}()

	return verified, nil
}
