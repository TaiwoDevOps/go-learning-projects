package utils

// handle Marshalling and Un-marshalling (handling json)
import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), v)
	if err != nil {
		return err
	}
	return nil
}

/**!SECTION
func ParseBody(r *http.Request, v interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), v); err != nil {
			log.Fatalln(err)
			return
		}
	}
}
*/
