package blogrenderer

import (
	"embed"
	"go-with-test/blogposts"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	htmlBody := template.HTML(string(mdToHTML([]byte(p.Body))))
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", struct {
		P        blogposts.Post
		HtmlBody template.HTML
	}{P: p, HtmlBody: htmlBody}); err != nil {
		return err
	}

	return nil
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.Replace(title, " ", "-", -1))
		},
	}).Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil
}
