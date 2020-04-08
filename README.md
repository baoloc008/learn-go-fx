# Learn Go Fx

Basic DI is straightforward in Go:

- Higher-layer package calls interface functions, not knowing (nor caring) which package implements them
- Lower-layer package implements interface functions
- startup code injects lower-layer implementation via higher-layer constructors

## Links
- https://blog.huyage.dev/posts/simple-dependency-injection-in-go/
- https://ewanvalentine.io/my-thoughts-on-uber-fx/

## Note
- Fx là dependency injection framework, vì là framework nên cần viết code theo cấu trúc của nó
- Fx có lifecycle, có thể gọi các hàm lúc app start, stop
