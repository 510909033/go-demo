package utils

import (
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"testing"
)

var (
	privateKey2 = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDWIqUso8MUQQEF
kynqOSUogN2p2VEllpO5RQbaQPzM6Su70eWSchVXrCrAaZ5JvTpLkFMMtv5z3D1H
+rQJgXG0SNWHqS+drB+SnEIUUZ7Ts1qtwbBmHe5Se2/AA9hwxz1rUvxsQsNBeMbq
3F66rJ7nc6hp4VB8DkeXLn0+00uuJ9c+DA2FxLj7NpL76cfBH+I79OXy54Pyr+t3
wvzo9zCm2Ix3DN0z/alrcawzUcQ/aVdHtQL7WGtVk1Wk0+j8RTO0iFwcdLqZKW7v
/Chiix5Pu1LLWgnvaMqsha9YSt0tIU08efHSZkbqHUNVI/uZnN0eTpXdHAcfGaFj
XdGNjNs9AgMBAAECggEAbPYMHn7mETRSlMehmHGjmsyX/ol6ZSoN2URqjDxtZsdv
XY8cxjwO6CHPT0FvLg5/z1uNJvcm20XCWRIU8pFPjnFRVkqc1Bp1pmCkrzopG4g0
KB5a6FtqncX5wzcxt5Fqi2Ya0+vU90fB0Fh6S3rzEMslc1sXXUjW3PLnAET8SY1F
a+eTEKRrcCuWy7xTCoAXvs2zt3N2Yt9qIqggNvajj7M1Lh2bBMMeA2oSIGTg0AZ2
vhpcElHU3I/o0TT4O3uxiqkncXO2/j9iUjU6a3p58+kYOAu0eE2DzRhe8+WsRWK2
1IL0rJKwoqV6IJeA+XDpEnrRzKimT8TbDgr6l8bWcQKBgQD72XAF/2t5iqQr7K41
v2YgPtL/v1YtYA2QLdVTwmujr7hS48kYAwF0dGCR9xgF64iNCyBt0rsDsAbzzbNO
JZ0TuXPQkB9YbMhlraFk1Qo8Bc6q3qG25zy0py1zvaE5JDrcsuQJ9YX2bdniyupI
60z+g1dvNmGyKCJ4Wm6TR+yN9wKBgQDZqhczdtCqK8Qf01yespIo7OiFKPr0UyZq
3yzdBVRaCTiw87FCyx2IZ+maI7q7by6Ef3NfS1RMi5CUEU1s3g689wl5gOLv3mWj
RU5KzzVdnlAsuLQszWJ/Qz9kEsA9TE8GNPHcz+Rrd1V87XTLwsdcI11PMBcDETnV
RqtwPbFjawKBgFQa0xxNphiawiShD7SHTJWsJU6fGSpiD2V8yOcF4GhzgDy+MSPB
rTS0wKM+P1mKxLlFY/kJDMc4e7njkeOhEAnMPu8BklY3Td0W1PUVaCuPTOsnGh44
rr8trTIfu72XacjWPO1OfwRCvybL2N2tJrNALWzIXmvel1RV7uNke40VAoGBALEk
VC12DOvxInQkN8SP35Oe+r0kVNhys1lWt9RIehWsW47nROPvGeHhb6QaX+Ya3ejH
MMXlMH37N3bJfhGTxrrS8csIsag1ftIZnVYUvIGsTeOXjkrH+9lGiViQShOY/tc2
T4A9UNCFFfXN6rfJlSXPQikoIhJE+qQwF2YJT0xxAoGBALEbifKKGYX7kkm+XU43
GlHSza+3XBBhoyRv0GdzIF3Ms+VssLJPM5vR+Quz6FL8UiiKD/qr0NBpMaOj3aOV
Dhs1/jPoEZlLhveblrrmsFuKJ6cmV9vlWCjU3ILL49j4avHUFIdC9MT7oOYHWRMN
NYqriZp362sZ0lA+7LxfoOkC
-----END PRIVATE KEY-----
`
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1iKlLKPDFEEBBZMp6jklKIDdqdlRJZaTuUUG2kD8zOkru9Hl
knIVV6wqwGmeSb06S5BTDLb+c9w9R/q0CYFxtEjVh6kvnawfkpxCFFGe07NarcGw
Zh3uUntvwAPYcMc9a1L8bELDQXjG6txeuqye53OoaeFQfA5Hly59PtNLrifXPgwN
hcS4+zaS++nHwR/iO/Tl8ueD8q/rd8L86PcwptiMdwzdM/2pa3GsM1HEP2lXR7UC
+1hrVZNVpNPo/EUztIhcHHS6mSlu7/woYoseT7tSy1oJ72jKrIWvWErdLSFNPHnx
0mZG6h1DVSP7mZzdHk6V3RwHHxmhY13RjYzbPQIDAQABAoIBAGz2DB5+5hE0UpTH
oZhxo5rMl/6JemUqDdlEaow8bWbHb12PHMY8Dughz09Bby4Of89bjSb3JttFwlkS
FPKRT45xUVZKnNQadaZgpK86KRuINCgeWuhbap3F+cM3MbeRaotmGtPr1PdHwdBY
ekt68xDLJXNbF11I1tzy5wBE/EmNRWvnkxCka3Arlsu8UwqAF77Ns7dzdmLfaiKo
IDb2o4+zNS4dmwTDHgNqEiBk4NAGdr4aXBJR1NyP6NE0+Dt7sYqpJ3Fztv4/YlI1
Omt6efPpGDgLtHhNg80YXvPlrEVittSC9KySsKKleiCXgPlw6RJ60cyopk/E2w4K
+pfG1nECgYEA+9lwBf9reYqkK+yuNb9mID7S/79WLWANkC3VU8Jro6+4UuPJGAMB
dHRgkfcYBeuIjQsgbdK7A7AG882zTiWdE7lz0JAfWGzIZa2hZNUKPAXOqt6htuc8
tKctc72hOSQ63LLkCfWF9m3Z4srqSOtM/oNXbzZhsigieFpuk0fsjfcCgYEA2aoX
M3bQqivEH9NcnrKSKOzohSj69FMmat8s3QVUWgk4sPOxQssdiGfpmiO6u28uhH9z
X0tUTIuQlBFNbN4OvPcJeYDi795lo0VOSs81XZ5QLLi0LM1if0M/ZBLAPUxPBjTx
3M/ka3dVfO10y8LHXCNdTzAXAxE51UarcD2xY2sCgYBUGtMcTaYYmsIkoQ+0h0yV
rCVOnxkqYg9lfMjnBeBoc4A8vjEjwa00tMCjPj9ZisS5RWP5CQzHOHu545HjoRAJ
zD7vAZJWN03dFtT1FWgrj0zrJxoeOK6/La0yH7u9l2nI1jztTn8EQr8my9jdrSaz
QC1syF5r3pdUVe7jZHuNFQKBgQCxJFQtdgzr8SJ0JDfEj9+Tnvq9JFTYcrNZVrfU
SHoVrFuO50Tj7xnh4W+kGl/mGt3oxzDF5TB9+zd2yX4Rk8a60vHLCLGoNX7SGZ1W
FLyBrE3jl45Kx/vZRolYkEoTmP7XNk+APVDQhRX1zeq3yZUlz0IpKCISRPqkMBdm
CU9McQKBgQCxG4nyihmF+5JJvl1ONxpR0s2vt1wQYaMkb9BncyBdzLPlbLCyTzOb
0fkLs+hS/FIoig/6q9DQaTGjo92jlQ4bNf4z6BGZS4b3m5a65rBbiienJlfb5Vgo
1NyCy+PY+Grx1BSHQvTE+6DmB1kTDTWKq4mad+trGdJQPuy8X6DpAg==
-----END RSA PRIVATE KEY-----
`
	publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1iKlLKPDFEEBBZMp6jkl
KIDdqdlRJZaTuUUG2kD8zOkru9HlknIVV6wqwGmeSb06S5BTDLb+c9w9R/q0CYFx
tEjVh6kvnawfkpxCFFGe07NarcGwZh3uUntvwAPYcMc9a1L8bELDQXjG6txeuqye
53OoaeFQfA5Hly59PtNLrifXPgwNhcS4+zaS++nHwR/iO/Tl8ueD8q/rd8L86Pcw
ptiMdwzdM/2pa3GsM1HEP2lXR7UC+1hrVZNVpNPo/EUztIhcHHS6mSlu7/woYose
T7tSy1oJ72jKrIWvWErdLSFNPHnx0mZG6h1DVSP7mZzdHk6V3RwHHxmhY13RjYzb
PQIDAQAB
-----END PUBLIC KEY-----
`
	content = `appId=a301020000000000a8ca6985e1deee35&appSecret=de2863203f87ee56f3067473571ee407&bizContent={"playlistId":"5CF4DA1F06D2AB3AC61AB1A665C7D588"}&signType=RSA_SHA256&timestamp=1661407884935`
)

func TestRsa(t *testing.T) {
	bs, err := RSAEncrypt([]byte(content), publicKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(hex.EncodeToString(bs))

	bs, err = RASDecrypt(bs, []byte(privateKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bs))
}

func TestRsaWithKey(t *testing.T) {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		t.Error(err)
	}

	bs, err := RSAEncryptWithKey([]byte(content), pub)
	t.Log(hex.EncodeToString(bs))

	pri, err := ParsePrivateKey(privateKey)
	if err != nil {
		t.Error(err)
	}

	bs, err = RsaDecryptWithKey(bs, pri)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bs))
}

func TestSha256WithRsa(t *testing.T) {
	// Sha256WithRSA
	pri, err := ParsePrivateKey(privateKey)
	if err != nil {
		t.Error(err)
	}
	bs, err := RsaSignWithKey([]byte(content), pri, crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	t.Log(hex.EncodeToString(bs))

	t.Log( base64.StdEncoding.EncodeToString(bs))
	//t.Log( base64.StdEncoding.EncodeToString(hex.EncodeToString(bs)))


	err = RSAVerify([]byte(content), bs, publicKey, crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
}

func TestRsaSign(t *testing.T) {
	data, err := RsaSign([]byte(content), []byte(privateKey), crypto.SHA256)
	if err != nil {
		t.Error(err)
	}

	//fmt.Printf("%x\n\n\n", data)
	//fmt.Printf("%X\n\n\n", data)
	t.Log(base64.StdEncoding.EncodeToString(data))
	t.Log(base64.RawStdEncoding.EncodeToString(data))

}