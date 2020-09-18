# multiunmarshal

`multiunmarshal` provides the `UnmarshalMultiple` method to deseralize a list of JSON documents into
a list of Go structs.

## Example Usage

```
type person struct {
    Name string `json:"name"`
}

peopleJson := [][]byte{
    []byte(`{"name":"person1"}`),
    []byte(`{"name":"person2"}`)
}

var people []person
if err := UnmarshalMultiple(peopleJson, &people); err != nil {
    // handle err
}
```