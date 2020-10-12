# go-bits

### Built-in 

**Bitwise operators that Go supports**
```
 &   bitwise AND
 |   bitwise OR
 ^   bitwise XOR
&^   AND NOT
<<   left shift
>>   right shift
```

**Go's integer is represented using two's complement arithmetic, according to its [spec](https://golang.org/ref/spec#Numeric_types)**

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

### Pitfall

**Absolute Number**  
We CAN'T just simply write absolute function like this 
```
if ( x < 0 ) x = -x;   // WRONG!
```
because in two's complement zero is not the only number that is equal to its negative. The value with just the highest 
bit set (the most negative value) also has this property. For example, the most negative value of int16 is -32768. We've 
got
```go
var int16 x = -32768 
fmt.Println(x == -x)             // TRUE!!! , -x is -32768 also  
```  
