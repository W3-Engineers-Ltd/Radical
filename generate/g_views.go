package generate

import (
	"fmt"
	"os"
	"path"

	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
	"github.com/W3-Engineers-Ltd/Radical/logger/colors"
	"github.com/W3-Engineers-Ltd/Radical/utils"
)

// recipe
// admin/recipe
func GenerateView(viewpath, currpath string) {
	w := colors.NewColorWriter(os.Stdout)

	radicalLogger.Log.Info("Generating view...")

	absViewPath := path.Join(currpath, "views", viewpath)
	err := os.MkdirAll(absViewPath, os.ModePerm)
	if err != nil {
		radicalLogger.Log.Fatalf("Could not create '%s' view: %s", viewpath, err)
	}

	cfile := path.Join(absViewPath, "index.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		f.WriteString(cfile)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		radicalLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "show.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		f.WriteString(cfile)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		radicalLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "create.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		f.WriteString(cfile)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		radicalLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "edit.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		f.WriteString(cfile)
		fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		radicalLogger.Log.Fatalf("Could not create view file: %s", err)
	}
}
