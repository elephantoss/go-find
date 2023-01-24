# Find

Find is a utility library that provides a convenient way to find elements in collections.


## Why do I need it?

If you are tired of writing the same 3 lines of code every time you want to get a
value out of a collection or use a default value and you wonder why there is not
a more convenient way to do it, then this lib is for you.


## Example

Get a value out of a map:

```go
value := maps.Find(myMap, "key", "some-default-value")
```


