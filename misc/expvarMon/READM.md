# expVarMon

install
> go get -u -v github.com/divan/expvarmon

prepare app
```go
import (
  _ "expvar"
  "net/http"
)
...
http.ListenAndServe(":1234", nil)
```

run
> expvarmon -ports="1234"
