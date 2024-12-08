package main

import (
	"fmt"
	"strings"
)

type Enigma struct {
	rotors    []*Rotor
	reflector *Reflector
}

func NewEnigmaMachine(rotors []*Rotor, reflector *Reflector) *Enigma {
	return &Enigma{
		rotors:    rotors,
		reflector: reflector,
	}
}

func (e *Enigma) StepRotors() {
	rotateNext := true
	for _, rotor := range e.rotors {
		if rotateNext {
			rotateNext = rotor.Step()
		} else {
			break
		}
	}
}

func (e *Enigma) EncryptChar(c int) string {
	if c < 'A' || c > 'Z' {
		return string(rune(c))
	}

	e.StepRotors()
	for _, rotor := range e.rotors {
		c = rotor.Forward(c)
		//println(string(rune(c)))
	}

	c = e.reflector.Reflect(c)

	for i := len(e.rotors) - 1; i >= 0; i-- {
		c = e.rotors[i].Reverse(c)
	}

	return string(rune(c))
}

func (e *Enigma) EncryptMessage(message string) string {
	var result strings.Builder
	for _, c := range strings.ToUpper(message) {
		result.WriteString(e.EncryptChar(int(c)))
	}
	return result.String()
}

func (e *Enigma) ResetPositions() {
	for _, rotor := range e.rotors {
		rotor.ResetPosition()
	}
}

func main() {
	//rotor1 := NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ")
	//rotor2 := NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE")
	//rotor3 := NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO")
	//reflector := NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")

	rotor1 := NewRotor("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rotor2 := NewRotor("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rotor3 := NewRotor("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	reflector := NewReflector("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	enigma := NewEnigmaMachine([]*Rotor{rotor1, rotor2, rotor3}, reflector)

	text := "Hello world"
	encryptedText := enigma.EncryptMessage(text)
	fmt.Println("Зашифрованное сообщение:", encryptedText)

	enigma.ResetPositions()

	decrypted := enigma.EncryptMessage(encryptedText)
	fmt.Println("Расшифрованное сообщение:", decrypted)
}
