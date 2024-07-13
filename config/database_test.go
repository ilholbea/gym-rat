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
		expected    *DatabaseConfig
		expectedErr error
	}{
		{
			name: "correct values",
			expected: &DatabaseConfig{
				Host:     MockHost,
				Port:     MockPort,
				User:     MockUser,
				Password: MockPassword,
				Database: MockDatabase,
				Schema:   MockSchema,
			},
			expectedErr: nil,
		},
		{
			name:        "missing values",
			expected:    &DatabaseConfig{},
			expectedErr: errors.New("unable to load environment variable"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setEnvironmentVariables(t, test.expected)
			conf, err := GetDatabaseConfig()
			if err != nil {
				assert.Equal(t, test.expectedErr, err, "expected %q, got %q", test.expectedErr, err)
			}
			assert.Equal(t, test.expected, conf, "expected %q, go %q", test.expected, conf)
		})
	}

}

func setEnvironmentVariables(t *testing.T, conf *DatabaseConfig) {
	t.Setenv("DB_HOST", conf.Host)
	t.Setenv("DB_PORT", conf.Port)
	t.Setenv("DB_USER", conf.User)
	t.Setenv("DB_PASSWORD", conf.Password)
	t.Setenv("DB_DATABASE", conf.Database)
	t.Setenv("DB_SCHEMA", conf.Schema)
}
