build:
	@go build -o server.run ./http
test:
	@go test ./... -cover
cover:
	@go test ./... -coverprofile=c.out && cat c.out | grep -v "mock_" > cover.out && go tool cover -func=cover.out
cover-html:
	@go test ./... -coverprofile=c.out && cat c.out | grep -v "mock_" > cover.out && go tool cover -html=cover.out