package methods

import "html/template"

var Tpl *template.Template

func Init() {
	Tpl = template.Must(Tpl.ParseGlob("templates/*.html"))
}
