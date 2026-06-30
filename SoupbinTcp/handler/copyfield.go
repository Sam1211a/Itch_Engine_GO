package handler

func CopyField(dst []byte, value string) {
	for i := range dst {
		dst[i] = ' '
	}
	copy(dst, []byte(value))
}
