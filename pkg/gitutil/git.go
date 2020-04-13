package gitutil

import (
	"errors"
	"github.com/pubgo/g/pkg/shutil"
	cryptossh "golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
	"log"
	"strings"
)

func GitSSHAuth(privateKeyPath string) transport.AuthMethod {
	sshKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("readfile: %v", err)
	}

	signer, err := cryptossh.ParsePrivateKey(sshKey)
	if err != nil {
		log.Fatalf("parseprivatekey: %v", err)
	}

	return &ssh.PublicKeys{User: "git", Signer: signer}
}

func GitCredentialAuth() (transport.AuthMethod, error) {
	_auth := &http.BasicAuth{}

	_dt, err := shutil.Execute("git credential fill")
	if err != nil {
		return _auth, err
	}

	for _, k := range strings.Split(strings.TrimSpace(_dt), "\n") {
		if _k := strings.ReplaceAll(k, "username=", ""); _k != k {
			_auth.Username = _k
		}

		if _k := strings.ReplaceAll(k, "password=", ""); _k != k {
			_auth.Password = _k
		}
	}

	return _auth, nil
}

func GitBasicAuth(username, password string) (transport.AuthMethod, error) {
	_auth := &http.BasicAuth{}

	if username == "" || password == "" {
		return _auth, errors.New("username or password is null")
	}

	_auth.Username = username
	_auth.Password = password
	return _auth, nil
}
