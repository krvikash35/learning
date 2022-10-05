* With java10+ we can use `var` keyword

## String
```
char[] ch={'a','b','c'};  
String s=new String(ch);

String s="abc";
String s2 = "abc" // no new object, reuse existing memory allocation
String s=new String("abc"); //this will always create new string in normal(non-heap) heap memory.
```

## Array
```
int data[] = {1,2};
int[] data = {1,2};

int[] data;
data = new int[] {1, 2};
```

## Reference vs Value compare
```
String a="1"
String b="1"
a==b        //False
a.equals(b) //True
```