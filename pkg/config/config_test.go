package config

import (
	"strings"
	"sync"
	"testing"
)

func TestNewPreservesInitializationError(t *testing.T) {
	t.Cleanup(func() {
		instance = nil
		instanceErr = nil
		once = sync.Once{}
	})

	instance = nil
	instanceErr = nil
	once = sync.Once{}
	t.Setenv("ENV", "development")
	t.Setenv("MYSQL_DSN", "")

	if _, err := New(); err == nil {
		t.Fatal("expected the first initialization to fail")
	}
	if _, err := New(); err == nil {
		t.Fatal("expected the cached initialization error on the second call")
	}
}

func TestLoadEnvRejectsMissingDatabaseDSN(t *testing.T) {
	t.Setenv("ENV", "development")
	t.Setenv("MYSQL_DSN", "")

	_, err := loadEnv()
	if err == nil {
		t.Fatal("expected an error when MYSQL_DSN is missing")
	}
	if !strings.Contains(err.Error(), "MYSQL_DSN") {
		t.Fatalf("expected MYSQL_DSN error, got %q", err)
	}
}

func TestLoadEnvRejectsMissingProductionSecrets(t *testing.T) {
	t.Setenv("ENV", "production")
	t.Setenv("MYSQL_DSN", "user:password@tcp(localhost:3306)/neuronews")
	t.Setenv("ADMIN_EMAIL", "")
	t.Setenv("ADMIN_PWD", "")
	t.Setenv("SESSION_SECRET", "")

	_, err := loadEnv()
	if err == nil {
		t.Fatal("expected an error when production secrets are missing")
	}

	for _, name := range []string{"ADMIN_EMAIL", "ADMIN_PWD", "SESSION_SECRET"} {
		if !strings.Contains(err.Error(), name) {
			t.Errorf("expected error to mention %s, got %q", name, err)
		}
	}
}

func TestLoadEnvAcceptsDevelopmentWithoutUnusedSecrets(t *testing.T) {
	const dsn = "user:password@tcp(localhost:3306)/neuronews"

	t.Setenv("ENV", "development")
	t.Setenv("MYSQL_DSN", dsn)
	t.Setenv("ADMIN_EMAIL", "")
	t.Setenv("ADMIN_PWD", "")
	t.Setenv("SESSION_SECRET", "")

	cfg, err := loadEnv()
	if err != nil {
		t.Fatalf("expected valid development config, got %v", err)
	}
	if cfg.MySQL.DSN != dsn {
		t.Fatalf("expected MYSQL_DSN %q, got %q", dsn, cfg.MySQL.DSN)
	}
}
