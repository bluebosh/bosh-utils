package crypto

import (
	"fmt"
	"io"

	bosherr "github.com/bluebosh/bosh-utils/errors"
	boshsys "github.com/bluebosh/bosh-utils/system"
	"os"
)

type digestImpl struct {
	algorithm Algorithm
	digest    string
}

func NewDigest(algorithm Algorithm, digest string) digestImpl {
	return digestImpl{
		algorithm: algorithm,
		digest:    digest,
	}
}

func (c digestImpl) Algorithm() Algorithm { return c.algorithm }

func (c digestImpl) String() string {
	if c.algorithm.Name() == DigestAlgorithmSHA1.Name() {
		return c.digest
	}

	return fmt.Sprintf("%s:%s", c.algorithm.Name(), c.digest)
}

func (c digestImpl) Verify(reader io.Reader) error {
	computedDigest, err := c.Algorithm().CreateDigest(reader)
	if err != nil {
		return bosherr.WrapError(err, "Computing digest from stream")
	}

	if c.String() != computedDigest.String() {
		return bosherr.Errorf("Expected stream to have digest '%s' but was '%s'", c.String(), computedDigest.String())
	}

	return nil
}

func (m digestImpl) VerifyFilePath(filePath string, fs boshsys.FileSystem) error {
	file, err := fs.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return bosherr.WrapErrorf(err, "Calculating digest of '%s'", filePath)
	}
	defer func() {
		_ = file.Close()
	}()
	return m.Verify(file)
}
