package methods

import "html/template"

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
}
