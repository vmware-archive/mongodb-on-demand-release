// Add digest authenticates header for http request
// reference resources: https://github.com/ryanjdew/http-digest-auth-client
//
// Example:
//
//  url := "https://xxxxxxxxx"
// 	username := "yourname"
// 	password := "pass"

// 	params := "params"
// 	req, err := http.NewRequest("POST", url, strings.NewReader(params))
// 	req.Header.Set("Content-Type", "application/json")

// 	err = httpAuthClient.ApplyHttpDigestAuth(username, password, url, req)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		resp, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		var b []byte
// 		b, err = ioutil.ReadAll(resp.Body)
// 		fmt.Println(resp.StatusCode, string(b), err)
// 	}

package httpAuthClient

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DigestHeaders tracks the state of authentication
type DigestHeaders struct {
	Realm     string
	Qop       string
	Method    string
	Nonce     string
	Opaque    string
	Algorithm string
	HA1       string
	HA2       string
	Cnonce    string
	Path      string
	Nc        int16
	Username  string
	Password  string
}

// add Auth authenticates against a given URI for the request
func ApplyHttpDigestAuth(username, password, uri string, request *http.Request) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 401 {
		authn := digestAuthParams(resp)
		algorithm := authn["algorithm"]
		d := &DigestHeaders{}
		u, _ := url.Parse(uri)
		d.Path = u.RequestURI()
		d.Realm = authn["realm"]
		d.Qop = authn["qop"]
		d.Nonce = authn["nonce"]
		d.Opaque = authn["opaque"]
		if algorithm == "" {
			d.Algorithm = "MD5"
		} else {
			d.Algorithm = authn["algorithm"]
		}
		d.Nc = 0x0
		d.Username = username
		d.Password = password

		d.applyAuth(request)
		return nil
	}
	return fmt.Errorf("response status code should have been 401, it was %v", resp.StatusCode)
}

func (d *DigestHeaders) digestChecksum() {
	switch d.Algorithm {
	case "MD5":
		// A1
		h := md5.New()
		A1 := fmt.Sprintf("%s:%s:%s", d.Username, d.Realm, d.Password)
		io.WriteString(h, A1)
		d.HA1 = fmt.Sprintf("%x", h.Sum(nil))

		// A2
		h = md5.New()
		A2 := fmt.Sprintf("%s:%s", d.Method, d.Path)
		io.WriteString(h, A2)
		d.HA2 = fmt.Sprintf("%x", h.Sum(nil))
	case "MD5-sess":
		// A1
		h := md5.New()
		A1 := fmt.Sprintf("%s:%s:%s", d.Username, d.Realm, d.Password)
		io.WriteString(h, A1)
		haPre := fmt.Sprintf("%x", h.Sum(nil))
		h = md5.New()
		A1 = fmt.Sprintf("%s:%s:%s", haPre, d.Nonce, d.Cnonce)
		io.WriteString(h, A1)
		d.HA1 = fmt.Sprintf("%x", h.Sum(nil))

		// A2
		h = md5.New()
		A2 := fmt.Sprintf("%s:%s", d.Method, d.Path)
		io.WriteString(h, A2)
		d.HA2 = fmt.Sprintf("%x", h.Sum(nil))
	default:
		//token
	}
}

// ApplyAuth adds proper auth header to the passed request
func (d *DigestHeaders) applyAuth(req *http.Request) {
	d.Nc += 0x1
	d.Cnonce = randomKey()
	d.Method = req.Method
	d.Path = req.URL.RequestURI()
	d.digestChecksum()
	response := doMD5(strings.Join([]string{d.HA1, d.Nonce, fmt.Sprintf("%08x", d.Nc),
		d.Cnonce, d.Qop, d.HA2}, ":"))
	AuthHeader := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=%08x, qop=%s, response="%s", algorithm=%s`,
		d.Username, d.Realm, d.Nonce, d.Path, d.Cnonce, d.Nc, d.Qop, response, d.Algorithm)
	if d.Opaque != "" {
		AuthHeader = fmt.Sprintf(`%s, opaque="%s"`, AuthHeader, d.Opaque)
	}
	// fmt.Printf("%v\n", AuthHeader)
	req.Header.Set("Authorization", AuthHeader)
}

/*
Parse Authorization header from the http.Request. Returns a map of
auth parameters or nil if the header is not a valid parsable Digest
auth header.
*/
func digestAuthParams(r *http.Response) map[string]string {
	s := strings.SplitN(r.Header.Get("Www-Authenticate"), " ", 2)
	if len(s) != 2 || s[0] != "Digest" {
		return nil
	}

	result := map[string]string{}
	for _, kv := range strings.Split(s[1], ",") {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			continue
		}
		result[strings.Trim(parts[0], "\" ")] = strings.Trim(parts[1], "\" ")
	}
	return result
}

func randomKey() string {
	k := make([]byte, 12)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			panic("rand.Read() failed")
		}
		bytes += n
	}
	return base64.StdEncoding.EncodeToString(k)
}

/*
H function for MD5 algorithm (returns a lower-case hex MD5 digest)
*/
func doMD5(data string) string {
	digest := md5.New()
	digest.Write([]byte(data))
	return fmt.Sprintf("%x", digest.Sum(nil))
}
