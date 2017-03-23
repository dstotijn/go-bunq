package bunq

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestGetDevice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Response":[{"DevicePhone":{"id":42,"created":"2015-06-13 23:19:16.215235","updated":"2015-06-30 09:12:31.981573","description":"Kees' iPhone","phone_number":"+31612345678","os":"IOS","status":"ACTIVE"}}]}`)
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

	got, err := client.GetDevice(42)
	if err != nil {
		t.Fatal(err)
	}

	exp := DevicePhone{
		ID:          42,
		CreatedAt:   time.Unix(1434237556, 215235000).UTC(),
		UpdatedAt:   time.Unix(1435655551, 981573000).UTC(),
		Description: "Kees' iPhone",
		PhoneNumber: "+31612345678",
		OS:          "IOS",
		Status:      "ACTIVE",
	}

	if eq := reflect.DeepEqual(exp, got); !eq {
		t.Errorf("Expected: `%#v`, got: `%#v`", exp, got)
	}
}

func TestListDevices(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Response":[{"DevicePhone":{"id":42,"created":"2015-06-13 23:19:16.215235","updated":"2015-06-30 09:12:31.981573","description":"Kees' iPhone","phone_number":"+31612345678","os":"IOS","status":"ACTIVE"}},{"DeviceServer":{"id":42,"created":"2015-06-13 23:19:16.215235","updated":"2015-06-30 09:12:31.981573","description":"Mainframe23 in Amsterdam","ip":"255.255.255.255","status":"ACTIVE"}}]}`)
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

	got, err := client.ListDevices()
	if err != nil {
		t.Fatal(err)
	}

	exp := []interface{}{
		DevicePhone{
			ID:          42,
			CreatedAt:   time.Unix(1434237556, 215235000).UTC(),
			UpdatedAt:   time.Unix(1435655551, 981573000).UTC(),
			Description: "Kees' iPhone",
			PhoneNumber: "+31612345678",
			OS:          "IOS",
			Status:      "ACTIVE",
		},
		DeviceServer{
			ID:          42,
			CreatedAt:   time.Unix(1434237556, 215235000).UTC(),
			UpdatedAt:   time.Unix(1435655551, 981573000).UTC(),
			Description: "Mainframe23 in Amsterdam",
			IP:          net.IPv4bcast,
			Status:      "ACTIVE",
		},
	}

	if eq := reflect.DeepEqual(exp, got); !eq {
		t.Errorf("Expected: `%#v`, got: `%#v`", exp, got)
	}
}
