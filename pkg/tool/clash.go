package tool
import (
	//"github.com/Dreamacro/clash/adapters/provider"
	"github.com/Dreamacro/clash/common/structure"
	//"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
)
func parseClash(src io.Reader) {
	yaml.NewDecoder(src).Decode(structure.Decoder{})
}