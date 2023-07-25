# cacooclip

Fork of Changkun Ou's [golang-design/clipboard](https://github.com/golang-design/clipboard) that handles clipboard data
from Chrome with Cacoo Editor's custom web MIME type, on MacOS only. 

Yes, it is very specialized to work on my machine. It might or might not support more OS'es and browsers in the future.

### Features

The Cacoo Editor supports copy and paste of shapes to/from the system clipboard by using a custom MIME type. 
When you CTRL+C a shape, you can only CTRL+V it in the Cacoo Editor. The cacooclip program allows you to 
read the raw shape data from the clipboard and output it as JSON text.

### Usage

This fork is designed as a command line program. It may evolve into a library at some point in the future.

At this stage, the program can be run with the go toolchain. 
Copy a shape from the Cacoo Editor, navigate your local copy of this repo and then run: 

```
go run _example/main.go
```
The program will output the raw shape JSON data.

### License

MIT | &copy; 2021 The golang.design Initiative Authors, written by [Changkun Ou](https://changkun.de) (original repo and the code that remains in this fork)
and 2023 Gabriele V. (Nulab Inc.) (all modifications and additions committed to this fork)
