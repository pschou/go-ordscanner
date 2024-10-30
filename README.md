# Ordered Scanner

A bufio.Scanner reader which reads from multiple scanners and returns an ordered list

```golang
func ExampleNew() {
  s := ords.New(
    bufio.NewScanner(strings.NewReader("a\nc\nd\n")),
    bufio.NewScanner(strings.NewReader("b\ne\nf\ng\n")),
    bufio.NewScanner(strings.NewReader("f\nz\n")),
  )

  for s.Scan() {
    fmt.Println(s.Text())
  }
  // Output:
  // a
  // b
  // c
  // d
  // e
  // f
  // f
  // g
  // z
}
```
