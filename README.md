# QRCODE-Auth
In order for this to work we need to tweak server.go in "http" package. You can find the source code in GOROOT\src\net\http\server.go, replace GOROOT by actual path.
Open server.go as administrator then add this peace of code :

```go
func (mux *ServeMux) UnhandleFunc(pattern string) {
	mux.mu.Lock()
	defer mux.mu.Unlock()
	delete(mux.m, pattern)

}
```
this will add the Unhandlefunc to ServeMux struct thant delete the handle pattern within ServeMux map containing the roots.
I hope this helps someone achieve whaterver they need it for.



https://user-images.githubusercontent.com/5898693/185229262-79579d42-8974-44a7-8f75-7f9c66e53251.mp4

