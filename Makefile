test:
	@cd golings && go test -coverprofile=coverage.out -v $$(go list ./... | grep -v fixtures/error1)
