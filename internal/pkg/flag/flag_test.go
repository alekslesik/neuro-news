package flag

import (
	"os"
	"testing"

	"github.com/alekslesik/neuro-news/pkg/config"
)

// TestInitWithInvalidArgs testing Init() with wrong args.
func TestInitWithInvalidArgs(t *testing.T) {
	cfg := &config.Config{}

	// save origin os.Args
	originalArgs := os.Args

	// Restore os.Args after test
	defer func() { os.Args = originalArgs }()

	// set wrong args for test
	os.Args = []string{"cmd", "-port=invalid"}

	// Вызов функции Init с неверными аргументами
	// call Init() with wrong args
	err := Init(cfg)
	if err == nil {
		t.Errorf("Init did not return an error for invalid args")
	}
}
