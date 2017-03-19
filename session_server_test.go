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

func TestSession(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Response":[{"Id":{"id":72}},{"Token":{"id":839,"token":"b165cce82bbd229b55962f90b4efedd706b3f616f0de831547ff62262f2924e3"}},{"UserCompany":{"id":42,"created":"2015-06-13 23:19:16.215235","updated":"2015-06-30 09:12:31.981573","public_uuid":"252e-fb1e-04b74214-b9e9467c3-c6d2fbf","name":"bunq","display_name":"bunq","public_nick_name":"bunq","alias":[{"type":"EMAIL","value":"bravo@bunq.com","name":""}],"chamber_of_commerce_number":"NL040492904","type_of_business_entity":"One man business","sector_of_industry":"Education","counter_bank_iban":"NL12BUNQ1234567890","avatar":{"uuid":"5a442bed-3d43-4a85-b532-dbb251052f4a","anchor_uuid":"f0de919f-8c36-46ee-acb7-ea9c35c1b231","image":[{"attachment_public_uuid":"d93e07e3-d420-45e5-8684-fc0c09a63686","content_type":"image/jpeg","height":380,"width":520}]},"address_main":{"street":"Example Boulevard","house_number":"123a","po_box":"09392","postal_code":"1234AA","city":"Amsterdam","country":"NL"},"address_postal":{"street":"Example Boulevard","house_number":"123a","po_box":"09392","postal_code":"1234AA","city":"Amsterdam","country":"NL"},"version_terms_of_service":"1.2","director_alias":{"uuid":"252e-fb1e-04b74214-b9e9467c3-c6d2fbf","avatar":{"uuid":"5a442bed-3d43-4a85-b532-dbb251052f4a","anchor_uuid":"f0de919f-8c36-46ee-acb7-ea9c35c1b231","image":[{"attachment_public_uuid":"d93e07e3-d420-45e5-8684-fc0c09a63686","content_type":"image/jpeg","height":380,"width":520}]},"public_nick_name":"Mary","display_name":"Mary","country":"NL"},"language":"en_US","region":"en_US","ubo":[{"name":"A. Person","date_of_birth":"1990-03-27","nationality":"NL"}],"status":"ACTIVE","sub_status":"APPROVAL","session_timeout":1,"daily_limit_without_confirmation_login":{"value":"12.50","currency":"EUR"},"notification_filters":[{"notification_delivery_method":"URL","notification_target":"https://my.company.com/callback-url","category":"PAYMENT"}]}}]}`)
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

	got, err := client.CreateSession()
	if err != nil {
		t.Fatal(err)
	}

	exp := &Session{
		ID: 72,
		Token: SessionToken{
			ID:    839,
			Token: "b165cce82bbd229b55962f90b4efedd706b3f616f0de831547ff62262f2924e3",
		},
		UserCompany: UserCompany{
			ID:             42,
			CreatedAt:      Time(time.Unix(1434237556, 215235000).UTC()),
			UpdatedAt:      Time(time.Unix(1435655551, 981573000).UTC()),
			PublicUUID:     "252e-fb1e-04b74214-b9e9467c3-c6d2fbf",
			Name:           "bunq",
			DisplayName:    "bunq",
			PublicNickName: "bunq",
			Alias: []Alias{
				Alias{
					Type:  "EMAIL",
					Value: "bravo@bunq.com",
					Name:  "",
				},
			},
			ChamberOfCommerceNumber: "NL040492904",
			TypeOfBusinessEntity:    "One man business",
			SectorOfIndustry:        "Education",
			CounterBankIBAN:         "NL12BUNQ1234567890",
			Avatar: Avatar{
				UUID:       "5a442bed-3d43-4a85-b532-dbb251052f4a",
				AnchorUUID: "f0de919f-8c36-46ee-acb7-ea9c35c1b231",
				Image: []AvatarImage{
					AvatarImage{
						AttachmentPublicUUID: "d93e07e3-d420-45e5-8684-fc0c09a63686",
						ContentType:          "image/jpeg",
						Height:               380,
						Width:                520,
					},
				},
			},
			AddressMain: Address{
				Street:      "Example Boulevard",
				HouseNumber: "123a",
				POBox:       "09392",
				PostalCode:  "1234AA",
				City:        "Amsterdam",
				Country:     "NL",
			},
			AddressPostal: Address{
				Street:      "Example Boulevard",
				HouseNumber: "123a",
				POBox:       "09392",
				PostalCode:  "1234AA",
				City:        "Amsterdam",
				Country:     "NL",
			},
			VersionTermsOfService: "1.2",
			DirectorAlias: DirectorAlias{
				UUID: "252e-fb1e-04b74214-b9e9467c3-c6d2fbf",
				Avatar: Avatar{
					UUID:       "5a442bed-3d43-4a85-b532-dbb251052f4a",
					AnchorUUID: "f0de919f-8c36-46ee-acb7-ea9c35c1b231",
					Image: []AvatarImage{
						AvatarImage{
							AttachmentPublicUUID: "d93e07e3-d420-45e5-8684-fc0c09a63686",
							ContentType:          "image/jpeg",
							Height:               380,
							Width:                520,
						},
					},
				},
				PublicNickName: "Mary",
				DisplayName:    "Mary",
				Country:        "NL",
			},
			Language: "en_US",
			Region:   "en_US",
			UBO: []UBO{
				UBO{
					Name:        "A. Person",
					DateOfBirth: Time(time.Unix(638496000, 0).UTC()),
					Nationality: "NL",
				},
			},
			Status:         "ACTIVE",
			SubStatus:      "APPROVAL",
			SessionTimeout: 1,
			DailyLimitWithoutConfirmationLogin: Limit{
				Value:    "12.50",
				Currency: "EUR",
			},
			NotificationFilters: []NotificationFilter{
				NotificationFilter{
					NotificationDeliveryMethod: "URL",
					NotificationTarget:         "https://my.company.com/callback-url",
					Category:                   "PAYMENT",
				},
			},
		},
	}

	if eq := reflect.DeepEqual(exp, got); !eq {
		t.Errorf("Expected: `%#v`, got: `%#v`", exp, got)
	}
}
