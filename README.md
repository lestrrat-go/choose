# choose 

`githbu.com/lestrrat-go/choose` is a simple tool that allows users to choose random
elements out of a container. It uses the new type parameters introduced in go 1.18

# Slices

```go
c := choose.Slice(src)
// Choose one random element
elem := c.One()

// Choose N random elements
elems := c.N(3)
```
