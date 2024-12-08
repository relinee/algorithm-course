package main

import (
	"strings"
)

type Rotor struct{}

var (
	fwdMap = make(map[int]int)
	resMap = make(map[int]int)

	position int
)

func NewRotor(forward string) *Rotor {
	for i, c := range forward {
		fwdMap[i] = int(c) - 'A'
		resMap[int(c)-'A'] = i
	}
	position = 0

	return &Rotor{}
}

func (r *Rotor) Forward(c int) int {
	offset := (c - 'A' + position) % 26
	return fwdMap[offset] + 'A'
}

func (r *Rotor) Reverse(c int) int {
	offset := resMap[c-'A']
	return 'A' + (offset-position+26)%26
}

func (r *Rotor) ResetPosition() {
	position = 0
}

func (r *Rotor) ConvertToString(codes []int) string {
	var builder strings.Builder
	for _, c := range codes {
		s := rune(c)
		builder.WriteString(string(s))
	}
	return builder.String()
}

func (r *Rotor) Step() bool {
	newPosition := (position + 1) % 26
	oldPosition := position
	position = newPosition

	return newPosition < oldPosition
}
