## Go compiler

编译成 `.o` 文件：
> go tool compile unsafe.go

检查 `.o` 文件
> ar t unsafe.o

编译成 `.a` 文件
> go tool compile --pack unsafe.go

检查 `.a` 文件
> ar t unsafe.a

使用 `goobj` 检查 `.a` 文件：
> go get -u -v github.com/s-matyukevich/goobj_explorer

编译汇编输出：
> go tool compile -S unsafe.go

### 资料
- [Golang Internals - Siarhei Matsiukevich](https://blog.altoros.com/golang-internals-part-3-the-linker-and-object-files.html)

## GC

查看 `GC` 的实际情况。
> go run gColl.go

更详细地查看 `GC` 的运行状况：
> GODEBUG=gctrace=1 go run gColl.go

```
gc 7 @0.112s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 47->47->0 MB, 48 MB goal, 4 P
```

`47->47->0 MB` 的意思是：第一个 `47` 是 `GC` 运行之前的堆大小；第二个 `47` 是 `GC` 运行结束后的堆大小；第三个 `0` 是当前堆（live heap）的大小

`Golang` 的 `GC` 是基于 `tricolor algorithm` 算法。严格说是 `tricolor mark-and-sweep algorithm` 算法。使用 `write barrier` 和
程序并行运行。


### 资料
- [Golang’s Real-time GC in Theory and Practice](https://making.pusher.com/golangs-real-time-gc-in-theory-and-practice/)
