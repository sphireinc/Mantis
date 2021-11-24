package _generics

type strings interface {
	type string, byte, rune
}

type pureStrings interface {
	type string, byte
}

type intStrings interface {
	type byte, rune
}