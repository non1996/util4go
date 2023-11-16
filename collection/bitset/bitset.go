package bitset

import (
	"bytes"
	"fmt"

	"github.com/non1996/util4go/function/value"
)

const (
	addressBitsPerWord int    = 6
	bitsPerWord        int    = 1 << addressBitsPerWord
	bitsIndexMask      int    = bitsPerWord - 1
	wordMask           uint64 = 0xffffffffffffffff
)

type Bitset struct {
	words      []uint64
	wordsInUse int
}

func New(nbits int) *Bitset {
	b := &Bitset{}
	b.initWords(nbits)
	return b
}

func ValueOf(b []byte) *Bitset {
	n := len(b)
	for ; n > 0 && b[n-1] == 0; n-- {
		value.Void(n)
	}

	words := make([]uint64, (n+7)/8)
	wordIdx := 0

	for i := 0; i < n; i++ {
		words[wordIdx] = uint64(b[i]) << (i % 8)
		if i != 0 && i%8 == 0 {
			wordIdx++
		}
	}

	return &Bitset{
		words:      words,
		wordsInUse: len(words),
	}
}

func (b *Bitset) IsEmpty() bool {
	return b.wordsInUse == 0
}

func (b *Bitset) Length() int {
	if b.wordsInUse == 0 {
		return 0
	}

	// TODO
	return bitsPerWord*(b.wordsInUse-1) +
		(bitsPerWord - numberOfLeadingZeros(b.words[b.wordsInUse-1]))
}

func (b *Bitset) Size() int {
	return len(b.words) * bitsPerWord
}

func (b *Bitset) Clear(bitIdx int) {
	wordIdx := wordIndex(bitIdx)
	if wordIdx >= b.wordsInUse {
		return
	}

	b.words[wordIdx] &= ^(1 << bitIdx)
	b.recalculateWordsInUse()
	b.checkInvariants()
}

func (b *Bitset) ClearRange(from, to int) {
	checkRange(from, to)

	if from == to {
		return
	}

	startWordIdx := wordIndex(from)
	if startWordIdx >= b.wordsInUse {
		return
	}

	endWordIdx := wordIndex(to - 1)
	if endWordIdx > b.wordsInUse {
		to = b.Length()
		endWordIdx = b.wordsInUse - 1
	}

	firstWordMask := wordMask << from
	lastWordMask := wordMask >> -to

	if startWordIdx == endWordIdx {
		b.words[startWordIdx] &= ^(firstWordMask & lastWordMask)
	} else {
		b.words[startWordIdx] &= ^firstWordMask
		for i := startWordIdx + 1; i < endWordIdx; i++ {
			b.words[i] = 0
		}
		b.words[endWordIdx] &= ^lastWordMask
	}

	b.recalculateWordsInUse()
	b.checkInvariants()
}

func (b *Bitset) ClearAll() {
	for b.wordsInUse > 0 {
		b.wordsInUse--
		b.words[b.wordsInUse] = 0
	}
}

func (b *Bitset) Flip(idx int) {
	checkBitIdx(idx)

	wordIdx := wordIndex(idx)
	b.expandTo(wordIdx)

	b.words[wordIdx] ^= 1 << idx

	b.recalculateWordsInUse()
	b.checkInvariants()
}

func (b *Bitset) FlipRange(from, to int) {
	checkRange(from, to)
	if from == to {
		return
	}

	startWordIdx := wordIndex(from)
	endWordIdx := wordIndex(to - 1)
	b.expandTo(endWordIdx)

	firstWordMask := wordMask << from
	lastWordMask := wordMask >> -to

	if startWordIdx == endWordIdx {
		b.words[startWordIdx] ^= firstWordMask & lastWordMask
	} else {
		b.words[startWordIdx] ^= firstWordMask

		for i := startWordIdx + 1; i < endWordIdx; i++ {
			b.words[i] ^= wordMask
		}

		b.words[endWordIdx] ^= lastWordMask
	}

	b.recalculateWordsInUse()
	b.checkInvariants()
}

func (b *Bitset) Set(idx int) {
	checkBitIdx(idx)

	wordIdx := wordIndex(idx)
	b.expandTo(wordIdx)

	b.words[wordIdx] |= 1 << idx

	b.checkInvariants()
}

func (b *Bitset) SetRange(from, to int) {
	checkRange(from, to)

	if from == to {
		return
	}

	startWordIdx := wordIndex(from)
	endWordIdx := wordIndex(to - 1)

	b.expandTo(endWordIdx)

	firstWordMask := wordMask << from
	lastWordMask := wordMask >> -to
	if startWordIdx == endWordIdx {
		b.words[startWordIdx] |= firstWordMask & lastWordMask
	} else {
		b.words[startWordIdx] |= firstWordMask

		for i := startWordIdx + 1; i < endWordIdx; i++ {
			b.words[i] = wordMask
		}

		b.words[endWordIdx] |= lastWordMask
	}

	b.checkInvariants()
}

func (b *Bitset) Get(idx int) bool {
	checkBitIdx(idx)
	b.checkInvariants()

	wordIdx := wordIndex(idx)
	return wordIdx < b.wordsInUse && (b.words[wordIdx]&(1<<idx)) != 0
}

func (b *Bitset) NextClearBit(from int) int {
	checkBitIdx(from)

	b.checkInvariants()

	u := wordIndex(from)
	if u >= b.wordsInUse {
		return from
	}

	word := ^b.words[u] & (wordMask << from)

	for {
		if word != 0 {
			return u*bitsPerWord + numberOfTrailingZeros(word)
		}
		u++
		if u == b.wordsInUse {
			return b.wordsInUse * bitsPerWord
		}
		word = b.words[u]
	}
}

func (b *Bitset) NextSetBit(from int) int {
	checkBitIdx(from)

	b.checkInvariants()

	u := wordIndex(from)
	if u >= b.wordsInUse {
		return -1
	}

	word := b.words[u] & (wordMask << from)

	for {
		if word != 0 {
			return u*bitsPerWord + numberOfTrailingZeros(word)
		}
		u++
		if u == b.wordsInUse {
			return -1
		}
		word = b.words[u]
	}
}

func (b *Bitset) PreviousClearBit(from int) int {
	checkBitIdx(from)

	b.checkInvariants()

	u := wordIndex(from)
	if u >= b.wordsInUse {
		return from
	}

	word := ^b.words[u] & (wordMask >> -(from + 1))

	for {
		if word != 0 {
			return (u+1)*bitsPerWord - 1 - numberOfLeadingZeros(word)
		}
		if u == 0 {
			return -1
		}
		u--
		word = ^b.words[u]
	}
}

func (b *Bitset) PreviousSetBit(from int) int {
	checkBitIdx(from)

	b.checkInvariants()

	u := wordIndex(from)
	if u >= b.wordsInUse {
		return b.Length() - 1
	}

	word := b.words[u] & (wordMask >> -(from + 1))

	for {
		if word != 0 {
			return (u+1)*bitsPerWord - 1 - numberOfLeadingZeros(word)
		}
		if u == 0 {
			return -1
		}
		u--
		word = b.words[u]
	}
}

func (b *Bitset) ToByteSlice() []byte {
	n := b.wordsInUse
	if n == 0 {
		return nil
	}

	l := 8 * (n - 1)
	for x := b.words[n-1]; x != 0; x >>= 8 {
		l++
	}

	res := bytes.NewBuffer(make([]byte, l))
	for i := 0; i < n-1; i++ {
		word := b.words[i]

		for j := 0; j < bitsPerWord/8; j++ {
			res.WriteByte(byte(word & 0xff))
			word >>= 8
		}
	}

	for word := b.words[n-1]; word != 0; word >>= 8 {
		res.WriteByte(byte(word & 0xff))
	}

	return res.Bytes()
}

func (b *Bitset) initWords(nbits int) {
	b.words = make([]uint64, wordIndex(nbits-1)+1)
}

func (b *Bitset) expandTo(wordIdx int) {
	if b.wordsInUse < wordIdx {
		b.wordsInUse = wordIdx + 1
	}
}

func (b *Bitset) recalculateWordsInUse() {
	var i int
	for i = b.wordsInUse - 1; i >= 0; i-- {
		if b.words[i] != 0 {
			break
		}
	}
	b.wordsInUse = i + 1
}

func (b *Bitset) checkInvariants() {

}

func wordIndex(bitIdx int) int {
	return bitIdx >> addressBitsPerWord
}

func checkBitIdx(idx int) {
	if idx < 0 {
		panic(fmt.Errorf("index < 0: %d", idx))
	}
}

func checkRange(fromIdx, toIdx int) {
	if fromIdx < 0 {
		panic(fmt.Errorf("fromIndex < 0: %d", fromIdx))
	}
	if toIdx < 0 {
		panic(fmt.Errorf("toIndex < 0: %d", toIdx))
	}
	if fromIdx > toIdx {
		panic(fmt.Errorf("fromIndex: %d > toIndex: %d", fromIdx, toIdx))
	}
}

func numberOfLeadingZeros(l uint64) int {
	intNumberOfLeadingZeros := func(i uint32) int {
		n := 31
		if i >= 1<<16 {
			n -= 16
			i >>= 16
		}
		if i >= 1<<8 {
			n -= 8
			i >>= 8
		}
		if i >= 1<<4 {
			n -= 4
			i >>= 4
		}
		if i >= 1<<2 {
			i -= 2
			i >>= 2
		}
		return int(uint32(n) - (i >> 1))
	}

	x := uint32(l >> 32)
	if x == 0 {
		return 32 + intNumberOfLeadingZeros(uint32(l>>32))
	}
	return intNumberOfLeadingZeros(uint32(l))
}

func numberOfTrailingZeros(l uint64) int {
	intNumberOfTrailingZeros := func(i uint32) int {
		i = ^i & (i - 1)

		n := 1
		if i > 1<<16 {
			n += 16
			i >>= 16
		}
		if i > 1<<8 {
			n += 8
			i >>= 8
		}
		if i > 1<<4 {
			n += 4
			i >>= 4
		}
		if i > 1<<2 {
			n += 2
			i >>= 2
		}
		return int(uint32(n) + (i >> 1))
	}

	x := uint32(l)
	if x == 0 {
		return 32 + intNumberOfTrailingZeros(uint32(l>>32))
	}
	return intNumberOfTrailingZeros(x)
}
