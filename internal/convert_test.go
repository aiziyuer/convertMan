package internal

import (
	"fmt"
	"testing"
)

func TestConvertFiles(t *testing.T) {

	s, _ := ConvertFiles("yaml", "yaml", []string{"default.yaml", "new.yaml"})
	fmt.Printf(s)

}
