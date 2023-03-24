# About
Flexible implementation of static(without erase)/dynamic hash table  
May be useful in high load projects as faster analogue of builtin map or as addressable map and etc
# About implementation
- I am using open addressing variant, but i think i will add chains variant
- To find cell i am using 2 independent hash functions to avoid clusterization, you may use custom gap function(h2), but it is bad idea to use custom itit function(h1)
- To store values i am using interface that generaly can store different types in same table (yes it is some kind of dynamic typization), but i think that it is ub because of oppotunity of clusterization
- I am doing rehash after reaching load factor >= 0.5, so total comlexy of rehashing O(nlogn) but in particular O(n) n - final size
- Dynamic hash table = static hash table with tombstones, so static is faster in 1.5-2 times than dynamic
