package exercises

type VerifiedResult struct {
	Verified chan Result
	Total    int
}

// Verify checks if all the exercises compile and pass the tests
func Verify() (VerifiedResult, error) {
	verified := make(chan Result)

	exs, err := List()
	if err != nil {
		return VerifiedResult{Verified: verified}, err
	}

	go func() {
		for _, ex := range exs {
			result, _ := Run(ex.Name)
			verified <- result
		}
		close(verified)
	}()

	return VerifiedResult{Verified: verified, Total: len(exs)}, nil
}
