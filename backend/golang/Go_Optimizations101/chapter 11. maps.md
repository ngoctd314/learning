# Maps

In Go, the capacity of a map is unlimited in theory, it is only limited by available memory. That is why the builtin cap function doesn't apply to maps.

In the official standard Go runtime implementation, maps are implemented as hashtables internally. Each map/hashtable maintains a backing array to store map entries (key-value pairs). Along with more and more entires are put into a map, the size of the backing array might be thought as to small to store more entires, thus a new larger backing array will be allocated and the current entires (in the old backing array) will be moved to it, then the old backing array will be discarded.  

In the official standard Go runtime implementation, the backing array of a map will never shrink, even if all entries are deleted from the map. This is a form of memory wasting.
