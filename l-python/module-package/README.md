## Module

* Any file in python is module. You can seperate your code in multiple files i.e multiple module.
* A module can be imported into other module or main module
* Within a module, the module’s name is available as the value of the global variable __name__
standard python module like sys


```python
# fibo.py
def fib(n):    # write Fibonacci series up to n
    print("fib", n)

def fib2(n):   # return Fibonacci series up to n
   print("fib2", n)
   
#dummy.py
import fibo 
fibo.fib(1000)
fibo.fib2(100)

from fibo import fib, fib2
from fibo import * #export all except _. not preferred
import fibo as fib
from fibo import fib as fibonacci
```

* module can be run as script. It will set __name__ to __main__
* file can be used as a script as well as an importable module, because the code that parses the command line only runs if the module is executed as the “main” file:

```python
python fibo.py <arguments>
if __name__ == "__main__":
    import sys
    fib(int(sys.argv[1]))
```

## Package
* Packages are a way of structuring Python’s module namespace by using “dotted module names”. For example, the module name A.B designates a submodule named B in a package named A
* The __init__.py files are required to make Python treat directories containing the file as packages

```
sound/                          Top-level package
      __init__.py               Initialize the sound package
      formats/                  Subpackage for file format conversions
              __init__.py
              wavread.py
              wavwrite.py
      effects/                  Subpackage for sound effects
              __init__.py
              echo.py
              surround.py
      filters/                  Subpackage for filters
              __init__.py
              equalizer.py
```

```python
import sound.effects.echo
sound.effects.echo.echofilter(input, output, delay=0.7, atten=4)

from sound.effects import echo
echo.echofilter(input, output, delay=0.7, atten=4)

from sound.effects.echo import echofilter
echofilter(input, output, delay=0.7, atten=4)

from sound.effects import * #bad idea, it will import if __init__.py define __all__
__all__ = ["echo", "surround", "reverse"]
```