package radiantpro

import (
	"io/ioutil"

	"github.com/W3-Engineers-Ltd/Radical/internal/pkg/utils"
	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
)

var CompareExcept = []string{"@radicalGenerateTime"}

func (c *Container) GenConfig() {
	if utils.IsExist(c.radiantProFile) {
		radicalLogger.Log.Fatalf("radiant pro toml exist")
		return
	}

	err := ioutil.WriteFile("radiantpro.toml", []byte(radiantToml), 0644)
	if err != nil {
		radicalLogger.Log.Fatalf("write radiant pro toml err: %s", err)
		return
	}
}
