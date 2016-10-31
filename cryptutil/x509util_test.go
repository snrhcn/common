package cryptutil

import (
	"strings"
	"testing"
)

func TestCertificateDecoding(t *testing.T) {

	_, err := ReadX509CertsFromFile(invalidFileName)
	if err == nil {
		t.Error("Attempting to load an invalid file should result in an error")
		return
	}

	googleCert := `
-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----
`

	c, err := ReadX509Certs([]byte(googleCert))
	if err != nil {
		t.Error(err)
		return
	}

	if len(c) != 1 {
		t.Error("Only one certificate should have been read")
		return
	}

	if res := Sha256CertFingerprint(c[0]); res != "d0:88:88:3c:7b:b3:da:b4:9e:d8:bf:ec:43:aa:92:cb:29:58:e8:e2:e1:c3:89:8d:73:50:6a:b8:c8:f1:12:21" {
		t.Error("Unexpected fingerprint:", res)
		return
	}

	if res := Sha1CertFingerprint(c[0]); res != "ee:b6:d4:d8:88:e5:75:5f:ff:c0:19:27:b6:67:9c:77:e8:0d:2c:7f" {
		t.Error("Unexpected fingerprint:", res)
		return
	}

	if res := Md5CertFingerprint(c[0]); res != "5c:a6:bd:96:9c:96:79:a7:90:ee:89:a6:ee:1a:04:a8" {
		t.Error("Unexpected fingerprint:", res)
		return
	}

	// Test error cases

	_, err = ReadX509Certs([]byte(googleCert[2:]))
	if err.Error() != "PEM not parsed" {
		t.Error("PEM parsing error expected:", err)
		return
	}

	_, err = ReadX509Certs([]byte(googleCert[0:29] + "Mi" + googleCert[31:]))
	if strings.HasPrefix("asn1: structure error", err.Error()) {
		t.Error("asn1 parsing error expected:", err)
		return
	}
}
