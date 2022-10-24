
types
* Brute force
* Recursive
* backtracking
* searching
* sorting
* hashing
* divide & conquer
* greedy
* dynamic programming
* randomized



complexity
* time: notation: Big-o(O), Omega(Ω), Theta(Θ)
* space


## hashing algo
* plain text + hasing algo --> hashed text
* one way, small hashed text,
* used in password storage, digital sign, 
* example: md5, SHA[-1, -2, -256 etc]

### md5
* It is easy to break but still popular and can be used for data integreity check.
* produce 128bit hash

## sha(secured hashing algo)
* modified version of md5

## recursive algo
* break the problem in two part: base case, recursive step.
* find sum of N natural number

```
func sum(n){
    if n == 0:
        return 0
    else
        return n + sum(n-1)
}
```

## Backtracking
* uses a brute force approach for finding the desired output.
* if current solution is not suitable then backtrack and try other solution, recursion is also used.
* for optimal solution, use dynamic programming.
* problems: N queen, all permutation of string, Hamiltonian path in graph, knight tour problem.

## Searching
* type: linear, binary, jump, exponential, fibonacci

## sorting
* type: bubble, selection, insertion, merge, quicksort, counting, radix, bucket, heap, shell
```
Sorting         T-Best	T-Worst	T-Avg	Space
Bubble Sort	    n	    n2	    n2	    1
Selection Sort	n2	    n2	    n2	    1
Insertion Sort	n	    n2	    n2	    1
Merge Sort	    nlogn	nlogn	nlogn	n
Quicksort	    nlogn	n2	    nlogn	logn
Counting Sort	n+k	    n+k	    n+k	    max
Radix Sort	    n+k	    n+k	    n+k	    max
Bucket Sort	    n+k	    n2	    n	    n+k
Heap Sort	    nlogn	nlogn	nlogn	1
Shell Sort	    nlogn	n2	    nlogn	1
```

## Divide and conquer
* breaks a problem into subproblems that are similar to the original problem, recursively solves the subproblems, and finally combines the solutions to the subproblems to solve the original problem.
* it uses recursion
* both merge & quick sort uses this.

## greedy
* solving a problem by selecting the best option available at the moment. It doesn't worry whether the current best result will bring the overall optimal result.
* doesn't always produce the optimal solution.
* used by selection sort, Huffman coding

## dynamic programming
* Optimization over recursion.
* The idea is to simply store the results of subproblems, so that we do not have to re-compute them when needed later.This simple optimization reduces time complexities from exponential to polynomial.
* example: find the fibonacci sequence upto 5th term
## randomized