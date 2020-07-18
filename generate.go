package textart

import (
	"strings"
)

func Gen(s string) string {
	ss := []string{}
	for _, v := range strings.Split(s, "\n") {
		builder := &builder{}
		builder.Append(v)
		ss = append(ss, builder.String())
	}
	ss = append(ss, "")
	return strings.Join(ss, "\n")
}

type builder struct {
	rows [][]rune
}

func (b *builder) String() string {
	s := []string{}
	for _, v := range b.rows {
		s = append(s, string(v))
	}
	return strings.Join(s, "\n")
}

func (b *builder) Append(s string) {
	for _, v := range []rune(s) {
		if int(v) >= len(mapping) {
			b.append(mapping[0])
			continue
		}
		m := mapping[v]
		if m == "" {
			b.append(mapping[0])
			continue
		}
		b.append(m)
	}
}

func (b *builder) fill() {
	max := 0
	for _, v := range b.rows {
		if len(v) > max {
			max = len(v)
		}
	}
	for i, v := range b.rows {
		if sub := max - len(v); sub > 0 {
			b.rows[i] = append(b.rows[i], []rune(strings.Repeat(" ", sub))...)
		}
	}
}

func (b *builder) append(s string) {
	rows := strings.Split(s, "\n")
	if sub := len(rows) - len(b.rows); sub > 0 {
		b.rows = append(b.rows, make([][]rune, sub)...)
		b.fill()
	}
	for i, v := range rows {
		b.rows[i] = append(b.rows[i], []rune(v)...)
	}
	b.fill()
}
