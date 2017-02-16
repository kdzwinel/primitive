[fogleman/primitive](https://github.com/fogleman/primitive) compiled from Go to JavaScript using [GopherJS](https://github.com/gopherjs/gopherjs)

## More details

I've created a simplified version of `main.go` file (`main-js.go`) that exposes API to JavaScript and consumes images encoded with Base64.

To recompile it using GopherJS use `gopherjs build main-js.go`.

Note that included `main-js.hand.js` was manually optimized (yes, original performance was even worse), so expect output from GopherJS to be 4x times slower.