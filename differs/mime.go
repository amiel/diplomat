package differs

import (
	"encoding/json"
	"mime"

	"github.com/google/go-cmp/cmp"
)

func diffBody(expected []byte, actual []byte, contentType string) (string, error) {
	mediaType, _, _ := mime.ParseMediaType(contentType)

	switch mediaType {
	case "application/json":
		return diffJson(expected, actual)
	default:
		return "", nil
		// TODO(schoon) - Once we have a "strict" mode:
		// return "", errors.New(fmt.Sprintf("Unsupported media type: %s", mediaType))
	}
}

func diffJson(expected []byte, actual []byte) (string, error) {
	parsedExp := make(map[string]interface{})
	parsedAct := make(map[string]interface{})

	err := json.Unmarshal(expected, &parsedExp)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(actual, &parsedAct)
	if err != nil {
		return "", err
	}

	return cmp.Diff(parsedExp, parsedAct), nil
}