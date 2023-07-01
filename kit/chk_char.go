//Check If The Rune is Chinese Character
package kit

import (
	"unicode"
)

func IsChinese(char rune) (b bool) {
	//chinese
	if unicode.Is(unicode.Han, char) {
		b = true
	} else {
		//not chinese
		b = false
	}
	return b
}
