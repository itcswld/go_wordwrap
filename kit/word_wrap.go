package kit

func SliceStr(txt []rune, breakLen int) (arr_str []string) {
	for i := 1; i < len(txt); i++ {
		if i%breakLen == 0 {
			//slicing
			arr_str = append(arr_str, string(txt[:i])) //[0][1]
			txt = txt[i:]                              //[i]~[len(s)]
			i = 1
		}
	}
	arr_str = append(arr_str, string(txt[:]))
	return arr_str
}

func WordWrap(text []rune, breakLen int) (data []string) {
	n := 0.0
	pre := 0
out:
	for j, v := range text {
		if IsChinese(v) {
			n += 1
		} else {
			n += 0.5
		}
		if len(text) <= breakLen || (j-pre == len(text)-1 && n < float64(breakLen)) {
			data = append(data, string(text[:]))
			break out
		} else if len(text) > breakLen && n >= float64(breakLen) {
			data = append(data, string(text[:j-pre]))
			text = text[j-pre:]
			pre += j - pre
			n = 0.0
		}
	}

	return data
}
