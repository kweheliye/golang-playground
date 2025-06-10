package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrNotExist = fmt.Errorf("not exist")

type ValidationError struct {
	Field, Issue string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s %s", e.Field, e.Issue)
}
func validate() error {
	return fmt.Errorf("input invalid: %w", &ValidationError{"age ", "must be > 18"})

}

func loadResource() error {
	return fmt.Errorf("loadResource: %w", ErrNotFound)
}

func main() {

	//********error.IS
	errIs := loadResource()
	if errors.Is(errIs, ErrNotExist) {
		fmt.Println("not found")
	} else {
		fmt.Println("unexpected error:", errIs)
	}
	//********error.has

	var ve *ValidationError
	errHas := validate()
	if errors.As(errHas, &ve) {
		fmt.Printf("Validation failed %s: %s \n", ve.Field, ve.Issue)
	}
}
