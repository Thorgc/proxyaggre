package getter

import (
	"errors"
	"github.com/oouxx/proxyaggre/pkg/tool"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestParseClash(t *testing.T) {
	path := "https://subcon.dlj.tf/sub?target=clash&new_name=true&url=https%3A%2F%2Fwangxinxing.ml%2Fvmess%2Fsub&insert=false&config=https%3A%2F%2Fraw.githubusercontent.com%2FACL4SSR%2FACL4SSR%2Fmaster%2FClash%2Fconfig%2FACL4SSR_Online.ini"
	data, err := readFile(path)
	if err != nil{
		return
	}
	ParseClashProxyFromYamlConfig(data)
}
func readFile(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := tool.GetHttpClient().Get(path)
		if err != nil {
			return nil, errors.New("config file http get fail")
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	} else {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, err
		}
		return ioutil.ReadFile(path)
	}
}
