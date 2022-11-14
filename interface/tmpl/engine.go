package tmpl

type Tmpl interface {
	Exec(data interface{})
}
