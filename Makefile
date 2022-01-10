
build:
	@echo "Building Template Fetcher CLI" 
	go build -o bin/tfetch app/tfetch.go

install:
	@echo "Installing Template Fetcher CLI" 
	go install app/tfetch.go 
