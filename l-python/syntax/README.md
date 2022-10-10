
* golable var:  Variables that are created outside of a function 

```python
x = y = z = "Orange" // one value to multiple variable
x, y, z = "Orange", "Banana", "Cherry" //many value to multiple variable

casting
x = str(3), y = int(3), z = float(3) #casted to '3', 3, 3.0

fruits = ["apple", "banana", "cherry"] 
x, y, z = fruits //unpack collection like list and tupple

print(x + y + z) //will concate/math if all str/num else error onmix, use comma for output

if b > a:
  print("b is greater than a")
  
i = 1
while i < 6:
  print(i)
  i += 1
  
fruits = ["apple", "banana", "cherry"]
for x in fruits:
  if x == "banana":
    break
  print(x)
  
def my_function():
  print("Hello from a function")
  
x = lambda a, b, c : a + b + c
print(x(5, 6, 2))

class Person:
  def __init__(self, name, age):
    self.name = name
    self.age = age

p1 = Person("John", 36)

class Student(Person):
  pass
  

mytuple = ("apple", "banana", "cherry")
myit = iter(mytuple)
print(next(myit))
```

## **kwargs vs *argv
```
def myFun(arg1, *argv):
    print("First argument :", arg1)
    for arg in argv:
        print("Next argument through *argv :", arg)
myFun('Hello', 'Welcome', 'to', 'GeeksforGeeks')


def myFun(**kwargs):
    for key, value in kwargs.items():
        print("%s == %s" % (key, value))
myFun(first='Geeks', mid='for', last='Geeks')

def myFun(arg1, arg2, arg3):
    print("arg1:", arg1)
    print("arg2:", arg2)
    print("arg3:", arg3)
 
args = ("Geeks", "for", "Geeks")
myFun(*args)
 
kwargs = {"arg1": "Geeks", "arg2": "for", "arg3": "Geeks"}
myFun(**kwargs)
```