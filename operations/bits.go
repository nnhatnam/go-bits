package operations

//Turn off the rightmost 1-bit in x, return 0 if x = 0
func TurnOffRightMostOne(x uint) uint {
	return x & (x - 1)
}

//Turn off the rightmost 0-bit in x, return all 1's if none
func TurnOnRightMostZero(x uint) uint {
	return x | (x + 1)
}

//Turn off the trailing 1’s in x, return x if none
func TurnOffTrailingOne(x uint) uint {
	return x & (x + 1)
}

//Turn on the trailing 0’s in x, return x if none
func TurnOnTrailingZero(x uint) uint {
	return x | (x - 1)
}

//create an uint with a single 1-bit at the position of the
//rightmost 0-bit in x, return 0 if none
func SingleOneAtRightMostZero(x uint) uint  {
	return ^x & (x + 1)
}

//Create an uint with a single 0-bit at the position of the
//rightmost 1-bit in x, return all 1’s if none
func SingleZeroAtRightMostOne(x uint) uint  {
	return ^x & (x + 1)
}