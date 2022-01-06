# choose 

`githbu.com/lestrrat-go/choose` is a simple tool that allows users to choose random
elements out of a container. It uses the new type parameters introduced in go 1.18

# Slices

Initialize a choose.SliceChooser:

```go
src := []int{....}
c := choose.Slice[int](src)
```

## Choose a random element from a slice

```go
elem := c.One()
```

## Choose N random elements from a slice

```go
elems := c.N(3)
```

# Map

Initialize a choose.MapChooser:

```go
src := map[string]int{...}
c := choose.Map[string,int](src)
```

## Choose a random element (as choose.MapElement) from a map

```go
elem := c.One()
fmt.Printf("%s -> %v", elem.Key, elem.Value)
```

## Choose N random elements (as []choose.MapElement) from a map

```go
elem := c.N(3)
```
