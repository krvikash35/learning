
## access modifier
```
package p1;
class C1

package p2;
class C2
class C3 extends C1
```

* `private`: method can be accessed within same class only.
* `default`: can be accessed within same pkg only.
* `protected`: can be access within other pkg but by subclass only. C3 can access C1's protected method.
* `public`: method can be accessed within other pkg as well.