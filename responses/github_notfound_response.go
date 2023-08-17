package responses

func GithubNotFoundResponse(http_code_message int) (int, string, string) {
	// Default Value
	status        := "failed"
	message       := "Yahhh... username github anda tidak tersedia"

	return http_code_message, status, message
}