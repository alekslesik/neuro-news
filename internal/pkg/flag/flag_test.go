package flag

import (
	"os"
	"strconv"
	"testing"

	"github.com/alekslesik/neuro-news/pkg/config"
)

// TestInitWithInvalidArgs testing Init() with wrong args
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

// TestInitWithNoValidEnvFlag testing Init() with invalid 'env' flag
func TestInitWithNoValidEnvFlag(t *testing.T) {
	testCases := []struct {
		desc      string
		env       string
		want      error
		textError string
	}{
		{
			desc: "valid env development flag",
			env:  "development",
			want: nil,
		},
		{
			desc: "valid env staging flag",
			env:  "staging",
			want: nil,
		},
		{
			desc: "valid env production flag",
			env:  "production",
			want: nil,
		},
		{
			desc: "invalid env flag",
			env:  "invalidflag",
			want: ErrWrongEnv,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cfg := &config.Config{}

			// save origin os.Args
			originalArgs := os.Args

			// Restore os.Args after test
			defer func() { os.Args = originalArgs }()

			env := "-env=" + tC.env

			// set wrong args for test
			os.Args = []string{"cmd", env}

			err := Init(cfg)
			if err != tC.want {
				t.Errorf("want %v return %v", tC.want, err)
			}
		})
	}
}

// TestInitInvalidPortFlag testing Init() with invalid 'port' flag
func TestInitInvalidPortFlag(t *testing.T) {
	testCases := []struct {
		desc      string
		port      int
		want      error
		textError string
	}{
		{
			desc: "valid port flag",
			port: 1,
			want: nil,
		},
		{
			desc: "invalid port0",
			port: 0,
			want: ErrWrongPort,
		},
		{
			desc: "invalid port<0",
			port: -1,
			want: ErrWrongPort,
		},
		{
			desc: "invalid port>65535",
			port: 65536,
			want: ErrWrongPort,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cfg := &config.Config{}

			// save origin os.Args
			originalArgs := os.Args

			// Restore os.Args after test
			defer func() { os.Args = originalArgs }()

			port := "-port=" + strconv.Itoa(tC.port)

			// set wrong args for test
			os.Args = []string{"cmd", port}

			err := Init(cfg)
			if err != tC.want {
				t.Errorf("want %v return %v", tC.want, err)
			}
		})
	}
}
