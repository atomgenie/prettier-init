package configurators

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
)

func getRawData() ([]byte, error) {
	data, err := http.Get("https://gitlab.com/atomgenie/shared/-/raw/master/.prettierrc.js")

	if err != nil {
		os.Stderr.WriteString("Can't download .prettierrc.js\n")
		return nil, err
	}

	return ioutil.ReadAll(data.Body)
}

func saveBytesInFile(data []byte) error {

	prettierFileName := ".prettierrc.js"

	_, err := os.Stat(prettierFileName)

	if err == nil {
		os.Stdout.WriteString(prettierFileName + " already exists\n")
		os.Stdout.WriteString("Override it? (y/n)\n")

		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		response := input.Text()

		if response != "y" {
			return nil
		}
	}

	file, err := os.Create(prettierFileName)

	if err != nil {
		os.Stderr.WriteString("Can't create .prettierrc.js\n")
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	return err
}

// InstallPrettier Install Prettier in current directory
func InstallPrettier() error {

	prettierData, err := getRawData()

	if err != nil {
		return err
	}

	return saveBytesInFile(prettierData)

}
