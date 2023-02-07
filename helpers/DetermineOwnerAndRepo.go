package helpers

import (
	"strings"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
)

// Determine the owner and repo name from arguments passed to the command
// ... or from the current repository.
func DetermineOwnerAndRepo(args []string) (string, string, error) {

	// Get current repository
	currentRepo, err := gh.CurrentRepository()
	if err != nil {
		return "", "", err
	}
	// Initialize owner and repo using the current repository
	owner, repo := currentRepo.Owner(), currentRepo.Name()

	// If arguments are provided, use them to determine owner and repo
	if len(args) > 0 {
		arg := args[0] // Only take the first argument (`owner/repo` or `repo`)

		// If the argument is a repo name, assume it's in the current owner namespace
		// and prepend the owner name to the argument to form a `owner/repo` string
		if !strings.Contains(arg, "/") {
			arg = owner + "/" + arg
		}

		// Parse the argument to determine owner and repo
		parsed, err := repository.Parse(arg)
		if err != nil {
			return "", "", err
		}
		// Override owner and repo
		owner, repo = parsed.Owner(), parsed.Name()
	}

	// Return owner and repo
	return owner, repo, nil
}
