[fogleman/primitive](https://github.com/fogleman/primitive) compiled from Go to JavaScript using [GopherJS](https://github.com/gopherjs/gopherjs).

See it in action here: https://kdzwinel.github.io/primitive/

## More details

I've created a simplified version of `main.go` file (`main-js.go`) that exposes API to JavaScript and consumes images encoded with Base64. For sample use, see `index.html`.

To compile the project to JavaScript use `gopherjs build main-js.go`. Note that included `main-js.hand.js` was manually optimized (yes, original performance was even worse), so expect output from GopherJS to be 4x times slower. There is still plenty opportunity to make it faster (rewrite parts of generated code, minify JS, use web worker, web assembly, etc.). This project is just a PoC.
