Clojure

https://www.braveclojure.com/clojure-for-the-brave-and-true/

* Def 
* Defn
* defn-
* ->
* condp
* If
* Else
* Fn
* Map




(def t0 (System/currentTimeMillis))
(defn t1 [] (System/currentTimeMillis))
(t1)
;; => 1318408717941
t0
;; => 1318408644243
t0
;; => 1318408644243
(t1)
;; => 1318408719361


If no arg then don’t use def. If no arg with def, then it bind the value i.e evaluated only once.

* defn = public
* defn- = private

Overloaded 
————
(defn tests 
  ([] 0) 
  ([a b] 1))


(-> expr f1 f2 f3)  ;same as (f3 (f2 (f1 expr)))


(defn myfunc []
  (if (and cond1 cond2) 
      something
      somethingelse))



(defn myfunc []
  (if cond1
      (if cond2 something newelse)
      somethingelse))


(defn my-func [a]
  (stuff a))

As 

(def my-func
  (fn [a] (stuff a)))

Using just fn creates an anonymous function that alone isn't bound to any symbol externally. It must be bound using let or def to be referred to outside of itself.

When you supply a name for fn, it can't be referred to outside of the function, but it can be used to refer to itself to create a recursive anonymous function:


(defn- hello [name]
  (condp = (name)
   Vikash :hello_vikash
    :hello_guest)



Vector are like arrayList - good for random access
Sequence are like linkedList - good for extending

(list 1 2 3)
(seq [1 2 3]))

(seq (list 1 2 3))


(map inc [1 2 3 4 5])
;;=> (2 3 4 5 6)


