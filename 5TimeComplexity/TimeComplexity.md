# Time Complexity

The time complexity of an algorithm uses a rough estimate of the execution time needed, which can be used to estimate `the upper bound of the program's running time`.

## Comparing Algorithms

When comparing algorithms, first compare `the communication complexity` and `I/O complexity`, then `the time complexity`, and finally `the space complexity`, because communication and I/O have very high costs.

先比较 `1 通信复杂度` 和 `2 IO复杂度`,然后是 `3时间复杂度` ,最后是 `4空间复杂度`

## Time Complexity List

Below is a comparison of the time complexities of different programs:

### Three Nested Loops

Many say the time complexity of three nested loops is `O(n^3) = n1 * n2 * n3 = n * n * n`.

But this `assumes each loop iteration is proportional to n`. (每层要和 n 成正比才行)

If the first loop iterates 10 times, the time complexity becomes 
` O(n^2) = n1 * n2 * n3 = 10 * n * n = 10 * n^2`

#### Time complexity of three nested loops:

Let T(n) be the number of executions of the three loops, and n be the input size.

##### Base Case：

When `n = 1`, each loop executes 1 time, so `T(1) = 1` holds.

Also assume for any k, `T(k) ≤ k^3 holds when n = k`.

##### Inductive step：

When `n = k+1`, the outer loop executes k+1 times, middle loop executes k+1 times, and inner loop executes k+1 times.

So `T(k+1) ≤ (k+1)^3` also holds.

##### Induction conclusion：

Therefore, `for any n, T(n) ≤ n^3 holds`, so the time complexity of three nested loops is `O(n^3)`.



https://zhuanlan.zhihu.com/p/320419705 
