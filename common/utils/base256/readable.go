package base256

import (
	"encoding/binary"
	"fmt"
	"math/bits"

	utls "github.com/cyrildever/go-utls/common/utils"
)

// Based on the 512-characters UTF-8 table - @see https://www.utf8-chartable.de/unicode-utf8-table.pl?number=512
var CHARSET = []rune("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^`abcdefghijklmnopqrstuvwxyz{|}€¡¢£¤¥¦§¨©ª«¬®¯°±²³´µ¶·¸¹»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷ùúûüýÿăąĊčđĕĘğħĩĭıĵķĿŀŁłňŋŏœŖřŝşŦŧũūůŲŵſƀƁƂƄƆƇƔƕƗƙƛƜƟƢƥƦƧƩƪƭƮưƱƲƵƸƺƾǀǁǂƿǬǮǵǶǹǻǿ")

//--- TYPES

// Readable represents a character in a readable base-256 charset
type Readable string

//--- METHODS

// Bytes ...
func (b256 Readable) Bytes() []byte {
	var barray []byte
	for _, char := range b256 {
		barray = append(barray, byte(IndexOf(char)))
	}
	return barray
}

// IsEmpty ...
func (b256 Readable) IsEmpty() bool {
	return b256.String() == ""
}

// Len ...
func (b256 Readable) Len() int {
	length := 0
	for range b256 {
		length++
	}
	return length
}

// NonEmpty ...
func (b256 Readable) NonEmpty() bool {
	return !b256.IsEmpty()
}

// String ...
//
// NB: By passing `true` as argument, you signify you don't want to use the readable charset but force the ASCII rendering,
// implying that you'd probably get unreadable characters if the underlying byte array uses byte that ranges over 127 but that
// you'd get the correct length when using the len() function on the resulting string
func (b256 Readable) String(useAscii ...bool) string {
	if len(useAscii) == 1 && useAscii[0] {
		return string(b256.Bytes())
	}
	return string(b256)
}

// ToHex ...
func (b256 Readable) ToHex() string {
	return utls.ToHex(b256.Bytes())
}

// ToNumber returns the stringified value of the underlying integer
//
// You may pass a minimum size for the returned number to add padding zeroes accordingly
func (b256 Readable) ToNumber(nbZeroPad ...int) string {
	var length int
	if len(nbZeroPad) > 0 {
		length = nbZeroPad[0]
	}
	format := "%" + fmt.Sprintf("0%dd", length)
	return fmt.Sprintf(format, b256.Uint64())
}

// Uint64 ...
func (b256 Readable) Uint64() uint64 {
	buf := make([]byte, 8)
	for i, b := range b256.Bytes() {
		buf[i+8-b256.Len()] = b
	}
	return binary.BigEndian.Uint64(buf)
}

//--- FUNCTION

// CharAt returns the character string at the passed index in the Base-256 character set
func CharAt(index int) string {
	return string(CHARSET[index : index+1])
}

// IndexOf ...
func IndexOf(char rune) int {
	for i, r := range CHARSET {
		if r == char {
			return i
		}
	}
	return -1
}

// ToBase256Readable ...
func ToBase256Readable(barray []byte) Readable {
	str := ""
	for _, b := range barray {
		str += CharAt(int(b))
	}
	return Readable(str)
}

// HexToReadable ...
func HexToReadable(hex string) (b256 Readable, err error) {
	bytes, err := utls.FromHex(hex)
	if err != nil {
		return
	}
	b256 = ToBase256Readable(bytes)
	return
}

// NumberToReadable ...
func NumberToReadable(n uint64) (b256 Readable, err error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	bytes := buf[bits.LeadingZeros64(n)>>3:]
	b256 = ToBase256Readable(bytes)
	return
}
