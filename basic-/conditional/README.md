# 控制语句

## if

语法：

```go
if condition {

}
```

或者加上 else ：

```go
if condition {  
} else if condition {
} else {
}
```

`condition` 右侧的花括号是强制的。

还存在 `if` 语句的变体：

```go
if statement; condition {

}
```

在 `statement` 中声明的变量只能在 `if` 语句块中使用。

## 循环

语法：

```go
for initialisation; condition; post {

}
```

在 `for` 语句块中，还有 `break` 和 `continue` 两个辅助控制的关键字。

`initialisation` 可以放在 `for` 语句之外进行。

当 `initialisation; condition; post` 均省略后，就是无限循环。

## switch

默认每个 `case` 都带有 `break`，匹配成功后不会自动向下执行其他 `case`，而是跳出整个 `switch`, 但是可以使用 `fallthrough` 强制执行后面的 `case` 代码。
