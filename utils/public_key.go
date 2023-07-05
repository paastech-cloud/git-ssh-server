package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/gliderlabs/ssh"
)

// ParsePublicKey parses the public key to a string.
func ParsePublicKey(key ssh.PublicKey) string {
	keyString := base64.StdEncoding.EncodeToString(key.Marshal())

	return fmt.Sprintf("%s %s", key.Type(), keyString)
}
