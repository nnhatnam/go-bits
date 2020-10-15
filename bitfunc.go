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
func CopyBit(x uint, src, dst int) uint {
	//x & 1 = 00000x[0]
	var temp uint = ((x >> src) ^ (x >> dst)) & 1
	return x ^ (temp << dst)
}
