package concate

import (
	"bytes"
	"strings"
)

func simplePlus(concateStep int) {
	var str string
	for i := 0; i < concateStep; i++ {
		str += "a"
	}
}

func buf(concateStep int) {
	var buf bytes.Buffer
	for i := 0; i < concateStep; i++ {
		buf.WriteString("a")
	}
}

func builder(concateStep int) {
	var b strings.Builder

	for i := 0; i < concateStep; i++ {
		b.WriteString("a")
	}
}
