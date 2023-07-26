# cacooclip

Fork of Changkun Ou's [golang-design/clipboard](https://github.com/golang-design/clipboard) that handles clipboard data
from Chrome with Cacoo Editor's custom web MIME type, on MacOS only. 

Yes, it is very specialized to work on my machine. It might or might not support more OS'es and browsers in the future.

### Features

The Cacoo Editor supports copy and paste of shapes to/from the system clipboard by using a custom MIME type. 
When you CTRL+C a shape, you can only CTRL+V it in the Cacoo Editor. The cacooclip program allows you to 
read the raw shape data from the clipboard and output it as JSON text.

### Usage

This fork is designed as a command line program or as a library.

#### Command line:

```
go install github.com/vibridi/cacooclip/cmd/cacooclip@latest
```
or from the cloned repo:
```
$ make build
```
then move the binary into your path. 
Copy a shape from the Cacoo Editor then run:

```
$ cacooclip -r
```
The command will print the Cacoo shape data to stdout. To write data, select and copy the shape data as JSON text
and (on MacOS) run:

```
$ pbpaste | cacooclip -w
```

#### Library:

```
$ go get github.com/vibridi/cacooclip
```
and then call it as:

```go
str, err := cacooclip.Read()
if err != nil {
	// ...
}
err = cacooclip.Write(str)
if err != nil {
	// ...
}
```

### License

MIT | &copy; 2021 The golang.design Initiative Authors, written by [Changkun Ou](https://changkun.de) (original repo and the code that remains in this fork)
and 2023 Gabriele V. (Nulab Inc.) (all modifications and additions committed to this fork)
