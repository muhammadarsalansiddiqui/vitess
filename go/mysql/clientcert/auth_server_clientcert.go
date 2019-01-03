package clientcert

import (
	"fmt"
	"net"
	"vitess.io/vitess/go/mysql"
	querypb "vitess.io/vitess/go/vt/proto/query"
)

type AuthServerClientCert struct {}

type UserData struct {
	conn *mysql.Conn
}

// Init is public so it can be called from plugin_auth_ldap.go (go/cmd/vtgate)
func Init() {
	ascc := &AuthServerClientCert{}
	mysql.RegisterAuthServerImpl("clientcert", ascc)
}

// AuthMethod is part of the AuthServer interface.
func (ascc *AuthServerClientCert) AuthMethod(user string) (string, error) {
	return "mysql_clear_password", nil
}

// Salt is not used for this plugin.
func (ascc *AuthServerClientCert) Salt() ([]byte, error) {
	return mysql.NewSalt()
}

// ValidateHash is unimplemented.
func (ascc *AuthServerClientCert) ValidateHash(salt []byte, user string, authResponse []byte, remoteAddr net.Addr) (mysql.Getter, error) {
	panic("unimplemented")
}

// Negotiate is part of the AuthServer interface.
func (ascc *AuthServerClientCert) Negotiate(c *mysql.Conn, user string, remoteAddr net.Addr) (mysql.Getter, error) {
	certs := c.GetTLSClientCerts()
	if certs == nil {
		return nil, fmt.Errorf("no client certs for connection ID %v", c.ConnectionID)
	}

	if _, err := mysql.AuthServerReadPacketString(c); err != nil {
		return nil, err
	}

	return &UserData{
		conn: c,
	}, nil
}

func (ud *UserData) Get() *querypb.VTGateCallerID {
	certs := ud.conn.GetTLSClientCerts()
	return &querypb.VTGateCallerID{
		Username: certs[0].Subject.CommonName,
		Groups: certs[0].DNSNames,
	}
}
