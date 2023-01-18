package rdiff

import (
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	DefaultHexPrivateKey string
	DefaultPrivateKey    *ecdsa.PrivateKey
)

func init() {
	// set DefaultHexPrivateKey from env var
	DefaultHexPrivateKey, ok := os.LookupEnv("EQLDefaultHexPrivateKey")
	if !ok {
		DefaultHexPrivateKey = "907d55a3b4365b8d1c6710445748308bc9f750368d59ad87fd03c46989096789"
		log.Println("You are using default privet key! Please set env var EQLDefaultHexPrivateKey!")
	}

	DefaultPrivateKey, _ = crypto.HexToECDSA(DefaultHexPrivateKey)
}

func Signature(data []byte) string {
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), DefaultPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	return hexutil.Encode(signature)
}

func SignatureVerify(data []byte, signature string) bool {
	pubKey := DefaultPrivateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error with public ECDSA key")
	}

	pubKeyB := crypto.FromECDSAPub(pubKeyECDSA)
	hash := crypto.Keccak256Hash(data)

	signatureNoRecoverID := signature[:len(signature)-1] // without recovery id
	signatureBytes, _ := hexutil.Decode(signatureNoRecoverID)
	return crypto.VerifySignature(pubKeyB, hash.Bytes(), signatureBytes)
}
