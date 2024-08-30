package entity

import (
	"testing"
)

func TestAuth_TableName(t *testing.T) {
    expected := "auth.auth"
    actual := (Auth{}).TableName()
    if actual != expected {
        t.Errorf("Expected table name %s, but got %s", expected, actual)
    }
}
func TestAuthConfirmation_TableName(t *testing.T) {
    expected := "auth.auth_confirmation"
    actual := (AuthConfirmation{}).TableName()
    if actual != expected {
        t.Errorf("Expected table name %s, but got %s", expected, actual)
    }
}
func TestAuthOAuth_TableName(t *testing.T) {
    expected := "auth.auth_oauth"
    actual := (AuthOAuth{}).TableName()
    if actual != expected {
        t.Errorf("Expected table name %s, but got %s", expected, actual)
    }
}
func TestAuthChangeEmail_TableName(t *testing.T) {
    expected := "auth.auth_change_email"
    actual := (AuthChangeEmail{}).TableName()
    if actual != expected {
        t.Errorf("Expected table name %s, but got %s", expected, actual)
    }
}
func TestAuthChangePassword_TableName(t *testing.T) {
    expected := "auth.auth_change_password"
    actual := (AuthChangePassword{}).TableName()
    if actual != expected {
        t.Errorf("Expected table name %s, but got %s", expected, actual)
    }
}