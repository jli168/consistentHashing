### consistentHashing

A demo of consistent hashing using Go

After reading that consistent hashing is widely used in cache machines in distributed system, and because it is quite simple,
I decided to write a water-down demo of consitent hashing using Go.

Suppose each node is a string, we add some nodes into the consistent hashing circle, or remove some anytime.

Then suppose we have some object to store, the object's key is also a string, we find the key's nearest counter-clockwise node.

References: 
*  https://en.wikipedia.org/wiki/Consistent_hashing
  
*  https://github.com/kkdai/consistent/blob/master/consistent.go
