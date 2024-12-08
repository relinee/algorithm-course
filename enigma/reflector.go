package main

type Reflector struct {
	mapping string
}

var rMap = make(map[int]int)

func NewReflector(mapping string) *Reflector {
	for i, c := range mapping {
		rMap[i] = int(c) - 'A'
	}

	return &Reflector{mapping: mapping}
}

func (r *Reflector) Reflect(c int) int {
	return rMap[c-'A'] + 'A'
}
