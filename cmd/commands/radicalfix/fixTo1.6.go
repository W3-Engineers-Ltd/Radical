package radicalfix

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/W3-Engineers-Ltd/Radical/cmd/commands"
	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
	"github.com/W3-Engineers-Ltd/Radical/logger/colors"
)

// fixTo16 upgrade radiant to 1.6
func fixTo16(cmd *commands.Command, args []string) int {
	output := cmd.Out()

	radicalLogger.Log.Info("Upgrading the application...")

	dir, err := os.Getwd()
	if err != nil {
		radicalLogger.Log.Fatalf("Error while getting the current working directory: %s", err)
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".exe") {
			return nil
		}
		err = fixFile(path)
		fmt.Fprintf(output, colors.GreenBold("\tfix\t")+"%s\n", path)
		if err != nil {
			radicalLogger.Log.Errorf("Could not fix file: %s", err)
		}
		return err
	})
	radicalLogger.Log.Success("Upgrade Done!")
	return 0
}

var rules = []string{
	"radiant.AppName", "radiant.BConfig.AppName",
	"radiant.RunMode", "radiant.BConfig.RunMode",
	"radiant.RecoverPanic", "radiant.BConfig.RecoverPanic",
	"radiant.RouterCaseSensitive", "radiant.BConfig.RouterCaseSensitive",
	"radiant.radiantServerName", "radiant.BConfig.ServerName",
	"radiant.EnableGzip", "radiant.BConfig.EnableGzip",
	"radiant.ErrorsShow", "radiant.BConfig.EnableErrorsShow",
	"radiant.CopyRequestBody", "radiant.BConfig.CopyRequestBody",
	"radiant.MaxMemory", "radiant.BConfig.MaxMemory",
	"radiant.Graceful", "radiant.BConfig.Listen.Graceful",
	"radiant.HttpAddr", "radiant.BConfig.Listen.HTTPAddr",
	"radiant.HttpPort", "radiant.BConfig.Listen.HTTPPort",
	"radiant.ListenTCP4", "radiant.BConfig.Listen.ListenTCP4",
	"radiant.EnableHttpListen", "radiant.BConfig.Listen.EnableHTTP",
	"radiant.EnableHttpTLS", "radiant.BConfig.Listen.EnableHTTPS",
	"radiant.HttpsAddr", "radiant.BConfig.Listen.HTTPSAddr",
	"radiant.HttpsPort", "radiant.BConfig.Listen.HTTPSPort",
	"radiant.HttpCertFile", "radiant.BConfig.Listen.HTTPSCertFile",
	"radiant.HttpKeyFile", "radiant.BConfig.Listen.HTTPSKeyFile",
	"radiant.EnableAdmin", "radiant.BConfig.Listen.EnableAdmin",
	"radiant.AdminHttpAddr", "radiant.BConfig.Listen.AdminAddr",
	"radiant.AdminHttpPort", "radiant.BConfig.Listen.AdminPort",
	"radiant.UseFcgi", "radiant.BConfig.Listen.EnableFcgi",
	"radiant.HttpServerTimeOut", "radiant.BConfig.Listen.ServerTimeOut",
	"radiant.AutoRender", "radiant.BConfig.WebConfig.AutoRender",
	"radiant.ViewsPath", "radiant.BConfig.WebConfig.ViewsPath",
	"radiant.StaticDir", "radiant.BConfig.WebConfig.StaticDir",
	"radiant.StaticExtensionsToGzip", "radiant.BConfig.WebConfig.StaticExtensionsToGzip",
	"radiant.DirectoryIndex", "radiant.BConfig.WebConfig.DirectoryIndex",
	"radiant.FlashName", "radiant.BConfig.WebConfig.FlashName",
	"radiant.FlashSeperator", "radiant.BConfig.WebConfig.FlashSeparator",
	"radiant.EnableDocs", "radiant.BConfig.WebConfig.EnableDocs",
	"radiant.XSRFKEY", "radiant.BConfig.WebConfig.XSRFKey",
	"radiant.EnableXSRF", "radiant.BConfig.WebConfig.EnableXSRF",
	"radiant.XSRFExpire", "radiant.BConfig.WebConfig.XSRFExpire",
	"radiant.TemplateLeft", "radiant.BConfig.WebConfig.TemplateLeft",
	"radiant.TemplateRight", "radiant.BConfig.WebConfig.TemplateRight",
	"radiant.SessionOn", "radiant.BConfig.WebConfig.Session.SessionOn",
	"radiant.SessionProvider", "radiant.BConfig.WebConfig.Session.SessionProvider",
	"radiant.SessionName", "radiant.BConfig.WebConfig.Session.SessionName",
	"radiant.SessionGCMaxLifetime", "radiant.BConfig.WebConfig.Session.SessionGCMaxLifetime",
	"radiant.SessionSavePath", "radiant.BConfig.WebConfig.Session.SessionProviderConfig",
	"radiant.SessionCookieLifeTime", "radiant.BConfig.WebConfig.Session.SessionCookieLifeTime",
	"radiant.SessionAutoSetCookie", "radiant.BConfig.WebConfig.Session.SessionAutoSetCookie",
	"radiant.SessionDomain", "radiant.BConfig.WebConfig.Session.SessionDomain",
	"Ctx.Input.CopyBody(", "Ctx.Input.CopyBody(radiant.BConfig.MaxMemory",
	".UrlFor(", ".URLFor(",
	".ServeJson(", ".ServeJSON(",
	".ServeXml(", ".ServeXML(",
	".ServeJsonp(", ".ServeJSONP(",
	".XsrfToken(", ".XSRFToken(",
	".CheckXsrfCookie(", ".CheckXSRFCookie(",
	".XsrfFormHtml(", ".XSRFFormHTML(",
	"radiant.UrlFor(", "radiant.URLFor(",
	"radiant.GlobalDocApi", "radiant.GlobalDocAPI",
	"radiant.Errorhandler", "radiant.ErrorHandler",
	"Output.Jsonp(", "Output.JSONP(",
	"Output.Json(", "Output.JSON(",
	"Output.Xml(", "Output.XML(",
	"Input.Uri()", "Input.URI()",
	"Input.Url()", "Input.URL()",
	"Input.AcceptsHtml()", "Input.AcceptsHTML()",
	"Input.AcceptsXml()", "Input.AcceptsXML()",
	"Input.AcceptsJson()", "Input.AcceptsJSON()",
	"Ctx.XsrfToken()", "Ctx.XSRFToken()",
	"Ctx.CheckXsrfCookie()", "Ctx.CheckXSRFCookie()",
	"session.SessionStore", "session.Store",
	".TplNames", ".TplName",
	"swagger.ApiRef", "swagger.APIRef",
	"swagger.ApiDeclaration", "swagger.APIDeclaration",
	"swagger.Api", "swagger.API",
	"swagger.ApiRef", "swagger.APIRef",
	"swagger.Infomation", "swagger.Information",
	"toolbox.UrlMap", "toolbox.URLMap",
	"logs.LoggerInterface", "logs.Logger",
	"Input.Request", "Input.Context.Request",
	"Input.Params)", "Input.Params())",
	"httplib.radiantHttpSettings", "httplib.radiantHTTPSettings",
	"httplib.radiantHttpRequest", "httplib.radiantHTTPRequest",
	".TlsClientConfig", ".TLSClientConfig",
	".JsonBody", ".JSONBody",
	".ToJson", ".ToJSON",
	".ToXml", ".ToXML",
	"radiant.Html2str", "radiant.HTML2str",
	"radiant.AssetsCss", "radiant.AssetsCSS",
	"orm.DR_Sqlite", "orm.DRSqlite",
	"orm.DR_Postgres", "orm.DRPostgres",
	"orm.DR_MySQL", "orm.DRMySQL",
	"orm.DR_Oracle", "orm.DROracle",
	"orm.Col_Add", "orm.ColAdd",
	"orm.Col_Minus", "orm.ColMinus",
	"orm.Col_Multiply", "orm.ColMultiply",
	"orm.Col_Except", "orm.ColExcept",
	"GenerateOperatorSql", "GenerateOperatorSQL",
	"OperatorSql", "OperatorSQL",
	"orm.Debug_Queries", "orm.DebugQueries",
	"orm.COMMA_SPACE", "orm.CommaSpace",
	".SendOut()", ".DoRequest()",
	"validation.ValidationError", "validation.Error",
}

func fixFile(file string) error {
	rp := strings.NewReplacer(rules...)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	fixed := rp.Replace(string(content))

	// Forword the RequestBody from the replace
	// "Input.Request", "Input.Context.Request",
	fixed = strings.Replace(fixed, "Input.Context.RequestBody", "Input.RequestBody", -1)

	// Regexp replace
	pareg := regexp.MustCompile(`(Input.Params\[")(.*)("])`)
	fixed = pareg.ReplaceAllString(fixed, "Input.Param(\"$2\")")
	pareg = regexp.MustCompile(`(Input.Data\[\")(.*)(\"\])(\s)(=)(\s)(.*)`)
	fixed = pareg.ReplaceAllString(fixed, "Input.SetData(\"$2\", $7)")
	pareg = regexp.MustCompile(`(Input.Data\[\")(.*)(\"\])`)
	fixed = pareg.ReplaceAllString(fixed, "Input.Data(\"$2\")")
	// Fix the cache object Put method
	pareg = regexp.MustCompile(`(\.Put\(\")(.*)(\",)(\s)(.*)(,\s*)([^\*.]*)(\))`)
	if pareg.MatchString(fixed) && strings.HasSuffix(file, ".go") {
		fixed = pareg.ReplaceAllString(fixed, ".Put(\"$2\", $5, $7*time.Second)")
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
		if err != nil {
			panic(err)
		}
		// Print the imports from the file's AST.
		hasTimepkg := false
		for _, s := range f.Imports {
			if s.Path.Value == `"time"` {
				hasTimepkg = true
				break
			}
		}
		if !hasTimepkg {
			fixed = strings.Replace(fixed, "import (", "import (\n\t\"time\"", 1)
		}
	}
	// Replace the v.Apis in docs.go
	if strings.Contains(file, "docs.go") {
		fixed = strings.Replace(fixed, "v.Apis", "v.APIs", -1)
	}
	// Replace the config file
	if strings.HasSuffix(file, ".conf") {
		fixed = strings.Replace(fixed, "HttpCertFile", "HTTPSCertFile", -1)
		fixed = strings.Replace(fixed, "HttpKeyFile", "HTTPSKeyFile", -1)
		fixed = strings.Replace(fixed, "EnableHttpListen", "HTTPEnable", -1)
		fixed = strings.Replace(fixed, "EnableHttpTLS", "EnableHTTPS", -1)
		fixed = strings.Replace(fixed, "EnableHttpTLS", "EnableHTTPS", -1)
		fixed = strings.Replace(fixed, "radiantServerName", "ServerName", -1)
		fixed = strings.Replace(fixed, "AdminHttpAddr", "AdminAddr", -1)
		fixed = strings.Replace(fixed, "AdminHttpPort", "AdminPort", -1)
		fixed = strings.Replace(fixed, "HttpServerTimeOut", "ServerTimeOut", -1)
	}
	err = os.Truncate(file, 0)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, []byte(fixed), 0666)
}
