﻿# QRCODE-Auth
In order for this to work we need to tweak server.go in "http" package. You can find the source code in GOROOT\src\net\http\server.go, replace GOROOT by actual path.
Open server.go as administrator as file is protected then add this peace of code :

```go
func (mux *ServeMux) UnhandleFunc(pattern string) {
	mux.mu.Lock()
	defer mux.mu.Unlock()
	delete(mux.m, pattern)

}
```
this will add the Unhandlefunc to ServeMux struct thant delete the handle pattern within ServeMux map containing the roots.
I hope this helps someone achieve whaterver they need it for.
