package cryptor

import (
	"golang.org/x/crypto/bcrypt"
)

type Cryptor interface {
	EncryptKeyword(pass string) (string, error)
	CompareHashAndPassword(hash, password string) error
}

type cryptor struct {
	pool *workerPool
}

func New(asyncHashingLimit int) Cryptor {
	return &cryptor{
		pool: NewWorkerPool(asyncHashingLimit),
	}
}

func (c cryptor) EncryptKeyword(pass string) (string, error) {
	resChan := make(chan string, 1)
	errChan := make(chan error, 1)

	c.pool.Add(func() {
		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			errChan <- err
			return
		}
		resChan <- string(hash)
	})

	select {
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return "", err
	}
}

func (c cryptor) CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
