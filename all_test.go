package textart

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	fmt.Print(Gen(`
1234567890
abcdefghijklmnopqrstuvwxyz
ABCDEFGHIJKLMNOPQRSTUVWXYZ
$%^&*~
!?@#,./|\;:-=+_()<>[]{}'"` + "`"))
}
