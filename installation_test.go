package bunq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCreateInstallation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Response":[{"Id":{"id":4971}},{"Token":{"id":11352,"created":"2017-03-11 12:40:03.613887","updated":"2017-03-11 12:40:03.613887","token":"2237d5026c1d70121ac6fef73bbe937c1a8bc7340ef0acd253d6d9f85a6dbfda"}},{"ServerPublicKey":{"server_public_key":"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu8U1n5MCfdo6kiv6m4yq\nDQ\/YDU3DI3lruoqKE8tXsTY9F1Hv8ySrcvin3YB3tLhP0w30rbaphvYQXg09OYgj\nqQI4Lix0F42+8G3kEmnbceNETF2rY90QJrLY2jp2+xKngsmprvGzTCG+eNCl5d+h\nuSr4mfYWCE05RRteDhikXulQGc1GdVTlAbFAP4OaBEYenEc2TUu92P8tPH0H9EZV\njEoj\/qyOR\/ZJJkHrUdiIbEXfFao0r00CMbfJjXa\/+gyACwKSaH7RDXfrREvn8doT\nGOWegVVIeqB7xfkp5BBnS\/Y7AKwuX+3FAsjOoS58cOhFNm2PQJ5DbgRdb\/vI6YHW\nUwIDAQAB\n-----END PUBLIC KEY-----\n"}}]}`)
	}))
	defer ts.Close()

	client := NewClient()
	client.BaseURL = ts.URL

	privKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA74uJ/q21yZk67NzktTL3khNpOnfsKVa7Xev/0W8rvPmhxPzT
N1OmrAFDzbVLtxdJ5T9yHtzEb0qfsESVPCayaa9tqfOm3Bg1qa5SWrSgCEW2BvPD
t12HDBVeelh4L3NFnj6ulLhfvi86+a2DHfdUx9OLt4huOpQxFstxLAa3zvJhH7OH
GotwzXyvJhMuJBUUSx7ZWenY2yz632D6N7+0t2DQyhtZ1s5e4SCtNV/1si+7KzXs
PVgi6slIYdUqaiAZJF0YGgQovGkr0s8povNm5KgR4Ovqv2M3In+hO04YqnLTHrrn
5VSL49kQGWphzTrp9T5wwn5/ZwPnnB7YcDAOEwIDAQABAoIBAHl3eHH8A8JGQOr6
175KKd+YmDNdvBL6N+hYU1AP303kB3OsAC597HYr7gXReKNO29mzYlrj93e3j2IC
ZOordSzCGAml02anoA56pqf4D24iazr7QLMqaeBmtZG0ar0k5phnkH85PtNhf7Y7
ldEMKaFqU96s/7gUjQ/R+YEppur4Yb/eXFfvYFUbFuc3CHQyUradKBJxKh4H7Lys
hRPpHVOQkCGYwiTrJz/9/bL86tnI/nbug59nXhTe3aai0qaHTeMJ/G2xbUkXefoh
c5xKCTBkw1KgZ24ZUYm6gJdFPRNYA5SieZWt3mIfLQ6zQc/k614kxRcew0fGW4Ri
5+sVSoECgYEA+hsdDBvAYaSDgHhs2XV/vX9kuUJiDB+3OGmXWQ3YN8mzv/tBTQPX
8goSwEOVKeiZfCKvpP2bJSwv4sBNujSDXJCrHZBJwf55Wj/kcgdp12D/rifHsFQv
rJ1WJjpSlchcbQ1PdQtQP7hO5QcjK6U7m5To1kiDcZYfkWqcATY9tU8CgYEA9TC2
TzD9pzjJBT86u2PO2IOfwkxV6HiEWA5OPPRNaPybKk4a+fvnaufG00zLrOwRj4Tp
asMqlYACY/x8xFIB5QxNCYQTLvMlaSovgr+rSkm+bGrYmHk/ITy4AbUiy8ioZvTm
t5CfcaO7WmrlV+dcJ0ftzE2UTX0BJyaNtzIUcf0CgYEA7+DbbkabsMsCGVDnTXaF
qzGpYIpL0ccFiwSzVYWS0IcTcNnCGuTJ1GpW67KmOUjPFSGLh1p52CBWWUwKAMLn
DvvuMu+13mt85tOK/tcfa6Sr9dRPkU5dX1iUTRv5I5HFHA79G4xbTpIukTnUQMM8
tY8P9p4b+/B5nJY8xGjKrL8CgYEAxDLYj4HqV1dPNA2ml7CEIikhO78Nt1pIvJWl
8YykLPCF0VJyr7rtMVSKeyamjJbSbn+ysCW/+6VVRGEUDZx5u6keNBElsJoMQ5zo
K73n+SgNYoAVFd1fsN7/dw5U67CDYO9zd0wY6jxUfUOwhaiyyxP5q1Qg6eivdX6a
RA+k4JkCgYBjaTlXVWFC45HKuttYNpIGPUwowiqzcmx8rx17ymrI8qxxhSAA+nf1
N3xaYHI6thoqfyq4JxMBxvBYDQBKhnCfLxk9AO2O/Uq7OOiRhWqtV4SdMD7Hb7O9
GZ4h0C1AqVJNAxfUH5p7vKxp4E73SYVry88zdHkFj6nYfgXkasBBVA==
-----END RSA PRIVATE KEY-----`
	r := strings.NewReader(privKey)
	if err := client.SetPrivateKey(r); err != nil {
		t.Fatal(err)
	}

	got, err := client.CreateInstallation()
	if err != nil {
		t.Fatal(err)
	}

	exp := &Installation{
		ID: 4971,
		Token: InstallationToken{
			ID:        11352,
			Token:     "2237d5026c1d70121ac6fef73bbe937c1a8bc7340ef0acd253d6d9f85a6dbfda",
			CreatedAt: time.Unix(1489236003, 613887000).UTC(),
			UpdatedAt: time.Unix(1489236003, 613887000).UTC(),
		},
		ServerPublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu8U1n5MCfdo6kiv6m4yq
DQ/YDU3DI3lruoqKE8tXsTY9F1Hv8ySrcvin3YB3tLhP0w30rbaphvYQXg09OYgj
qQI4Lix0F42+8G3kEmnbceNETF2rY90QJrLY2jp2+xKngsmprvGzTCG+eNCl5d+h
uSr4mfYWCE05RRteDhikXulQGc1GdVTlAbFAP4OaBEYenEc2TUu92P8tPH0H9EZV
jEoj/qyOR/ZJJkHrUdiIbEXfFao0r00CMbfJjXa/+gyACwKSaH7RDXfrREvn8doT
GOWegVVIeqB7xfkp5BBnS/Y7AKwuX+3FAsjOoS58cOhFNm2PQJ5DbgRdb/vI6YHW
UwIDAQAB
-----END PUBLIC KEY-----
`,
	}

	if eq := reflect.DeepEqual(exp, got); !eq {
		t.Errorf("Expected: `%#v`, got: `%#v`", exp, got)
	}
}

func TestGetInstallation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Response":[{"Id":{"id":12}}]}`)
	}))
	defer ts.Close()

	client := NewClient()
	client.BaseURL = ts.URL

	privKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA74uJ/q21yZk67NzktTL3khNpOnfsKVa7Xev/0W8rvPmhxPzT
N1OmrAFDzbVLtxdJ5T9yHtzEb0qfsESVPCayaa9tqfOm3Bg1qa5SWrSgCEW2BvPD
t12HDBVeelh4L3NFnj6ulLhfvi86+a2DHfdUx9OLt4huOpQxFstxLAa3zvJhH7OH
GotwzXyvJhMuJBUUSx7ZWenY2yz632D6N7+0t2DQyhtZ1s5e4SCtNV/1si+7KzXs
PVgi6slIYdUqaiAZJF0YGgQovGkr0s8povNm5KgR4Ovqv2M3In+hO04YqnLTHrrn
5VSL49kQGWphzTrp9T5wwn5/ZwPnnB7YcDAOEwIDAQABAoIBAHl3eHH8A8JGQOr6
175KKd+YmDNdvBL6N+hYU1AP303kB3OsAC597HYr7gXReKNO29mzYlrj93e3j2IC
ZOordSzCGAml02anoA56pqf4D24iazr7QLMqaeBmtZG0ar0k5phnkH85PtNhf7Y7
ldEMKaFqU96s/7gUjQ/R+YEppur4Yb/eXFfvYFUbFuc3CHQyUradKBJxKh4H7Lys
hRPpHVOQkCGYwiTrJz/9/bL86tnI/nbug59nXhTe3aai0qaHTeMJ/G2xbUkXefoh
c5xKCTBkw1KgZ24ZUYm6gJdFPRNYA5SieZWt3mIfLQ6zQc/k614kxRcew0fGW4Ri
5+sVSoECgYEA+hsdDBvAYaSDgHhs2XV/vX9kuUJiDB+3OGmXWQ3YN8mzv/tBTQPX
8goSwEOVKeiZfCKvpP2bJSwv4sBNujSDXJCrHZBJwf55Wj/kcgdp12D/rifHsFQv
rJ1WJjpSlchcbQ1PdQtQP7hO5QcjK6U7m5To1kiDcZYfkWqcATY9tU8CgYEA9TC2
TzD9pzjJBT86u2PO2IOfwkxV6HiEWA5OPPRNaPybKk4a+fvnaufG00zLrOwRj4Tp
asMqlYACY/x8xFIB5QxNCYQTLvMlaSovgr+rSkm+bGrYmHk/ITy4AbUiy8ioZvTm
t5CfcaO7WmrlV+dcJ0ftzE2UTX0BJyaNtzIUcf0CgYEA7+DbbkabsMsCGVDnTXaF
qzGpYIpL0ccFiwSzVYWS0IcTcNnCGuTJ1GpW67KmOUjPFSGLh1p52CBWWUwKAMLn
DvvuMu+13mt85tOK/tcfa6Sr9dRPkU5dX1iUTRv5I5HFHA79G4xbTpIukTnUQMM8
tY8P9p4b+/B5nJY8xGjKrL8CgYEAxDLYj4HqV1dPNA2ml7CEIikhO78Nt1pIvJWl
8YykLPCF0VJyr7rtMVSKeyamjJbSbn+ysCW/+6VVRGEUDZx5u6keNBElsJoMQ5zo
K73n+SgNYoAVFd1fsN7/dw5U67CDYO9zd0wY6jxUfUOwhaiyyxP5q1Qg6eivdX6a
RA+k4JkCgYBjaTlXVWFC45HKuttYNpIGPUwowiqzcmx8rx17ymrI8qxxhSAA+nf1
N3xaYHI6thoqfyq4JxMBxvBYDQBKhnCfLxk9AO2O/Uq7OOiRhWqtV4SdMD7Hb7O9
GZ4h0C1AqVJNAxfUH5p7vKxp4E73SYVry88zdHkFj6nYfgXkasBBVA==
-----END RSA PRIVATE KEY-----`
	r := strings.NewReader(privKey)
	if err := client.SetPrivateKey(r); err != nil {
		t.Fatal(err)
	}

	got, err := client.GetInstallation(42)
	if err != nil {
		t.Fatal(err)
	}

	exp := &Installation{
		ID: 12,
	}

	if eq := reflect.DeepEqual(exp, got); !eq {
		t.Errorf("Expected: `%#v`, got: `%#v`", exp, got)
	}
}

func TestGetInstallationID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Response":[{"Id":{"id":72}}]}`)
	}))
	defer ts.Close()

	client := NewClient()
	client.BaseURL = ts.URL

	privKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA74uJ/q21yZk67NzktTL3khNpOnfsKVa7Xev/0W8rvPmhxPzT
N1OmrAFDzbVLtxdJ5T9yHtzEb0qfsESVPCayaa9tqfOm3Bg1qa5SWrSgCEW2BvPD
t12HDBVeelh4L3NFnj6ulLhfvi86+a2DHfdUx9OLt4huOpQxFstxLAa3zvJhH7OH
GotwzXyvJhMuJBUUSx7ZWenY2yz632D6N7+0t2DQyhtZ1s5e4SCtNV/1si+7KzXs
PVgi6slIYdUqaiAZJF0YGgQovGkr0s8povNm5KgR4Ovqv2M3In+hO04YqnLTHrrn
5VSL49kQGWphzTrp9T5wwn5/ZwPnnB7YcDAOEwIDAQABAoIBAHl3eHH8A8JGQOr6
175KKd+YmDNdvBL6N+hYU1AP303kB3OsAC597HYr7gXReKNO29mzYlrj93e3j2IC
ZOordSzCGAml02anoA56pqf4D24iazr7QLMqaeBmtZG0ar0k5phnkH85PtNhf7Y7
ldEMKaFqU96s/7gUjQ/R+YEppur4Yb/eXFfvYFUbFuc3CHQyUradKBJxKh4H7Lys
hRPpHVOQkCGYwiTrJz/9/bL86tnI/nbug59nXhTe3aai0qaHTeMJ/G2xbUkXefoh
c5xKCTBkw1KgZ24ZUYm6gJdFPRNYA5SieZWt3mIfLQ6zQc/k614kxRcew0fGW4Ri
5+sVSoECgYEA+hsdDBvAYaSDgHhs2XV/vX9kuUJiDB+3OGmXWQ3YN8mzv/tBTQPX
8goSwEOVKeiZfCKvpP2bJSwv4sBNujSDXJCrHZBJwf55Wj/kcgdp12D/rifHsFQv
rJ1WJjpSlchcbQ1PdQtQP7hO5QcjK6U7m5To1kiDcZYfkWqcATY9tU8CgYEA9TC2
TzD9pzjJBT86u2PO2IOfwkxV6HiEWA5OPPRNaPybKk4a+fvnaufG00zLrOwRj4Tp
asMqlYACY/x8xFIB5QxNCYQTLvMlaSovgr+rSkm+bGrYmHk/ITy4AbUiy8ioZvTm
t5CfcaO7WmrlV+dcJ0ftzE2UTX0BJyaNtzIUcf0CgYEA7+DbbkabsMsCGVDnTXaF
qzGpYIpL0ccFiwSzVYWS0IcTcNnCGuTJ1GpW67KmOUjPFSGLh1p52CBWWUwKAMLn
DvvuMu+13mt85tOK/tcfa6Sr9dRPkU5dX1iUTRv5I5HFHA79G4xbTpIukTnUQMM8
tY8P9p4b+/B5nJY8xGjKrL8CgYEAxDLYj4HqV1dPNA2ml7CEIikhO78Nt1pIvJWl
8YykLPCF0VJyr7rtMVSKeyamjJbSbn+ysCW/+6VVRGEUDZx5u6keNBElsJoMQ5zo
K73n+SgNYoAVFd1fsN7/dw5U67CDYO9zd0wY6jxUfUOwhaiyyxP5q1Qg6eivdX6a
RA+k4JkCgYBjaTlXVWFC45HKuttYNpIGPUwowiqzcmx8rx17ymrI8qxxhSAA+nf1
N3xaYHI6thoqfyq4JxMBxvBYDQBKhnCfLxk9AO2O/Uq7OOiRhWqtV4SdMD7Hb7O9
GZ4h0C1AqVJNAxfUH5p7vKxp4E73SYVry88zdHkFj6nYfgXkasBBVA==
-----END RSA PRIVATE KEY-----`
	r := strings.NewReader(privKey)
	if err := client.SetPrivateKey(r); err != nil {
		t.Fatal(err)
	}

	exp := 72
	got, err := client.GetInstallationID()
	if err != nil {
		t.Fatal(err)
	}

	if eq := reflect.DeepEqual(exp, got); !eq {
		t.Errorf("Expected: `%#v`, got: `%#v`", exp, got)
	}
}
