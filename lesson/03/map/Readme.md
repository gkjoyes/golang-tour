# Maps

Maps provide a data structure that allows for the storage and management of key/value pair data.

- Map provides a way to store and retrieve key/value pairs.
- Reading an absent key returns the zero value for the map's value type.
- Iterating over a map is always random.
- The map key must be a value that is comparable.
- Elements in a map are not addressable.
- Maps are a reference type.

## Examples

- [Declare, write, read and delete](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/example1/example1.go)
- [Absent keys](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/example2/example2.go)
- [Map key restrictions](https://play.golang.org/p/43QckRQ82hN)
- [Random behaviour of map iteration](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/example3/example3.go)
- [Sorting maps by keys](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/example4/example4.go)
- [Can't take the address of a map element](https://play.golang.org/p/Ifl6hUa8Adq)
- [Map are reference types](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/example5/example5.go)

## Macro view of the map

[Maps](https://github.com/golang/go/blob/master/src/runtime/map.go#L115) in Go are implemented as a [hash table](https://en.wikipedia.org/wiki/Hash_table). The hash table for a Go map is structured as an array of buckets. The number of buckets is always equal to a power of 2. When performing a map operation such as adding a key/value pair, a hash key is generated against the specified key. The low order bits(LOB) of the generated hash key is used to select a bucket.

![Buckets](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/images/buckets.png)

Once the bucket is selected, the key/value pair needs to be stored, removed or looked up, depending on the type of operation. If we look inside any bucket, we will find two data structures. First, there is an array with the top 8 high order bits (HOB) from the same hash key that was used to select the bucket. This array distinguishes each key/value pair stored in the respective bucket. Second, there is an array of bytes that store the key/value pairs. The byte array packs all the keys and then all the values together for the respective bucket.

![Map Hash Table](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/images/map-hash-table.png)

When you are iterating through the map, the iterator will select [bucket](https://github.com/golang/go/blob/master/src/runtime/map.go#L827) and [offset](https://github.com/golang/go/blob/master/src/runtime/map.go#L832) within that bucket randomly, then return the key/value pairs in the order they are laid out in the byte array. This is why maps are unsorted collections.

### Memory and Bucket Overflow

There is a reason the key/value pairs are packed like this in a single byte array. If the keys and values were stored like key/value/key/value, padding allocations would be needed between each key/value pair to maintain proper [alignment boundaries](https://github.com/gkjoyes/golang-tour/tree/master/lesson/02/syntax/struct-types/padding).

An example where this would apply is with a map that looks like this:

```go
    map[int64]int8
```

The 1-byte value in this map would result in 7 extra bytes padding per key/value pair. By packing key/value pairs as key/key/value/value, the padding only has to be appended to the end of the byte array and not in between. Eliminating the padding bytes saves the bucket and the map a good amount of memory.

A bucket is configured to store 8 key/value pairs. If you need to add the ninth key to a full bucket, an overflow bucket is created and reference from inside the respective bucket.

![Map Hash Table](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/images/overflow-bucket.png)

### How Map Grow

As we continue to add and remove key/value pairs from the map, the efficiency of the map lookups begins to deteriorate. The load threshold values that determine when to grow the hash table are based on these four factors:

- **overflow**: percentage of buckets which have an overflow bucket.
- **bytes/entry**: overhead bytes used per key/value pair.
- **hitprobe**: number of entries to check when looking up a present key.
- **missprobe**: number of entries to check when looking up an absent key.

The hash table starts growing by adding a pointer to the current bucket array called "old bucket" pointer.
Then a new bucket array has been allocated to hold twice the number of existing buckets. This can result in large allocations, but the memory is not initialized so the allocation is fast.

Once the memory for the new bucket is available, the key/value pairs from the old bucket array can be moved or "evacuated" to the new bucket array. The key/value pairs that are together in an old bucket could be moved to different buckets inside the new bucket array. The evacuation algorithm attempts to distribute the key/value pairs evenly across the new bucket array.

This is a very delicate dance because iterators still need to run through the old buckets until every old bucket has been evacuated. This also affects the performance of iteration operations.

### Map in other languages

![map-in-other-languages](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/images/map-in-other-languages.png)

### Speed

![Speed](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/images/speed.png)

### Size

![size](https://github.com/gkjoyes/golang-tour/blob/master/lesson/03/map/images/size.png)

## Reference

- [Inside the Map Implementation by Keith Randall](https://youtu.be/Tl7mi9QmLns)
- [Macro view of Map by William Kennedy](https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html)
