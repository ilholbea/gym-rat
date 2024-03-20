package config

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const MockHost = "mock-host"
const MockPort = "mock-port"
const MockUser = "mock-user"
const MockPassword = "mock-password"
const MockDatabase = "mock-database"
const MockSchema = "mock-schema"

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name        string
		expected    *Config
		expectedErr error
	}{
		{
			name: "correct values",
			expected: &Config{
				DbConfig: Database{
					Host:     MockHost,
					Port:     MockPort,
					User:     MockUser,
					Password: MockPassword,
					Database: MockDatabase,
					Schema:   MockSchema,
				},
			},
			expectedErr: nil,
		},
		{
			name:        "missing values",
			expected:    &Config{},
			expectedErr: errors.New("unable to load environment variable"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setEnvironmentVariables(t, test.expected)
			conf, err := NewConfig()
			if err != nil {
				assert.Equal(t, test.expectedErr, err, "expected %q, got %q", test.expectedErr, err)
			}
			assert.Equal(t, test.expected, conf, "expected %q, go %q", test.expected, conf)
		})
	}

}

func setEnvironmentVariables(t *testing.T, conf *Config) {
	t.Setenv("DB_HOST", conf.DbConfig.Host)
	t.Setenv("DB_PORT", conf.DbConfig.Port)
	t.Setenv("DB_USER", conf.DbConfig.User)
	t.Setenv("DB_PASSWORD", conf.DbConfig.Password)
	t.Setenv("DB_DATABASE", conf.DbConfig.Database)
	t.Setenv("DB_SCHEMA", conf.DbConfig.Schema)
}
