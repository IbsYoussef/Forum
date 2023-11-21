package methods

import "html/template"

var Tpl *template.Template

func InitGlobalTemplate() {
	Tpl = template.Must(Tpl.ParseGlob("templates/*.html"))
}
