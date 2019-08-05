# Casbin

访问控制模型基于 `PERM` (Policy, Effect, Request, Matcher) 的一个文件。

最基本、最简单的 model 是 ACL：

```
# Request definition
[request_definition]
r = sub, obj, act

# Policy definition
[policy_definition]
p = sub, obj, act

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))

# Matchers
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act

```

ALC model 的 policy：
```
p,abc123,user,GET
p,abc123,user,POST
p,admin,test,*
g,super_admin,admin
```

表示：
> 用户 abc123 对 user 有 GET 权限，下面同理
>
> 角色或用户组 admin 对 test 有所有权限
>
> 用户 super_admin 属于 admin 组或角色
