package utils

import (
	"os"
	"testing"
)

func TestGetEnvValue(t *testing.T) {
	err := os.Setenv("GO_TEST_KEY", "somevalue")
	if err != nil {
		t.Errorf("failed, erroring in setting up test")
	}
	expected := "somevalue"
	result := GetEnv("GO_TEST_KEY", "")

	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}

func TestGetEnvDefaultValue(t *testing.T) {
	err := os.Unsetenv("GO_TEST_KEY")
	if err != nil {
		t.Errorf("failed, erroring in setting up test")
	}
	expected := "defaultvalue"
	result := GetEnv("GO_TEST_KEY", "defaultvalue")

	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}
