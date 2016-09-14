package strings

const (
	RuneSelf  = 0x80         // Characters below Runeself are represented as themselves in a single byte.
	MaxASCII  = '\u007F'     // Maximum ASCII value.
	MaxLatin1 = '\u00FF'     // Maximum Latin-1 value.
	MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
)
