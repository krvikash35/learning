
* data types: print(type(x))
    * text: str x = "Hello World"
    * number: 
        * int x = 20
        * float x = 20.5
        * complex x = 1j
    * sequence:
        * list x = ["apple", "banana", "cherry"]
        * tuple x = ("apple", "banana", "cherry")
        * range x = range(6)
    * mapping: dict x = {"name" : "John", "age" : 36}
    * set:
        * set x = {"apple", "banana", "cherry"}
        * frozenset x = frozenset({"apple", "banana", "cherry"})
    * bool: bool x = True
    * binary:
        * bytes x = b"Hello"
        * bytearray x = bytearray(5)
        * memoryview memoryview(bytes(5))
    * None: NoneType x = None


string
```python
b = "Hello, World!"
print(b[2:5]) #slice from 2 to 5(exclude)
print(b[:5]) #slice from start to 5(exclude)

quantity,itemno,price  = 3, 567, 49.95
myorder = "I want {} pieces of item {} for {} dollars."
print(myorder.format(quantity, itemno, price))

txt = "hello \"Vikings\" " //esacpe

capitalize()
strip()
```

list
```
x = ["apple", "banana", "cherry"]
x[1] first item 
x[-1] last item
x[2:5] index 2-5, exclude 5

x[1]="hello"
x[1:3] = ["blackcurrant", "watermelon"] //replace index 1 & 2
x[1:2] = ["blackcurrant", "watermelon"] //replace index 1 and insert index 2
x.insert(2, "watermelon") // insert at index 2 without replacing
x.append("orange") //add item to last
x.pop(1) // remove index
x.pop() // remove last
del x[1] // remove index
del x // delete list completely
x.remove("banana") //remove item
x.clear() //empty list

for x in thislist:
  print(x)
  
for i in range(len(thislist)):
  print(thislist[i]
 
i = 0 
while i < len(x):
  print(x[i])
  i = i + 1
  
[print(y) for y in y] //List Comprehension

List comprehension offers a shorter syntax when you want to create a new list 
based on the values of an existing list.
newlist = [x for x in fruits if "a" in x] //fruits with a in name

x.sort(reverse = True) //descending, default ascending


y=x.copy
y=list(x)

z=x+y
z=x.extend(y)

count()
index()

```


tupple
```
#tuple are immutable but workaround is conver to list, update then convert back totuple
x = ("apple", "banana", "cherry")
y = list(x)
y[1] = "kiwi"
x = tuple(y)
```

sets
```
for x in {"v", "i"}:
exist = "v" in {"v", "i"}

x.add("a")
x.update(list/tupple/set)

x.remove("a")
x.discard("a")
x.pop() //remove last item, since unordered so not sure which one gets removed
x.clear()
del x

x.union(y) //new set with items from both set

issubset()
issuperset()
```

dict
```
x = {
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
x.get("model")
x.keys() // keys get updated if dict gets updated as it is reference

x["year"] = 2018
x.update({"year": 2020}) //upsert

x.pop("model") //remove the item with key
del x["model"]
del x
x.clear() //empty dict

for key in x:
  print(key, x[key])
  
for value in x.values():
  print(value)
  
for key in x.keys():
  print(key)
  
for key,value in x.items():
  print(key)
  
y = x.copy()
y = dict(x)

fromkeys()
setdefault()
```

## list vs tuple
* tuple are fixed size and hence immutable while list are dynamic and mutable.
* can't add/remove element from tupple. no method as append, extend, remove, pop