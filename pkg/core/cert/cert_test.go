package cert

import (
	"github.com/chcp/bsn-sdk-go/pkg/util/keystore"
	"github.com/chcp/bsn-sdk-go/third_party/github.com/hyperledger/fabric/bccsp"
	"encoding/hex"
	"fmt"
	"github.com/cloudflare/cfssl/csr"
	"testing"
)

func TestGetCSRPEM(t *testing.T) {

	name := "abc"
	cr := GetCertificateRequest(name)
	var ks bccsp.KeyStore

	fks, err := keystore.NewFileBasedKeyStore(nil, "./test/msp/keystore", false)

	ks = fks

	key, cspSigner, err := keystore.BCCSPKeyRequestGenerate(ks)

	if err != nil {
		fmt.Println(err)

	}
	csrPEM, err := csr.Generate(cspSigner, cr)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("key:", hex.EncodeToString(key.SKI()))

	fmt.Println("csrPEM：", string(csrPEM))

}
