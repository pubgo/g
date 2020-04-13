package version

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"strings"
)

var (
	// version is the version of project Dragonfly
	// populate via ldflags
	version string

	// revision is the current git commit revision
	// populate via ldflags
	revision string

	// buildDate is the build date of project Dragonfly
	// populate via ldflags
	buildDate string

	// goVersion is the running program's golang version.
	goVersion = runtime.Version()

	// os is the running program's operating system.
	os = runtime.GOOS

	// arch is the running program's architecture target.
	arch = runtime.GOARCH

	// DFDaemonVersion is the version of dfdaemon.
	DFDaemonVersion = version

	// DFGetVersion is the version of dfget.
	DFGetVersion = version

	// SupernodeVersion is the version of supernode.
	SupernodeVersion = version
)

// versionInfoTmpl contains the template used by Info.
var versionInfoTmpl = `
{{.program}} version  {{.version}}
  Git commit:     {{.revision}}
  Build date:     {{.buildDate}}
  Go version:     {{.goVersion}}
  OS/Arch:        {{.OS}}/{{.Arch}}
`

// Print returns version information.
func Print(program string) string {
	m := map[string]string{
		"program":   program,
		"version":   version,
		"revision":  revision,
		"buildDate": buildDate,
		"goVersion": goVersion,
		"OS":        os,
		"Arch":      arch,
	}
	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// Handler returns build information.
func Handler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(Print(""))
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding JSON: %s", err), http.StatusInternalServerError)
	} else if _, err := w.Write(data); err != nil {
		http.Error(w, fmt.Sprintf("error writing the data to the connection: %s", err), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// HandlerWithCtx returns build information.
func HandlerWithCtx(context context.Context, w http.ResponseWriter, r *http.Request) (err error) {
	Handler(w, r)
	return
}
