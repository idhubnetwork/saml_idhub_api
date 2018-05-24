package keypairs

import (
	"errors"
	"os/exec"
)

func GenPrivateKey(dir string) (string, error) {
	out, err := exec.Command("openssl", "genrsa", "-out", dir, "4096").CombinedOutput()
	if err != nil {
		return "", errors.New(err.Error() + " : " + string(out))
	}
	return "Key generated!", nil
}

func GenKeyAndCert(keyDir, certDir string) (string, error) {
	out, err := exec.Command("openssl", "req", "-newkey", "rsa:4096", "-nodes", "-keyout",
		keyDir, "-x509", "-days", "3650", "-out", certDir, "-subj",
		"/C=CN/ST=GD/L=BJ/O=idhub/OU=product/CN=idhub.network/emailAddress=support@idhub.network").CombinedOutput()
	if err != nil {
		return "", errors.New(err.Error() + " : " + string(out))
	}
	return "Cert and Key generated!", nil
}
