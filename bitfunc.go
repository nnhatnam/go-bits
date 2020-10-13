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