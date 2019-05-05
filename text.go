package main

import (
	"strings"
	"unicode"
)

type RuneType int

const (
	Letter RuneType = iota
	Digit
	Symbol
)

type ActionType int

const (
	Continue ActionType = iota
	SkipLast
	AddWord
	AddWord2
)

func ToConstantCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		strs[i] = strings.ToUpper(v)
	}
	return strings.Join(strs, "_")
}

func ToUnderlineCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		strs[i] = strings.ToLower(v)
	}
	return strings.Join(strs, "_")
}

func ToTitleCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		strs[i] = strings.ToTitle(v)[0:1] + strings.ToLower(v)[1:]
	}
	return strings.Join(strs, "")
}

func ToCamelCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		if i == 0 {
			strs[i] = strings.ToLower(v)
		} else {
			strs[i] = strings.ToTitle(v)[0:1] + strings.ToLower(v)[1:]
		}
	}
	return strings.Join(strs, "")
}

func ToDashCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		strs[i] = strings.ToLower(v)
	}
	return strings.Join(strs, "-")
}

func ToPathCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		strs[i] = strings.ToLower(v)
	}
	return strings.Join(strs, "/")
}

func ToPackageCase(str string) string {
	strs := SplitWord(str)
	for i, v := range strs {
		strs[i] = strings.ToLower(v)
	}
	return strings.Join(strs, ".")
}

func SplitWord(str string) (result []string) {
	//  Trim text
	str = strings.TrimSpace(str)

	// Helper Function...
	checkRuneType := func(r rune) RuneType {
		if unicode.IsLetter(r) {
			return Letter
		} else if unicode.IsDigit(r) {
			return Digit
		} else {
			return Symbol
		}
	}

	// Decision Function...
	handlers := map[RuneType]func(r rune, lastRune rune, lastType RuneType) ActionType{
		Letter: func(r rune, lastRune rune, lastType RuneType) ActionType {
			if lastType == Letter {
				if unicode.IsUpper(r) &&
					unicode.IsLower(lastRune) {
					return AddWord
				} else if unicode.IsLower(r) &&
					unicode.IsUpper(lastRune) {
					return AddWord2
				}
			}

			if lastType == Digit {
				return AddWord
			}
			if lastType == Symbol {
				return SkipLast
			}
			return Continue
		},
		Digit: func(r rune, lastRune rune, lastType RuneType) ActionType {
			if lastType == Symbol {
				return SkipLast
			}
			return Continue
		},
		Symbol: func(r rune, lastRune rune, lastType RuneType) ActionType {
			if lastType == Letter || lastType == Digit {
				return AddWord
			}
			return Continue
		},
	}

	// Start spliting string
	head := 0
	var lastRune rune
	lastType := Letter

	for i, r := range str {
		runeType := checkRuneType(r)

		// Make Decision
		var action ActionType
		if i != 0 {
			action = handlers[runeType](r, lastRune, lastType)
		}

		// Take Action
		switch action {
		case AddWord:
			result = append(result, str[head:i])
			head = i
		case AddWord2:
			s := str[head : i-1]
			if len(s) != 0 {
				result = append(result, s)
				head = i - 1
			}
		case SkipLast:
			head = i
		}

		// Save last value
		lastRune = r
		lastType = runeType
	}

	// Add last word to array
	if lastType == Letter || lastType == Digit {
		result = append(result, str[head:])
	}

	// Convert string Title Case
	for i, v := range result {
		result[i] = strings.Title(strings.ToLower(v))
	}

	return
}
