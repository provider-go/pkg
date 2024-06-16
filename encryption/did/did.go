package did

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/provider-go/pkg/encryption/sm2"
	"github.com/provider-go/pkg/encryption/sm3"
	"strconv"
	"strings"
	"time"
)

/**
 * {
 *     "@context": "https://www.w3.org/ns/did/v1",
 *     "id": "did:cmid:000000026atznpmye111",
 *     "versionId": "1",
 *     "created": "2020-03-10T04:24:12.164Z",
 *     "update": "2020-03-10T04:24:12.164Z",
 *     "verificationMethod": [
 *         {
 *             "id": "did:cmid:000000026atznpmye111#key-0",
 *             "type": "Secp256k1",
 *             "controller": "did:cmid:000000026atznpmye111",
 *             "publicKeyHex": "02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71"
 *         }
 *     ],
 *     "proof": {
 *         "type": "Secp256k1",
 *         "creator": "did:cmid:000000026atznpmye111",
 *         "publicKeyHex": "02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
 *         "signatureValue": "eyJhbGciOiJSUzI...GlcTwLtjPAnKb78"
 *     }
 * }
 */

/**
 * {
 *     "@context": "https://www.w3.org/2018/credentials/v1",
 *     "id": "did:cmid:000000054abcdbc0fd42",
 *     "type": ["UniversityDegreeCredential"],
 *     "issuer": "did:cmid:000000026atznpmye222",
 *     "issuanceDate": "2020-03-10T04:24:12.164Z",
 *     "expirationDate": "2020-01-01T19:23:24Z",
 *     "credentialSubject": {
 *         "id": "did:cmid:000000026atznpmye111",
 *         "degree": {
 *             "type": "BachelorDegree",
 *             "name": "Bachelor of Science and Arts"
 *         }
 *     },
 *     "proof": {
 *         "type": "Secp256k1",
 *         "creator": "did:cmid:000000026atznpmye111",
 *         "publicKeyHex": "02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
 *         "signatureValue": "eyJhbGciOiJSUzI...GlcTwLtjPAnKb78"
 *     }
 * }
 */

/**
 * {
 *     "@context": "https://www.w3.org/2018/credentials/v1",
 *     "type": "VerifiablePresentation",
 *     "verifiableCredential": [
 *         {
 *             "@context": [
 *                 "https://www.w3.org/ns/did/v1",
 *                 "https://w3id.org/security/suites/jws-2020/v1"
 *             ],
 *             "id": "did:cmid:000000054abcdbc0fd42",
 *             "type": ["UniversityDegreeCredential"],
 *             "issuer": "did:cmid:000000026atznpmye222",
 *             "issuanceDate": "2020-03-10T04:24:12.164Z",
 *             "expirationDate": "2020-01-01T19:23:24Z",
 *             "credentialSubject": {
 *                 "id": "did:cmid:000000026atznpmye111",
 *                 "degree": {
 *                     "type": "BachelorDegree",
 *                     "name": "Bachelor of Science and Arts"
 *                 }
 *             },
 *             "proof": {
 *                 # 签名算法
 *                 "type": "JsonWebKey2020",
 *                 # 签名时间
 *                 "created": "2021-06-18T21:19:10Z",
 *                 # 证明目的
 *                 "proofPurpose": "authentication",
 *                 # 验证签名公钥
 *                 "verificationMethod": "https://www.chainmaker.org.cn/keys/000000026atznpmye222#key-0",
 *                 # 签名内容
 *                 "jws": "eyJhbGciOiJSUzI...GlcTwLtjPAnKb78"
 *             }
 *         }
 *     ],
 *     "proof": {
 *         "type": "Secp256k1",
 *         "creator": "did:cmid:000000026atznpmye111",
 *         "publicKeyHex": "02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
 *         "signatureValue": "eyJhbGciOiJSUzI...GlcTwLtjPAnKb78"
 *     }
 * }
 */

type Proof struct {
	ProofType      string `json:"type"`
	Creator        string `json:"creator"`
	PublicKeyHex   string `json:"publicKeyHex"`
	SignatureValue string `json:"signatureValue"`
}

type VerificationMethod struct {
	Id           string `json:"id"`
	VMtype       string `json:"type"`
	Controller   string `json:"controller"`
	PublicKeyHex string `json:"publicKeyHex"`
}

type CredentialSubject struct {
	Id   string `json:"id"`
	Info string `json:"info"`
}

type DIDDocument struct {
	Context            string                `json:"context"`
	Id                 string                `json:"id"`
	Version            string                `json:"version"`
	Created            string                `json:"created"`
	Update             string                `json:"update"`
	VerificationMethod []*VerificationMethod `json:"verificationMethod"`
	Proof              *Proof                `json:"proof"`
}

type VCDocument struct {
	Context           string             `json:"context"`
	Id                string             `json:"id"`
	VCtype            string             `json:"type"`
	Issuer            string             `json:"issuer"`
	IssuanceDate      string             `json:"issuanceDate"`
	ExpirationDate    string             `json:"expirationDate"`
	CredentialSubject *CredentialSubject `json:"credentialSubject"`
	Proof             *Proof             `json:"proof"`
}

type VPDocument struct {
	Context              string `json:"context"`
	VCtype               string `json:"type"`
	VerifiableCredential string `json:"verifiableCredential"`
	Proof                *Proof `json:"proof"`
}

/**
 *
 */
func CreateDIDDocument(priKey, did string) string {
	// Initialization information
	didDDocument := new(DIDDocument)
	didDDocument.Context = "https://www.w3.org/ns/did/v1"
	didDDocument.Id = did
	didDDocument.Version = "1"
	currTZTime := getCurrTZTime()
	didDDocument.Created = currTZTime
	didDDocument.Update = currTZTime
	didDDocument.VerificationMethod = make([]*VerificationMethod, 0)
	didDDocument.VerificationMethod = append(didDDocument.VerificationMethod, createVerificationMethod(priKey, did))
	verificationMethod := ""
	for _, v := range didDDocument.VerificationMethod {
		verificationMethod = verificationMethod + v.Id + v.VMtype + v.Controller + v.PublicKeyHex
	}
	hash := sm3Hash(did + "1" + currTZTime + currTZTime + verificationMethod)
	didDDocument.Proof = createProof(priKey, did, hash)
	byte, err := json.Marshal(didDDocument)
	if err != nil {
		return ""
	} else {
		return string(byte)
	}
}

func CreateVCDocument(priKey, id, vcType, issuer, days, issueTo, info string) string {
	vcDocument := new(VCDocument)
	vcDocument.Context = "https://www.w3.org/2018/credentials/v1"
	vcDocument.Id = id
	vcDocument.VCtype = vcType
	vcDocument.Issuer = issuer
	vcDocument.IssuanceDate = getCurrTZTime()
	daysInt, _ := strconv.Atoi(days)
	vcDocument.ExpirationDate = getExpirationDate(daysInt)
	vcDocument.Issuer = issuer
	vcDocument.CredentialSubject = new(CredentialSubject)
	vcDocument.CredentialSubject.Id = issueTo
	vcDocument.CredentialSubject.Info = info
	hash := sm3Hash(id + vcType + issuer + vcDocument.IssuanceDate + vcDocument.ExpirationDate + issueTo + info)
	vcDocument.Proof = createProof(priKey, issuer, hash)
	byte, err := json.Marshal(vcDocument)
	if err != nil {
		return ""
	} else {
		document := string(byte)
		document = strings.ReplaceAll(document, "\"{", "{")
		document = strings.ReplaceAll(document, "}\"", "}")
		document = strings.ReplaceAll(document, "\\\"", "\"")
		return document
	}
}

func CreateVPDocument(priKey, vpType, grantFrom, info string) string {
	vpDocument := new(VPDocument)
	vpDocument.Context = "https://www.w3.org/2018/credentials/v1"
	vpDocument.VCtype = vpType
	vpDocument.VerifiableCredential = info
	hash := sm3Hash(vpType + info)
	vpDocument.Proof = createProof(priKey, grantFrom, hash)
	byte, err := json.Marshal(vpDocument)
	if err != nil {
		return ""
	} else {
		document := string(byte)
		document = strings.ReplaceAll(document, "\"[", "[")
		document = strings.ReplaceAll(document, "]\"", "]")
		document = strings.ReplaceAll(document, "\\\"", "\"")
		return document
	}
}

func createVerificationMethod(prikey, did string) *VerificationMethod {
	// Initialization information
	verificationMethod := new(VerificationMethod)
	verificationMethod.Id = did
	verificationMethod.VMtype = "Secp256k1"
	verificationMethod.Controller = did
	verificationMethod.PublicKeyHex = priToPub(toPrikey(prikey))

	return verificationMethod
}

func toPrikey(prikey string) *sm2.PrivateKey {
	// Private key string conversion type
	pBytes, err := hex.DecodeString(prikey)
	if err != nil {
		return nil
	}
	pk := sm2.NewPrivateKey(pBytes)
	return pk
}

func priToPub(pk *sm2.PrivateKey) string {
	pubByte := sm2.Compress(&pk.PublicKey)
	return hex.EncodeToString(pubByte)
}

// _createProof 生成签名
func createProof(prikey, did, hash string) *Proof {
	// Private key string conversion type
	pk := toPrikey(prikey)
	pubHex := priToPub(pk)
	// Initialization information
	proof := new(Proof)
	proof.ProofType = "Secp256k1"
	proof.Creator = did
	proof.PublicKeyHex = pubHex
	// sign
	signByte, err := pk.Sign(rand.Reader, []byte(hash), nil)
	if err != nil {
		return nil
	} else {
		proof.SignatureValue = hex.EncodeToString(signByte)
		return proof
	}
}

func getCurrTZTime() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}

func getExpirationDate(days int) string {
	return time.Now().AddDate(0, 0, days).Format("2006-01-02T15:04:05Z")
}

func sm3Hash(msg string) string {
	e := sm3.New()
	/*bf,_ := hex.DecodeString(msg)
	e.Write(bf)*/
	e.Write([]byte(msg))
	return hex.EncodeToString(e.Sum(nil)[:32])
}
