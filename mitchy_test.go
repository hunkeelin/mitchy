package mitchy

import (
	"fmt"
	"testing"
)

func TestMitchy(t *testing.T) {
	hi()
	panic(fmt.Errorf("you suck"))
}
