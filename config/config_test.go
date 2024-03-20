package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const MY_CONSTANT_VALUE = "my-value"

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		expected *Config
	}{
		{name: "test", expected: &Config{Database{Host: MY_CONSTANT_VALUE}}},
	}

	setEnvironmentVariables()
	defer unsetEnvironmentVariables()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := NewConfig()
			assert.Equal(t, test.expected, conf, "expected %q, go %q", test.expected, conf)
		})
	}

}

func setEnvironmentVariables() {
	_ = os.Setenv("MYCONSTANT", MY_CONSTANT_VALUE)
}

func unsetEnvironmentVariables() {
	_ = os.Unsetenv("MYCONSTANT")
}
