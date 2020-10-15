package readable




//Compute floor( (x + y ) / 2 ). FloorAverage will give the correct value even if in case (x + y ) is overflow
func FloorAverage(x, y uint) uint{
	return (x & y) + ( (x ^ y) >> 1 )
}

// x | y is a way to turn certain bits of x on based on a mask constant y, then x &^ y is doing the the opposite
// and turn those same bits off. When y = 1, the operation only impact the last bit. x | 1 => turn last bit on,
// and x &^ i => turn last bit off. We can use this characteristic to find the closet odd number >= x using x | 1, and
//find the closet even number <= x using x &^ 1. Then we've got two function below

func findOddNumGreaterOrEqual(x uint) uint {
	return x | 1
}

func findEvenNumLessOrEqual(x uint) uint {
	return x | 1
}

//Compute previous even value. The algorithm is find the closet even number <= (x -1)
func PrevEven(x uint) uint { return findEvenNumLessOrEqual(x - 1) }

//Compute next even value. The algorithm is find the closet odd number >= x then plus 1
func NextEven(x uint) uint { return findOddNumGreaterOrEqual(x) + 1 }


//Compute previous odd value. The algorithm is find the closet even number <= x then minus 1
func PrevOdd(x uint) uint { return findEvenNumLessOrEqual(x) - 1 }

//Compute next odd value. The algorithm is increase x by 1 then find the closet odd number
func NextOdd(x uint) uint { return findOddNumGreaterOrEqual(x + 1) }

