# blocky

semi-fast half block art encoder for go

## examples

```go
// encode image to terminal
blocky.NewEncoder(os.Stdout).Encode(img)

// encode and put in variable
var art string = blocky.EncodeToString(img)
```