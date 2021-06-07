package handlers

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type anyDirs interface {
	len() int
	name(i int) string
	isDir(i int) bool
	size(i int) int64
}
type dirEntryDirs []fs.DirEntry

func (d dirEntryDirs) len() int          { return len(d) }
func (d dirEntryDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d dirEntryDirs) name(i int) string { return d[i].Name() }
func (d dirEntryDirs) size(i int) int64 {
	info, _ := d[i].Info()
	return info.Size()
}

//func (d dirEntryDirs) info(i int) (fs.FileInfo, error) { return d[i].Info() }

type fileInfoDirs []fs.FileInfo

func (d fileInfoDirs) len() int          { return len(d) }
func (d fileInfoDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d fileInfoDirs) name(i int) string { return d[i].Name() }
func (d fileInfoDirs) size(i int) int64  { return d[i].Size() }

type DownloadHandler struct {
	Root http.FileSystem
}

func (f *DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	file, err := f.Root.Open(path.Clean(upath))
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer file.Close()

	dirList(w, r, file)
}

func dirList(w http.ResponseWriter, r *http.Request, f http.File) {
	// Prefer to use ReadDir instead of Readdir,
	// because the former doesn't require calling
	// Stat on every entry of a directory on Unix.
	var dirs anyDirs
	var err error
	if d, ok := f.(fs.ReadDirFile); ok {
		var list dirEntryDirs
		list, err = d.ReadDir(-1)
		dirs = list
	} else {
		var list fileInfoDirs
		list, err = f.Readdir(-1)
		dirs = list
	}

	if err != nil {
		logf(r, "http: error reading directory: %v", err)
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs.name(i) > dirs.name(j) })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<table>\n")
	fmt.Fprintf(w, "\t<tr>\n")
	fmt.Fprintf(w, "\t\t<td>Имя</td>\n")
	fmt.Fprintf(w, "\t\t<td>Расширение</td>\n")
	fmt.Fprintf(w, "\t\t<td>Размер в байтах</td>\n")
	fmt.Fprintf(w, "\t</tr>\n")
	for i, n := 0, dirs.len(); i < n; i++ {
		name, ext := splitName(dirs.name(i))
		queryExtension := r.URL.Query().Get("extension")
		if "."+queryExtension != ext && queryExtension != "" {
			continue
		}
		size := ""
		if dirs.isDir(i) {
			name += ext + "/"
			ext = "каталог"
		} else {
			size = strconv.FormatInt(dirs.size(i), 10)
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		url := url.URL{Path: name}
		fmt.Fprintf(w, "\t<tr>\n")
		fmt.Fprintf(w, "\t\t<td><a href=\"%s\">%s</a></td>\n", url.String(), htmlReplacer.Replace(name))
		fmt.Fprintf(w, "\t\t<td>%s</td>\n", ext)
		fmt.Fprintf(w, "\t\t<td>%s</td>\n", size)
		fmt.Fprintf(w, "\t</tr>\n")
		// fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), htmlReplacer.Replace(name))
	}
	fmt.Fprintf(w, "</table>\n")
}

func toHTTPError(err error) (msg string, httpStatus int) {
	if os.IsNotExist(err) {
		return "404 page not found", http.StatusNotFound
	}
	if os.IsPermission(err) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}
func logf(r *http.Request, format string, args ...interface{}) {
	s, _ := r.Context().Value(http.ServerContextKey).(*http.Server)
	if s != nil && s.ErrorLog != nil {
		s.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

func splitName(nameIn string) (name, extension string) {
	lastDot := strings.LastIndex(nameIn, ".")
	if lastDot == -1 {
		lastDot = len(nameIn)
	}
	return string([]byte(nameIn)[:lastDot]), string([]byte(nameIn)[lastDot:])
}
