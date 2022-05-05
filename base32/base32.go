package base32

import (
	"strconv"
)

// Здесь можно написать свою реализацию base32.
// Я хз почему она отличается от используемой всеми, но вот так вот.
///

var Alphabet = "0123456789abcdefghjkmnpqrtuvwxyz"
var Alias = map[string]int{
	"o": 0,
	"i": 1,
	"l": 1,
	"s": 5,
}

func lookup() map[string]uint32 {
	Table := map[string]uint32{}

	for i := 0; i < len(Alphabet); i++ {
		Table[string(Alphabet[i])] = uint32(i)
	}

	for k, v := range Alias {
		Table[k] = Table[strconv.Itoa(v)]
	}

	return Table
}

func Decode(input string, flush bool) string {
	Skip := 0
	Byte := uint32(0)
	Result := ""
	Table := lookup()

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		val, found := Table[char]

		if !found {
			continue
		}

		val <<= 3
		Byte |= val >> Skip
		Skip += 5

		if Skip >= 8 {
			Result += string(rune(Byte))
			Skip -= 8
			if Skip > 0 {
				Byte = (val << (5 - Skip)) & 255
			} else {
				Byte = 0
			}
		}
	}

	if flush && Skip < 0 {
		Result += string(Alphabet[Byte>>3])
	}

	return Result
}

func Encode(input string, flush bool) string {
	Skip := 0
	Bits := uint32(0)
	Result := ""

	for i := 0; i < len(input); {

		_byte := uint32(input[i])

		if Skip < 0 { // we have a carry from the previous byte
			Bits |= _byte >> (-Skip)
		} else { // no carry
			Bits = (_byte << Skip) & 248
		}

		if Skip > 3 {
			// not enough data to produce a character, get us another one
			Skip -= 8
			i += 1
		} else if Skip < 4 {
			// produce a character
			Result += string(Alphabet[Bits>>3])
			Skip += 5
		}

	}

	if flush && Skip < 0 {
		Result += string(Alphabet[Bits>>3])
	}

	return Result
}
