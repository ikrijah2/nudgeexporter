// Code generated by "esc -pkg mappings -o plugin/storage/es/mappings/gen_assets.go -ignore assets -prefix plugin/storage/es/mappings plugin/storage/es/mappings"; DO NOT EDIT.

package mappings

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/.nocover": {
		name:    ".nocover",
		local:   "plugin/storage/es/mappings/.nocover",
		size:    43,
		modtime: 1551330067,
		compressed: `
H4sIAAAAAAAC/youSSzJzFYoSEzOTkxPVcjILy4pVkgsLcnXTU/NSy1KLElNUUjLzEkt1uMCBAAA//8y
IKK1KwAAAA==
`,
	},

	"/jaeger-dependencies-7.json": {
		name:    "jaeger-dependencies-7.json",
		local:   "plugin/storage/es/mappings/jaeger-dependencies-7.json",
		size:    279,
		modtime: 1612587018,
		compressed: `
H4sIAAAAAAAC/2yPQWrDMBBF9z6FmGVIRDfd+BAttAcwivTtTJGnqmYMBaO7F5l00dDt478Hfx+cI5aE
76kEM1RRGh2dPgIW1EtCgSRIZOjlROe+VpixLEpjl391L9t6RZ0+50lvoaae2XfnXw78Or8f0LV2/l+q
KJljeNTe7vhRXEMpLIsXqCFNMyMn9ZlXNhqfn/5sK742qKmPId7gIeGaQaPVDYNzR5fuvX6qDW34CQAA
//8AfUo/FwEAAA==
`,
	},

	"/jaeger-dependencies.json": {
		name:    "jaeger-dependencies.json",
		local:   "plugin/storage/es/mappings/jaeger-dependencies.json",
		size:    273,
		modtime: 1612587018,
		compressed: `
H4sIAAAAAAAC/2yPQWrDMBBF9z7FMMuQiG660SFaaA9gFOvbmSJPVc0YCsZ3LzbZNGT7+O/BXzsidsy1
JAdH4tNXwoR2yajQDB0Edjnxed8Z3EUn47hrRCya8Rt0ma9o/ffY2y21bBxpXSm8Hfh9/Dwgbdv5udRQ
iwzpUfu440dxTrWKTkFhjtyPgpItFJnFOb6+/Ns2/CwwtzCk4YYATdcCjt4WdERHl++9/dTWbd1fAAAA
//8zWep3EQEAAA==
`,
	},

	"/jaeger-service-7.json": {
		name:    "jaeger-service-7.json",
		local:   "plugin/storage/es/mappings/jaeger-service-7.json",
		size:    1148,
		modtime: 1612587018,
		compressed: `
H4sIAAAAAAAC/8ySQY/TMBCF7/kVozmuWgshLYfcOSDBglhxQsiaOpPU4NjGdrsbRf7vKMbbZgvdEwdO
kWbmvXzzPHMDgNp2/Cg9pcTBRmwBb+YZxNv7T4F7/Qg5fyceOGwjh6NWvL3BTQMwz1vQPYgvkd+9/wA5
L15kNEVeTBZrAHzZKTB1CC3Mizg/ubLtql3klLQdIrbVrrAKexh3HKTrZdxT6MrvZhB3pfyxvy9FyMXw
L6LA3mhFVXbWfa71S+VI3ms7CMsxcSd7zaaLwuhRJ2xvXz2bDfzzwDFFoUjtWbClnWFsUzhw2e6PyAA2
aHTPalLLYM2tOFoalwrW0LQZt94ZrSbcnIeCM8YdOcgS/TL+cuQPQSfGos+XeZe1sa57Dr2bLI1aycSj
N5SW5/1aAVa00ZOViYYoR/JP2tqrls+rAJgmz9jiD54eXOhWa/1OdLAusKSdOzK2r2/frNp5PYue0l6O
lNQeW0w0iBs8tXNzoVgx++AUx/hfYFcWcQ2/fL/VY/PBeQ5JczzDYX3iu3I2qy2v4l5FPWHi8htK2tl/
4Xq6utzk5lcAAAD//16uKiV8BAAA
`,
	},

	"/jaeger-service.json": {
		name:    "jaeger-service.json",
		local:   "plugin/storage/es/mappings/jaeger-service.json",
		size:    1056,
		modtime: 1612587018,
		compressed: `
H4sIAAAAAAAC/8xTy47bIBTd+yvQXUYJqiqlCz4ildplVaEbc23TAqaA00aR/73CQ8aPibKaxXjhx+E8
OBhuFWOQyHqDiUAw2P1CaikcIoWLrumwg32mREpJuzaCyArGQDtF/7gb7JmC7BsZOwwqgmC3G+OnCf7a
fJ9ANo77x6JA3ugat7JvBd4KLXqvXcsdxURKNpqMitxoqxOI46cVN9CfgWKKvMa6I04Oz4ZApDDQG08K
XF0dWl2DaNBEqhibkqEkzrWlogYHk+QdyRgaM38yBi9hajbLV6nCGJQseV/2COLHq3i2yavu0cmEbZQW
/TJiGi2T2+L5h149gYDfdP3bBwX77bhuXR9I4rm/EIjPxy8rwrjmg8fUSYup7kBAwpbvYEEYqwe6VQsf
+ppi/GBFyqz4s0Ll7We1cINyMhYbwIfeU0ia4mobFOIJLa2rPav1pNKiDuRATLp37+deLZ/5PlZj9T8A
AP//IuFOPCAEAAA=
`,
	},

	"/jaeger-span-7.json": {
		name:    "jaeger-span-7.json",
		local:   "plugin/storage/es/mappings/jaeger-span-7.json",
		size:    3680,
		modtime: 1612587018,
		compressed: `
H4sIAAAAAAAC/+xXwW7bMAy95ysEHovEGAZ0B5+3Q4G1G9buNAwGY9OOWlnSJCZtYPjfBytO6sR2OmDu
MAw7BZb4nh4pPsauZkKA1Bk9JRaZyWkPsYCLqhLRh9vPjnL5JOr6Hqkgt/AW9eIC5jMhqmohZC6ir56u
Pl6Lum6IUEn01DA0vELAGRpHmDWBDbDeM5LOWipPzFIXHuKWKoiM9LpckktMnvgVuiwcVYnoJix/ym/D
oqgD4QDIkVUyxVPYl3b5FFiitVIXkSbPlCW5JJX5SMlSMsSXb45iHf1Yk2cfpZiuKCKNS0UQs1tTCKuq
Xr2EmIOSOaXbtAltiyYEaCybZ2jrJVW5sEbJdAvzfYgzSpkNuSTUvAk+U+tHJ5kgQOu9mEOpQ8rQpvpc
72yrsZRpwlRahdzc6rf28L3O5pos6oSx8EmJdo9t91rK41UhgLeWIIYH2j4alx1SandloY2jBJdmQxC/
vXzX2a67sWCRV0mJnK4gBsYiuoDDdj07QXQ0W2dS8v6vkN1qicbkh9/vbaNZZyw5luSfxQE7TOnqfVfu
uNRRmQeJYNGR5luLekJSPy1dUwZkafRNMMpUIhkd38lhRmV0AeOQa6mU9EPADJm6OnLjSmSIgaxJV0m5
A/aYs/Uuw1/XkissBhVIzc0o6COUGQbsxl1XdDsMIM5ReZofeanXkjsqWZJnLO2Yj7pJ9E2yG7Vj2J7A
syLPCg2bD7Ttr75k+Rdt38sqIDao1vTHTmMs7gLt65w3G3uqxyfxfuYddd5oH3lyG5nSqdEnnciMo9Pe
LO8pZTgH/d+l/2SXOsrJkU7p1Uako3wo6+nauv9mMCl//0/9d+nPXMeJR4f9eRT+avc24MPJajrouykH
3dQdN/LKGj4sZvXsZwAAAP//iNEyH2AOAAA=
`,
	},

	"/jaeger-span.json": {
		name:    "jaeger-span.json",
		local:   "plugin/storage/es/mappings/jaeger-span.json",
		size:    3826,
		modtime: 1612587018,
		compressed: `
H4sIAAAAAAAC/+xWTW/bPAy++1cIPBap8eIFuoPPu+ywDlh7GwaDsWlHrb4mMd2Cwv998EeX2JWdDYiH
AVsOSUzpIfmI5GM9J0IAk3YKmSATcPWAVJO/Dg7N9RVs2vVAzNLUAbJ2uxAgTUnfUrPXW/K5rfKwQ18G
yMTzs0hvO/OH6q4ziqbZxEGenJIFTmEfB/MUqNE5aerUUGAq80qSKkOqpJYM2c1/o72evuwpcEgLLHaU
ksGtIsjY7+mVT/JpeTCoZQFZhSpQIkQXGYaIR9p5SRXuFecvltaGSh0fhYA+WHl01n4GKkLAECt/OfMA
2acf4KOb9tQdmpyxDrlGdxqiWx2Sm9rbah4cQQaPdPhqfQmb6bqsjfWU49Y+EWT/37wZbWjG+8Eh73KN
XOwgA8Y6vYKTDU0SwY1YOG8LCuEPIzJklS4RGv59Tk68dUU5qb7z1pFnSWHUA+yxoHdvx5yW+CxwOeEB
Dj0ZvnNoVnAe1nHbHhCytOYWNV0+aUbP93LOs7KmhmXge6mUDHF42YriKKvKeo0MGZCzxS7XPTgaodz3
vH81s0phPZOPNNyKcxyn7Bysl8wxkbHmbSYjG2nq3qHUFBi1mx/XMbHYJPbSPe8hku6ZlM+k3S0/0iFm
P68zP6E1EZYd6gnVnn57VMb6vnO+Ztxk6bk594p4EeBJvy50XiD/JAt6LSIrvCoYF15HdvtABcM5B//6
+6/ub08VeTIFrS7Jnqr4aVx6KGJ3mhXixK4hlwhztmSvpn5u4ieg1SscnewLn/rMJF9eWNfp1cUbfP/b
fjdJk3wPAAD//5ROSoHyDgAA
`,
	},

	"/": {
		name:  "/",
		local: `plugin/storage/es/mappings`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"plugin/storage/es/mappings": {
		_escData["/.nocover"],
		_escData["/jaeger-dependencies-7.json"],
		_escData["/jaeger-dependencies.json"],
		_escData["/jaeger-service-7.json"],
		_escData["/jaeger-service.json"],
		_escData["/jaeger-span-7.json"],
		_escData["/jaeger-span.json"],
	},
}