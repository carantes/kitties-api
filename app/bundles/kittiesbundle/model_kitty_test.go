package kittiesbundle_test

import (
	"testing"

	"github.com/carantes/kitties-api/app/bundles/kittiesbundle"
)

func TestNewKittySpec(t *testing.T) {
	k := kittiesbundle.NewKitty("kittyName", "kittyBreed", "2020-01-01")

	t.Run("kitty should have name", func(t *testing.T) {
		expected := "kittyName"
		if k.Name != expected {
			t.Errorf("TestNewKittySpec result is incorrect, return %s, expected %s", k.Name, expected)
		}
	})

	t.Run("kitty should have breed", func(t *testing.T) {
		expected := "kittyBreed"
		if k.Breed != expected {
			t.Errorf("TestNewKittySpec result is incorrect, return %s, expected %s", k.Breed, expected)
		}
	})

	t.Run("kitty should have birthDate", func(t *testing.T) {
		expected := "2020-01-01"
		if k.BirthDate != expected {
			t.Errorf("TestNewKittySpec result is incorrect, return %s, expected %s", k.BirthDate, expected)
		}
	})
}

func TestNewKittyValidateSpec(t *testing.T) {

	t.Run("kitty is invalid if name is empty", func(t *testing.T) {
		k := kittiesbundle.NewKitty("", "kittyBreed", "2020-01-01")
		result := k.Validate()
		expected := false

		if result != expected {
			t.Errorf("TestNewKittyValidateSpec result is incorrect, return %t, expected %t", true, expected)
		}
	})

	t.Run("kitty is valid if all fields are correct", func(t *testing.T) {
		k := kittiesbundle.NewKitty("kittyName", "kittyBreed", "2020-01-01")
		result := k.Validate()
		expected := true

		if result != expected {
			t.Errorf("TestNewKittyValidateSpec result is incorrect, return %t, expected %t", true, expected)
		}
	})
}
