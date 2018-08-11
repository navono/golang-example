# 并发的安全操作

- 为共享内存进行同步（sync.Mutex）
- 通过通信来同步（channels）
- 使用`Immutable`数据类型
- 使用`confinement`保护数据

## 数据的限制（`confinment`）

有两种方式：

- ad hoc
- lexical

# for-select loop

- Sending iteration variables out on a channel

```go
for _, s := range []string{"a", "b", "c"} {
  select {
    case <-done:
      return
    case stringStream <- s:
  }
}
```

- Looping infinitely waiting to be stopped

```go
for {
  select {
    case <-done:
      return
    default:
  }

  // Do non-preemptable work
}
```
