package helpers

import (
	"strings"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
)

// Determine owner and repo name from argument
func DetermineOwnerAndRepo(arg string) (string, string, error) {

	// Declare owner and repo
	var owner, repo string

	// Get current repository
	currentRepo, err := gh.CurrentRepository()
	if err != nil {
		return "", "", err
	}
	// Set default owner and repo
	owner, repo = currentRepo.Owner(), currentRepo.Name()

	// If arguments are provided, parse it to get the specified owner and repo
	if len(arg) > 0 {
		// If owner is not provided, add it to the argument before parsing
		if !strings.Contains(arg, "/") {
			arg = owner + "/" + arg
		}
		// Parse argument to get owner and repo
		parsed, err := repository.Parse(arg)
		if err != nil {
			return "", "", err
		}
		owner, repo = parsed.Owner(), parsed.Name()
	}

	// Return owner and repo
	return owner, repo, nil
}
