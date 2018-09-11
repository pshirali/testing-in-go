# Samples


The subfolders here match the example/pattern tag from the slides.
Unless otherwise stated, you can `chdir` to the respective folders
below and run `go test -v`


## Test composition


#### samples/ex1 --- Hello World through tests


#### samples/ex2 --- Table driven tests & subtests

You can also try running individual tests:
```
go test -v -run=TestAdderUsingTable
go test -v -run=TestAdderUsingSubtests
go test -v -run=TestAdderUsingSubtests/Positive
```


#### samples/ex3/test_normal --- Tests for `UnsafeCounter`


#### samples/ex3/test_suite --- Suite

In addition to `-v`, you can also add `-d` and `-w` switches.
`-d` shows the Type and address of the instances under test per-test.
`-w` shows `before` and `after` wrapper functions at Suite level and per-test


## Fixtures


#### samples/F1 --- No code here. The pattern is used in `ex3`


#### samples/F2 --- Idempotent teardowns

Run `go test -v -d` for debug output.

Additionally, the code comments mention changes that can be made to
simulate errors and obesrve behavior in the output


#### samples/F3 --- Fixture that returns a teardown func

The code has commented `t.Fatalf` lines which can be uncommented to
induce errors in both SETUP and TEARDOWN


#### samples/F4 and samples/F5

These are variants with minor changes.


#### samples/G1 --- Gotcha 1

Lines that demonstrate behavior are printed directly. Just run `go test -v`


#### samples/G1 --- Gotcha 2

Lines that demonstrate behavior are printed directly. Just run `go test -v`
