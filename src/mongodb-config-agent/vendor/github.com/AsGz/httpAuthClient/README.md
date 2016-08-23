# httpAuthClient

## Add digest authenticates header for http request 
reference resources:
- [http-digest-auth-client](https://github.com/ryanjdew/http-digest-auth-client)
- [go-http-auth-server](https://github.com/abbot/go-http-auth)

## example

```
go get github.com/AsGz/httpAuthClient

```


```go

url := "https:xxxxxxxxx"
username := "yourname"
password := "pass"

params := "params"
req, err := http.NewRequest("POST", url, strings.NewReader(params))
req.Header.Set("Content-Type", "application/json")

err = httpAuthClient.ApplyHttpDigestAuth(username, password, url, req)
if err != nil {
	log.Fatal(err)
} else {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, string(b), err)
}

```
