// Code generated by "esc -o assets_gen.go -pkg logs -ignore .go assets/"; DO NOT EDIT.

package logs

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
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
	return nil, nil
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

	"/assets/control_panel.html": {
		local:   "assets/control_panel.html",
		size:    6096,
		modtime: 1518879672,
		compressed: `
H4sIAAAAAAAC/+xYX4vcOBJ/n09RpztIArFNjnsKbh+zM9kQSGYDyS7sU5Ctsi1Glowkd6cx/u6LJLun
e8bd7ckfyC7ph7ZsVf1KKtVf9T3DkksEUrSfhKoqLisyDBcXfW+xaQW1YapGygjEw3CRMr6GQlBjVkSr
DckuAAD2vxZKRKKKXvx3nPPz9YtpuqUVRg4PNcneBpFpUr/YI26zd1TSCqFBY9xTqMoAlQzWVHPVGVC2
Rg225rIyMTz9WHMD3MgnFjY1aoScysBwy4tbA1R7jArZc9iqDgoqwaAFW1MLXQudZB4OoVEMNbVcSbA0
fxanSTvuMGF8PQ7/FUWQxLt9QhRlF+P8Pb1RgdqaUXOBTatNYPgyRe7Nt1SiAP8fMSxpJ+we5Sy1V7w7
40M693uNEjUVTi/WqfUQ6G73x7FzxbYzwGmpdANaCVwRNyTQoK0VW5FWGTvDcB/9TjvnKEeN/e8E+QOW
GovbXH0+wwKQCpqjOEcFkHLZdhbstsU9dJC0wRX53aB2o9HuX0maC2QE+p6XEF8pWfIqnicaBvBoyPoe
JRuG80sJnAnjxj2hG3FhdPQ018l5kD9Hf/HCgXqUJ2Z0px2mgQ23tXehTa24gUI1DZUsXiZk/0RyExVU
CNVZuBtGXJaKZGme3SiLL9Mkz+BN6cUH91YSqAxhwaBeo75bUK6smx7fuQEMGn3uZqX74/pgI0JAPgUM
wDVK4KWPG5MmuQfcqm4SFrY54yUzW00WGNIZpHPTfyeXuOHF7VmXmCf6epeQI+63dIkJ86tc4se3EhfK
o0qrriULLCDs5hdBi1vBjYWiplKiMFBq1Rzk+WnnCwOHQYGFnRbVdMLy8IkAZyuSTxKjSeJkd7u1ILva
TXn+1mWqabRgc+7X983Hu8Qf8D6p1hUS5pPHIkDu5MSXheVrfN1xweLpM5APfuWOh+xM/z3VBtneaqeN
DMN57SRBF9/LjE5N/TMTeChM3wV7NVdU/sFxc40C7YOYdYr0CyLXpRBqAw02uUt5PrSkhWI41cqToDTx
X6FF3XBjnDGBVbDmuAEWpE/+Zn6mqkcd/qs16q2SePLY54m++MBxhPtBz/B7RYCDpmeWJ++sdV2aPynT
5Q23ZFdCWgm5lZHpigKN8WNR+UcuVHFLIPtA1+hUDB/GfidNAmIGAN9QE2niEmV2qpu6/xr6xNDduU7x
G/d+UAikuuRzTvO+ywUvDvtup2RXNY9l9cWp02s7ISLNq9qSzHuFVBbiX7k29j2tcBhSCrXGckX+T0uL
etX38Q1u0Nhh2J2epOtIcHkLu2Ns/KPVvKF6SzLHodOEZqMvpXQp6yQ8x1JpdNJ/EyxIz9zIo841u6f7
X+uKyUijaZU0fD1XNKSe5oDhWOtr9QlXsXX25jpNbH2a5kojtcjOE152tlZ6AWCoOhYAhgLmOGGaHNtg
3/+ngperw+rozfWRQqfvNZUVQvxWVceKoTO6ZNm/+95LSBPLTlP2vXNkaj/yBiEe9Xtpl7HGQc3DAE93
L07ss/PMboWj8m9og0vknQ6c/jpmwRXMTFocrbfmjKEkY+y1+NlOGZIzAmsqOlyRUbEEkkdm3SmWe3V7
Y1qRshMiJL6HMT54OHO2oHfSr0fiJcLpEch7QSNpfXBMnJUOQ+JiYzJtMnOp3sWOM2n2YS645xns8T7j
Q+BM3vExJnv8NV6plEV9Kks4h4N3W2io/JkNluXvhdfHh3fE+zIfXMmHYwpXyxeTFfwVAAD//4LETr/Q
FwAA
`,
	},

	"/assets/log_view.html": {
		local:   "assets/log_view.html",
		size:    5244,
		modtime: 1518879672,
		compressed: `
H4sIAAAAAAAC/6xYeW/cuBX/fz7FC+PCUmJJ4xS9PMfWR7trINsumqToYrtIOOKThghFaklqxl5B370g
dYxGsR23aP7IUHwHH3/vpOuaYcYlAimrjeDpR4N6h/qjULkhTTOb1bXFohTUIpC0/LhFygjETTNbGnsv
EOx9iSti8c4mqTFkPZsBAPy5QMZpUHAZ7Tmz2wv4w+//WN6FUHuy+2fpRuDoe9iLBL1Xlb3I+B2yxRE9
edVqO5/Pf7N4lRzRMiVtZPiveAHz+E9YHEvulWbRRiP9fAH+J6JCHFiaYTUsXhYmj1IloB6dCaMbnc/n
5d2igeTVO7SWyxzsFsETQWWQKlEVEn4bj+x8aXmBR0p/53U8reL8SAWt7FbpsZI382coeXOsJLVcSTO5
39e1DLdpAfP/xQwFWmRRgcbQHA8uTZVQ+gJ078Zmtkx81KxnS8Z3kApqzIpotSdrzzDeTZWIRB6dv+lo
U3pJc4xcOKIecXiuTOkC2iuuSFJQSXNM6jq+TC3f4bcVFyy+vWmaRKg85zJPskqI9hYECrRbxVakVMZO
FHvlXJaVHYU9AUkLXJHbG9KbtuWMoSSwo6LCFanr+K3KjT/yIY3b8/X1lkqJAlzawcTSv9ECmwaWpqBC
rJcUthqzFUnIenMPP15++8PNVXx3/+syoetl0jLVNc8gvjWXrODSyY6NNtWm4HYwdmMlbKyMRO5/GJU5
6sH2mw6VZF3XKFnTLJPt+QTuxOE9clLC+K7z52j5IoogiQevQhStZx19WmSoQG1NV2b+D4HikPX/Rwwz
WompX7/g9nHFZT7he9l7snNX55mg+7y9uYAJh3N5OEFrgGTYaqE5OtnD86iRbZ3UaEolDd/hQ1HV1tex
QFdet2qHulsbq3mJjABnKyJUHvntB9S1Kp1xD9Naun6c2CnwB/VVkKzf8wIh+PD+OoQl743NKGQ0SqlG
G1UlWS8Tvl4mdvs85Yf6SNaXfv182a7kk/X3bSnzkpNkep4Rh/pK1pftR6fMZ9HjGCaPgejEPfyP+Gaj
2P3jt6zrk+4GcLF6zm3q+uSayn9y3LclgHm5460nxbUrJHDCz+Bk52V9XnTAmqcgsHoIR632UV2f8KYZ
6lWpShfAkdU8z1FDx0aAUUsjq/Jc4MDV77a8K5IKnn4mYLl1TJ0xwGWmOs5S0BQLlHZFrCoJtLupktbv
tQEFXaa3Xy7HoVfVUbrPRyr+8W3Zuq5jlwjG0qJ0FdayZ8m0x38wqKUvRC+HvRtuUs2LZ+o6hD4KEdV1
26igjfvB1T3+k27v+HxIt3miNARS2UEuhJNpzNR1fN3i6dYoDDbN+y030GmELTWwQZSgsVA7ZJBpVfjW
GIOS4h6oi10DKZVgEIHbeGhOHhmewclz87XjH9vcNMtNZa2SAzDt5wGaSes0xVHrlCpVMuO6ACV9wPWg
dVERnHaKTkOybs9cJu0R6+OL/K+1os3Ax6WXySPlYpn4BrB+qmVNP0fdy3etEU//07JotW+7vgvO0q5n
O+rT19zewArGU9JillXSV004Rq4wedjOlzwLXnQwB+RSI9yrCkzVLfZUWrCqkwY7Cq9v/i1/dKxbVQnW
M3ALXIJLG6UZcGtQZE6+oJ+x1crtqYFcSWyjkYp2WItJGB4GXo220rIfdmft1i+wAol7+Nf3b7+ztvwH
/lKhsUG46OkxZewvO5T2LTcWJeqACEUZOYMehQB39gwopaOzeAaBEzaW2srAixW8mc/DyYPKT1IBeafc
VOvGij1KC3utZH4B5LWT76YIDEevIZeUx4pOAvJyKBHktfODM/vaJULwRU0IFw8Kd2nUibfZHYTTV1jz
JDKotdJjaEaQdLe9lOC5QKVppZH15owVqxJlQH74+7v35AzI1trSXCTujfCdMrZpnvNoKEzevRnCRetr
F8+uX3QO/6vSxQ21tL+iI8UGbUDeqvz2hpx1wf8Feege5AwcVAezDUoWOMZwMWtmPoH8nPShhBWcPj1C
ncKB/0bt5eMSTO3lkYwb2K6oQScyGtlOPdEobZF9KPdUMwMrsLrCxewkmHrIRcEw+YWxr4zBQ25MlTRK
YCxUHrQd++DDoyf69OQXRxuL2VGuHNGmeTKxbWsLEQyXft1jPAlrp/G9K5eBEz9Mz+FP85/PPAzh11Lq
a2c6P/1Xp2ZUmPGxXeDPmnB2qKkHFV70DDTuUBvsUfEu3zhXOnJsrxTjaLz+JIHKIHzq5s1PrkjyXCrt
9/yA+gmoZO4rU8quP4FWe3P4q4+GFVxqTe/jUiur3JM0NoKnGKdUiMBuXJswZzAP/WHu2eo2gEurgDrB
g1c7D3udVsfuWkNAQUDPYBNC7bQ4yrEd7opcMry7hBXQmLNYo5/9gmHudIVhHHaDyBWsYPNskZJqg8wd
41e30gbtwQ/yXX3Bd/VQ5LeNpvcbfDOcEg16LoZVv3d5XAYzpQMOK5gvgMPSAShQ5na7gNeveQh2E9Oy
RMmut1ywwOqf+M/hwqHZbgPSdOtAdZ1TaYbaVaRl0jf3yZveBQPq7k+K7WTynwAAAP//BGKfp3wUAAA=
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/assets": {
		isDir: true,
		local: "assets",
	},
}
