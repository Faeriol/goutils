package configutil

import (
	"encoding/json"
	"os"
	"path"
)

func WriteConfFile(v interface{}, file string) error {
	os.MkdirAll(path.Dir(file), 0755)

	fh, err := os.Create(file)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(v, "", "    ")
	if nil != err {
		return err
	}

	_, err = fh.Write(data)
	return err

}
