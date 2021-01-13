package configurators

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var vscodeConfigPath string = ".vscode/"
var vscodeConfigFilename string = "settings.json"

func insertFormOnSave() error {
	os.MkdirAll(vscodeConfigPath, 0755)
	file, err := ioutil.ReadFile(vscodeConfigPath + vscodeConfigFilename)
	var resultFile map[string]interface{} = map[string]interface{}{}

	if err == nil {

		err = json.Unmarshal(file, &resultFile)

		if err != nil {
			return err
		}
	}

	resultFile["editor.formatOnSave"] = true

	resultData, err := json.Marshal(resultFile)

	if err != nil {
		return err
	}

	fileOpen, err := os.Create(vscodeConfigPath + vscodeConfigFilename)

	if err != nil {
		return err
	}

	defer fileOpen.Close()
	fileOpen.Write(resultData)

	return nil
}

// ConfigureVSCode Configure VSCode
func ConfigureVSCode() {
	insertFormOnSave()
}
