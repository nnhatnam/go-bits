package bits



//Compute floor( (x + y ) / 2 ). FloorAverage will give the correct value even if in case (x + y ) is overflow
func FloorAverage(x, y uint) uint{
	return (x & y) + ( (x ^ y) >> 1 )
}

//Compute previous even value
func PrevEven(x uint) uint { return (x - 1) &^ 1 }

//Compute next even value
func NextEven(x uint) uint { return (x | 1) + 1 }


//Compute previous odd value
func PrevOdd(x uint) uint { return (x &^ 1) - 1 }

//Compute next odd value
func NextOdd(x uint) uint { return (x + 1) | 1 }

//Return zero if bit[i] is zero, else return a number with only 1 bit set at index i
func TestBit(x uint, i int) uint {
	return x & (1 << i)
}

//Turn on bit at index i
func SetBit(x uint, i int) uint {
	return x | (1 << i)
}

//Turn off bit at index i
func ClearBit(x uint, i int) uint {
	return x &^ (1 << i)
}

//flip bit at index i
func ChangeBit(x uint, i int) uint {
	return x ^ (1 << i)
}

//copy bit at position src to position dst
func CopyBit(x uint, isrc, idst int) uint {
	//x & 1 = 00000x[0]
	var temp uint = ((x >> isrc) ^ (x >> idst)) & 1
	return x ^ (temp << idst)
}

//copy bit according at src-mark (msrc) to the bit
//according to the desk-mask (mdst)
//Both msrc and mdst must have exactly one bit set
//func MaskCopyBit(x, msrc, mdst uint) uint {
//
//}

//Return an unint number with bit at positions i and j swapped
func SwapBit(x uint, i , j int) uint {
	temp := ((x >> i) ^ (x >> j)) & 1
	x ^= temp << i
	x ^= temp << j
	return x
}


// Return word where only the lowest set bit in x is set.
// Return 0 if no bit is set.
func LowestOne(x uint) uint {
	return x & -x
}


// Return word where only the lowest unset bit in x is set.
// Return 0 if all bits are set.
func LowestZero(x uint) uint {
	return LowestOne(^x)
}

//Turn off right most bit
func ClearLowestOne(x uint) uint{
	return x & (x - 1)
}

//Turn on right most zero
func SetLowestZero(x uint) uint {
	return x | (x + 1)
}

//func asm_bsf


// Return word where all the (low end) ones are set.
// Example: 01011011 --> 00000011
// Return 0 if lowest bit is zero:
// 10110110 --> 0
func LowOnes(x uint) uint {
	x = ^x
	x &= -x
	x -= 1
	return x
}


// Return word where all the (low end) zeros are set.
// Example: 01011000 --> 00000111
// Return 0 if all bits are set.
func LowZeros(x uint) uint {
	x &= -x
	x -= 1
	return x
}

// Isolate lowest block of ones.
func LowestBlock(x uint) uint {
	l := x & -x
	y := x + l
	x ^= y
	return x & (x >> 1)
}

//Let x = ∗0W and y = ∗1W, the following function computes W
func LowMatch(x, y uint) uint {
	x ^= y
	x &= ^x
	x -= 1
	x &= y
	return x
}

// Return word with length-n bit block starting at bit p set.
// Both p and n are effectively taken modulo BITS_PER_LONG.
func BitBlock(p, n uint) uint {
	var x uint = (1 << n) - 1
	return x << p
}

// Return word with only the isolated ones of x set.
func SingleOnes(x uint) uint {
	return x &^ (( x << 1) | (x >> 1))
}

// Return word with only the isolated zeros of x set.
func SingleZerosXI(x uint) uint {
	return SingleOnes(^x)
}

// Return word with only the isolated zeros of x set.
func SingleZeros(x uint) uint {
	return ^x & ((x << 1 ) & (x >> 1))
}

//// Return word where only the isolated ones and zeros of x are set.
func SingleValues(x uint) uint {
	return (x ^ (x << 1 )) & (x ^ (x >> 1))
}

// Return word where only the isolated ones and zeros of x are set.
func SingleValuesXI(x uint) uint {
	return SingleOnes(x) | SingleZeros(x)
}

// Return word where only those ones of x are set that lie next to a zero.
func BorderOnes(x uint) uint {
	return x &^ ((x << 1) & (x >> 1))
}

// Return word where those bits of x are set that lie on a transition.
func BorderValues(x uint) uint {
	return ( x ^ ( x << 1 )) | 	( x ^ (x >> 1))
}

// Return word where only those ones of x are set
// that lie right to (i.e. in the next lower bin of) a zero.
func HighBorderOnes(x uint) uint {
	return x & (x ^ ( x >> 1))
}

func LowBorderOnes(x uint) uint {
	return x & ( x ^ (x<<1) )
}

// Return word where only those ones of x are set
// that are at the border of a block of at least 2 bits.
func BlockBorderOnes(x uint) uint {
	return x & ((x << 1) ^ (x >> 1))
}

// Return word where only those bits of x are set
// that are at left of a border of a block of at least 2 bits.
func LowBlockBorderOnes(x uint) uint{
	var t = x & ( (x<<1) ^ (x>>1) )
	return t & (x>>1)
}

// Return word where only those bits of x are set
// that are at right of a border of a block of at least 2 bits.
func HighBlockBorderOnes(x uint) uint {
	t := x & ( (x<<1) ^ (x>>1) )
	return t & (x<<1)
}

// Return word where only those bits of x are set
// that are part of a block of at least 2 bits
func BlockOnes(x uint) uint {
	return x & ( (x<<1) | (x>>1) )
}

// Return word where all the (high end) zeros are set.
// e.g.: 00011001 --> 11100000
// Returns 0 if highest bit is set:
// 11011001 --> 00000000
//func HighZeros()


// Return floor(log2(x)),
// i.e. return k so that 2^k <= x < 2^(k+1)
// If x==0, then 0 is returned (!)
func  Ld(x uint) uint {
	var k uint = 0
	for ; x >= 0 ; x >>= 1 {
		k++
	}
	return k
}

// Return whether x \in {1,2,4,8,16,...}
func OneBitQ(x uint) bool {
	var m uint = x - 1
	return (((x^m)>>1) == m)
}

// Return whether x == 0(!) or x == 2**k
func IsPowerOf2(x uint) bool {
	return (x & (x-1)) == 0
}

// Return x if x=2**k
// else return 2**ceil(log_2(x))
// Exception: returns 0 for x==0
//func NextPowOf2(x uint) uint {
//
//}

// Return whether the set bits of u are a subset of the set bits of e.
// That is, as bitsets, test whether u is a subset of e.
func IsSubset(u , e uint) bool {
	return  (u & e)==u
}

// Return whether u (as bitset) is a proper subset of e.
func IsProperSubset(u , e uint) bool {
	return  (u<e) && ((u & e)==u)
}