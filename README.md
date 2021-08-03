# assertion
go assertion without testing.

# why I need it
I just want to find a tool that could make me write go assertions like:
```go
assertion.Equals(expected, actual)
assertion.NotNull(actual)
```
instead of:
```go
if expected != actual {
  panic("xxxxx")
}
```
but I cannot find a satisfied one. `stretchr/testify` is good but its code like:
```go
assert.Equal(t, 123, 123, "they should be equal")
```
I don't write assertion code in testing so I have not parameter `t` with is `tesing.T`.
I just think: "Why not I write a new tesing lib? I didn't want any powerful functions, I just wan't write too many `if...panic...`" ... and here it is.
