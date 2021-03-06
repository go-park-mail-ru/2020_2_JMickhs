package csrfUsecase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/sessions/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/sessions/internal/csrf"
)

type CsrfUsecase struct {
	CsrfRep csrf.Repository
}

type CsrfToken struct {
	SessionID string
	Timestamp int64
}

func NewCsrfUsecase(r csrf.Repository) *CsrfUsecase {
	return &CsrfUsecase{
		CsrfRep: r,
	}
}

func (u *CsrfUsecase) CreateToken(sesID string, timeStamp int64) (string, error) {

	block, err := aes.NewCipher([]byte(configs.SecretTokenKey))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	td := &CsrfToken{SessionID: sesID, Timestamp: timeStamp}
	data, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)

	return token, nil

}

func (u *CsrfUsecase) CheckToken(sesID string, token string) (bool, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(token)

	block, err := aes.NewCipher([]byte(configs.SecretTokenKey))
	if err != nil {
		return false, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, errors.New("ciphertext < nonceSize")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, err
	}

	CsrfTok := CsrfToken{}
	err = json.Unmarshal(plaintext, &CsrfTok)
	if err != nil {
		return false, err
	}

	if time.Now().Unix()-CsrfTok.Timestamp >
		int64(time.Duration(viper.GetInt(configs.ConfigFields.CsrfExpire))*time.Minute) {
		return false, errors.New("token expired")
	}

	expected := CsrfToken{SessionID: sesID, Timestamp: CsrfTok.Timestamp}

	err = u.CsrfRep.Check(token)

	if CsrfTok != expected || err != nil {
		return false, nil
	}

	err = u.CsrfRep.Add(token)
	if err != nil {
		return false, err
	}
	return true, nil
}
