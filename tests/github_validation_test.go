package tests

import (
	"testing"

	validation "backend_community_rest/validations"
	exception "backend_community_rest/exceptions"
)


// Github Validation Test
func TestGithubValidation(t *testing.T) {

	username := "belajarqywok"

	value, err := validation.ValidateGitHubUsername(username)
	if err != nil {
		exception.TryCatchError(err)
	}

	if !value {
		t.Errorf("Username \"%q\" not found!", username)
	}
}