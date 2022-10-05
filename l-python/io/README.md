```python
f = open("demofile.txt", "rt") //r-read, a-append, w-write, x-create, t-text, b-binary
print(f.read())
print(f.read(5)) //first 5 char

f = open("demofile2.txt", "a")
f.write("Now the file has more content!")
f.close()

import os
if os.path.exists("demofile.txt"):
  os.remove("demofile.txt")
else:
  print("The file does not exist")
  
os.rmdir("myfolder")
```