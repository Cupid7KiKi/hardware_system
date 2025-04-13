package pkg

import "html/template"

type BaseComponent interface {
	GetContent() template.HTML
}
