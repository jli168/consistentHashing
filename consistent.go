/**
 * Try to implement consistent hashing.
 *
 * steps:
 * 	1. create a utility function `hash()` to hash a string into int
 * 	2. create a nodes map[int]string to map hashed int to name string
 * 	3. create a keys slice to store the hashed keys
 * 	4. Now we can add/remove a node string into the nodes and keys
 * 	5. Check which ndoe a key would map into. `sort.Search` function is very helpful!
 *
 * Extra note: The reason sort `map[int]string` is not prefered is because in Go, the runtime randomizes map iteration order
 * If you require a stable iteration order you must maintain a separate data structure that specifies that order.
 * from https://blog.golang.org/go-maps-in-action#TOC_7.
 *
 * So we need a `keys` array to store the key hashes
 */

package main

import (
	"hash/fnv"
	"fmt"
	"sort"
)

type Consistent struct {
	nodes map[int]string // nodes names map: key is hashed int, value is node name
	maxSlot int  // max number of slots in the circle
	keys []int  // hashed keys
}

func newConsistent(slots int) Consistent {
	con := Consistent{}
	con.maxSlot = slots
	con.nodes = make(map[int]string)
	con.keys = []int{}

	return con
}

// add new node into the circle
func (con *Consistent) Add(node string) {
	key := con.hash(node) % con.maxSlot

	con.nodes[key] = node
	con.keys = append(con.keys, key)

	// sort keys so we can use binary search for clockwise closest node
	sort.Ints(con.keys)
}

// remove node from the circle
func (con *Consistent) Remove(node string) {
	key := con.hash(node) % con.maxSlot

	if node, ok := con.nodes[key]; ok == false {
		fmt.Println(node, " is a no existing node!")
		return
	}

	// remove it from nodes map
	delete(con.nodes, key)

	// remove it from keys
	for i, v := range con.keys {
		if v == key {
			con.keys = append(con.keys[:i], con.keys[i+1:]...)
			break
		}
	}
}

// get node for a key string
func (con *Consistent) Get( name string) {
	// find hash of the key
	hash := con.hash(name) % con.maxSlot

	// find closest node key's index in the keys array
	i := sort.Search(len(con.keys), func(i int) bool {
		return con.keys[i] > hash
	})

	// find the closest node key
	slot := con.keys[len(con.keys) - 1]
	if i > 0 {
		slot = con.keys[i - 1]
	}

	fmt.Println("hash: ", hash, " map to ", slot," name : ", con.nodes[slot])
}

// hash function: map a string to a hashed int
func (con *Consistent) hash(s string) int {
    h := fnv.New32a()
    h.Write([]byte(s))
    return int(h.Sum32())
}

// main function, see if it is what we get
func main() {
	c := newConsistent(32)

	c.Add("A")
	c.Add("B")
	c.Add("C")
	c.Add("D")
	c.Add("E")

	fmt.Println(c.nodes)

	c.Get("user233")
	c.Get("user333")
	c.Get("user343")
	c.Get("user353")
	c.Get("user363")
	c.Get("user373")

	fmt.Println("---Now do remove---")

	c.Remove("C")
	c.Remove("D")
	c.Remove("E")
	
	fmt.Println(c.nodes)

	c.Get("user233")
	c.Get("user333")
	c.Get("user343")
	c.Get("user353")
	c.Get("user363")
	c.Get("user373")
}