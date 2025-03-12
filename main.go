package main

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func cleanString(s string) (string, bool) {
	s = strings.ToLower(s)

	// Transform para eliminar acentos
	t := transform.Chain(
		norm.NFD,                           // 1️⃣ Descompone caracteres en base + acento
		runes.Remove(runes.In(unicode.Mn)), // 2️⃣ Elimina los caracteres de la categoría "Mn" (acentos)
		norm.NFC) // 3️⃣ Recomponer el texto limpio sin acentos

	s, _, _ = transform.String(t, s)

	if strings.Contains(s, " ") {
		return "", false
	}

	return s, true
}

func isIsogram(s string) bool {
	if s == "" {
		return true
	}

	cleaned, valid := cleanString(s)
	if !valid {
		return false
	}

	letters := make(map[rune]bool)
	for _, char := range cleaned {
		if letters[char] {
			return false
		}
		letters[char] = true
	}

	return true
}

// Test
func main() {
	fmt.Println(isIsogram("hello"))
	fmt.Println(isIsogram("lumberjack"))
	fmt.Println(isIsogram("isogram"))
	fmt.Println(isIsogram("alphabet"))
	fmt.Println(isIsogram("áéíóú"))
	fmt.Println(isIsogram("España"))
	fmt.Println(isIsogram("Mañana"))
	fmt.Println(isIsogram("único"))
	fmt.Println(isIsogram("camión"))
	fmt.Println(isIsogram("palabra clave"))
	fmt.Println(isIsogram(""))
}
