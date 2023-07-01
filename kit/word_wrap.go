package kit

import "fmt"

// "strconv"

func SliceStr(s string, n int) (arr_str []string) {
	txtRune := []rune(s)
	for i := 1; i < len(txtRune); i++ {
		if i%n == 0 {
			//slicing
			arr_str = append(arr_str, string(txtRune[:i])) //[0][1]
			txtRune = txtRune[i:]                          //[i]~[len(s)]
			i = 1
		}
	}
	arr_str = append(arr_str, string(txtRune[:]))
	return arr_str
}

func WordWrap(text []rune, breakLen int) (data []string) {
	n := 0.0
	t := 0
out:
	for j, v := range text {
		fmt.Println("j=", j, "n =", n)
		fmt.Println("new:", string(text[:]), "| len= ", len(text))
		if IsChinese(v) {
			n += 1
		} else {
			n += 0.5
		}
		if len(text) <= breakLen || (j-t == len(text)-1 && n < float64(breakLen)) {
			fmt.Println("j=", j, "n =", n)
			data = append(data, string(text[:]))
			fmt.Println("last:", string(text[:]))
			break out
		} else if len(text) > breakLen && n >= float64(breakLen) {
			fmt.Println("j=", j, ";n =", n, ";t =", j)
			data = append(data, string(text[:j-t]))
			fmt.Println("cut:", string(text[:j-t]))
			text = text[j-t:]
			t += j
			n = 0.0
		}
	}

	return data
}
