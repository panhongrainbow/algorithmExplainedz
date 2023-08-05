# Operation of Max Heap

## Parent and Child Nodes

A maxheap has `all nodes within a slice`, and each node has `a parent-child relationship`.

The parent node's index can be used to deduce the index of its two child nodes.

The formula is as follows:
$$
formula\ 1:\ index_{left\ child\ node} = 2 \times index_{parent\ node} + 1 \\
formula\ 2:\ index_{left\ child\ node} = 2 \times index_{parent\ node} + 2 \\
$$
For Example:

If the index of the parent node is 0, substituting into formulas 1 and 2 yields 1 and 2, respectively, This corresponds to the illustration below.

<img src="../assets/image-20230803005118736.png" alt="image-20230803005118736" style="zoom: 33%;" /> 

Why do we `multiply by 2 and add 1 or 2` in those formulas?

## Explanation of the Formula

Multiplying by 2 is used to locate the index, while `adding 1 or 2` serves as `an offset`.

(乘以2是索引，加1 加2是偏移量)



By observing `the circled section` in the diagram below, it's evident that `the child nodes consistently increase by a factor of 2`.

(圈起来部分，子节点2个等距增加，乘2合理)



Therefore, multiplying by 2  Therefore, multiplying by 2 is a logical choice.<img src="../assets/image-20230803104935465.png" alt="image-20230803104935465" style="zoom:100%;" />

## Max Value & DownShift

> Because the goal is to ensure that `the maximum value is in the parent node`, the process of extracting the maximum value involves a `downward shift`.



The entire process of selecting the maximum value and shifting downwards.

---

In the initial state, `there are five nodes`, with `node1 being the maximum value` and `assuming node5 to be the minimum value`.



Here, node1 is the maximum value, and node5 is the `assumed` minimum value.

`The minimum value` of the max heap `may not necessarily be at the end of the slice.`

The heap size is `5.`

 

<img src="../assets/image-20230804023152929.png" alt="image-20230804023152929" style="zoom:70%;" /> 

---

When operating on a max heap, it's necessary to first `swap node1 and node5.`

This is because we want to `move the maximum value, node5, to the end of the slice.`

(最小值并不一定会在切片的取后面)

<img src="../assets/image-20230804023343239.png" alt="image-20230804023343239" style="zoom:67%;" /> 

---

 Retrieve `the maximum value, node1`, from the max heap.

Later, the entire max heap will need to be restructured.

<img src="../assets/image-20230804024316602.png" alt="image-20230804024316602" style="zoom:100%;" /> 

---

Since node2 and node3 are greater than node5, node5 needs to be moved down to the position of either node2 or node3.

However, `the root of the max heap` must hold `the maximum value`.

Therefore, nodes 2 and 5 are swapped.

Next, node5 will be adjacent to node4.

As node4 is smaller than node5, the downward movement of node5 stops.

The restructuring of the max heap is now complete.

<img src="../assets/image-20230804031519974.png" alt="image-20230804031519974" style="zoom:100%;" />

## Append Value & UpShift

> The process of shifting up in a max heap is relatively simpler compared to the downward operation, not as complicated. (和下滑比较起来，简单多了)



The entire process appends arbitrary values and shifts them upwards.

---

This is the initial state, with a heap size of 5, currently 4 have been used, and there's room `for 1 more to be added`.

<img src="../assets/image-20230805010455589.png" alt="image-20230805010455589" style="zoom:55%;" /> 

---

Adding `a new node, node6`, to the end of the entire slice.

Next, `the entire max heap will need to be reorganized`.

<img src="../assets/image-20230805012016008.png" alt="image-20230805012016008" style="zoom:60%;" /> 

---

Because node6 is larger than node5, nodes node5 and node6 can exchange positions.



Next, node6 is adjacent to node2, but node2 is larger.

Therefore, `node6 cannot continue to move up.`



The reorganization is complete.

<img src="../assets/image-20230805012944563.png" alt="image-20230805012944563" style="zoom:65%;" /> 

