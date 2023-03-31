package helpers

import (
	"os"

	"github.com/cli/go-gh/pkg/browser"
)

// OpenInBrowser opens the specified URL in the default browser.
func OpenInBrowser(url string) error {
	b := browser.New("", os.Stdout, os.Stderr)
	return b.Browse(url)
}
