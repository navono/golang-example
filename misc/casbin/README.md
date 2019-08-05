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

## 详细解释

PERM(Policy, Effect, Request, Matchers) 模型访问控制：
- Policy：定义权限的规则
- Effect：定义组合了多个 Policy 之后的结果，allow/deny
- Request：访问请求，也就是操作
- Match而：判断 Request 是否满足 Policy


使用文件定义权限模型，支持以下模型：
- ACL
- ACL with superuser
- ACL without users
- ACL without resources
- RBAC
- RBAC with resource roles
- RBAC with domains/tenants
- ABAC

model file 语法：
casbin 是基于 PERM 的, 所有 model file 中主要就是定义 PERM 4 个部分.

- Request definition

    `[request_definition]`
    
    r = sub, obj, act
    分别表示 request 中的
    
    - accessing entity (Subject)
    - accessed resource (Object)
    - the access method (Action)

- Policy definition

    `[policy_definition]`
    
    p = sub, obj, act
    p2 = sub, act
    定义的每一行称为 policy rule, p, p2 是 policy rule 的名字. p2 定义的是 sub 所有的资源都能执行 act

- Policy effect

    `[policy_effect]`
    
    e = some(where (p.eft == allow))
    上面表示有任意一条 policy rule 满足, 则最终结果为 allow

- Matchers

    `[matchers]`
    
    m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
    定义了 request 和 policy 匹配的方式, p.eft 是 allow 还是 deny, 就是基于此来决定的

- Role

    `[role_definition]`
    
    g = _, _
    g2 = _, _
    g3 = _, _, _
    g, g2, g3 表示不同的 RBAC 体系, _, _ 表示用户和角色 _, _, _ 表示用户, 角色, 域(也就是租户)


https://www.cnblogs.com/wang_yb/p/9987397.html