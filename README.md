# go-bits

### Built-in 

**Bitwise operators that Go supports**
```
 &   bitwise AND
 |   bitwise OR
 ^   bitwise XOR (0011 ^ 0101)
 ^   bitwise NOT (^0101)
&^   Bitclear (AND NOT)
<<   left shift
>>   right shift
```

**Go's integer is represented using two's complement arithmetic, more detail at [spec](https://golang.org/ref/spec#Numeric_types)**

### Shift and Division

Use a shift instruction for division with unsigned types. Use it with signed integer may cause undefined 
  behavior such as **-1 >> 1 = -1**, but **-1/2 = 0**
```go
a := -1
fmt.Println(a >> 1)         // -1 
fmt.Println(a/2)            //  0
```

Computing remainders modulo x, where x is unsigned types and a power of 2, is equivalent to **&(x-1)**
```go
var x uint = 32                     // 2^5
fmt.Println( a % x == a & (x - 1))  // true 
```

### Condition 

**Check whether at least a or b equal zero**
```
if( !( a & b ) )    // ==   if ( a == 0 || b == 0 )
```

**Check whether several variables is zero**
```
if ( (a|b|c|...|z) == 0 )   // == if ( (a == 0)  && (b == 0) && (c == 0) ... && (z == 0) )
``` 

**Check whether exactly one of two variables is zero**
```
if ( (^a) ^ (^b) )         // == if ( ( a == 0 && b != 0 ) || ( a != 0 && b == 0) )
```

### Pitfall

**Absolute Number**  
We CAN'T just simply write absolute function like this 
```
if ( x < 0 ) x = -x;  
```
because in two's complement zero is not the only number that is equal to its negative. The value with just the highest 
bit set (the minimum negative value) also has this property. For example, the most minimum value of int16 is -32768. Then we 
got
```go
var int16 INT16_MIN = -32768 
fmt.Println(INT16_MIN == -INT16_MIN)             // print true
```  
So, to avoid that, we need to handle that case. Either overflow panic or point out in the document.