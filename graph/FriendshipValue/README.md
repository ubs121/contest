First of all we know the graph we are given is comprised of several components. Each component has its own size. At first if a component has  nodes, you just need to add  edges to make the component connected (a subtree of the component), add the rest of the edges at the end when all the components are connected in themselves. 

At the end, when all of the components are connected, add the extra edges. 

But what about the order of the components? Its better to add larger components first so that larger numbers are repeated more. 
What about a component in itself? Try to make a tree from that component. 

This is a sample of a graph with 15 nodes with 3 components of sizes 7, 5, 3 with 10, 8 and 3 edges respectively. 

First component is getting connected
```
2   1   0   0   0   0   0       0   0   0   0   0       0   0   0
2   2   2   0   0   0   0       0   0   0   0   0       0   0   0
3   3   3   3   0   0   0       0   0   0   0   0       0   0   0
4   4   4   4   4   0   0       0   0   0   0   0       0   0   0
5   5   5   5   5   5   0       0   0   0   0   0       0   0   0
6   6   6   6   6   6   6       0   0   0   0   0       0   0   0
```
6 edges of 10 edges in component 1 are used to make it connected


First component is connected and second component is getting connected
```
6   6   6   6   6   6   6       1   1   0   0   0       0   0   0
6   6   6   6   6   6   6       2   2   2   0   0       0   0   0
6   6   6   6   6   6   6       3   3   3   3   0       0   0   0
6   6   6   6   6   6   6       4   4   4   4   4       0   0   0
```
4 edges of 8 edges in component 2 are used to make it connected

First and second component are connected, third component is getting connected
```
6   6   6   6   6   6   6       4   4   4   4   4       1   1   0
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
```
2 edges of 3 edges in component 3 are used to make it connected

now that all three components are connected we will add extra edges 
(4 from the first, 4 from the second, 1 from the third)

```
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2

6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
6   6   6   6   6   6   6       4   4   4   4   4       2   2   2

6   6   6   6   6   6   6       4   4   4   4   4       2   2   2
```