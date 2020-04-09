# Learn Go Fx

Basic DI is straightforward in Go:

- Higher-layer package calls interface functions, not knowing (nor caring) which package implements them
- Lower-layer package implements interface functions
- startup code injects lower-layer implementation via higher-layer constructors

## Links
- https://blog.huyage.dev/posts/simple-dependency-injection-in-go
- https://ewanvalentine.io/my-thoughts-on-uber-fx
- https://pkg.go.dev/go.uber.org/fx?tab=doc

## Note
- Fx là dependency injection framework, vì là framework nên cần viết code theo cấu trúc của nó
- Fx có lifecycle, có thể gọi các hàm lúc app start, stop

## Fx Basic
- All Fx applications start with an `fx.App` that can be constructed from `fx.New()`
- Run application with [Run Method](https://pkg.go.dev/go.uber.org/fx?tab=doc#App.Run): `app.Run()`
- `func Provide`: cung cấp các dependency
    - Constructors được gọi lazily
    - Constructors must return one or more objects, and may return an error
    - Constructors should perform as little external interaction as possible, and should avoid spawning goroutines.
- `func Invoke`: gọi hàm khi app start, các tham số được dùng từ `fx.Provide`
    - Invocations được gọi eagerly
    - Invocations có thể return error, khi return các giá trị khác sẽ bị bỏ qua
