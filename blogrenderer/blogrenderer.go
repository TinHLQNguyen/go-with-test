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

//go:embed "templates/*"
var postTemplates embed.FS

// Create a dedicated view model type, such as PostViewModel with exactly the data we need to render
type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

func NewPostView(p blogposts.Post) PostViewModel {
	return PostViewModel{
		Title:          p.Title,
		SanitisedTitle: strings.ToLower(strings.ReplaceAll(p.Title, " ", "-")),
		Description:    p.Description,
		Body:           p.Body,
		Tags:           p.Tags,
	}
}

func convertPostsToPostViews(posts []blogposts.Post) []PostViewModel {
	postviews := make([]PostViewModel, len(posts))
	for i, p := range posts {
		postviews[i] = NewPostView(p)
	}
	return postviews
}

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	pv := NewPostView(p)
	if err := templ.ExecuteTemplate(w, "blog.gohtml", pv); err != nil {
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
	pv := NewPostView(p)
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", struct {
		P        PostViewModel
		HtmlBody template.HTML
	}{P: pv, HtmlBody: htmlBody}); err != nil {
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
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitisedTitle}}">{{.Title}}</a></li>{{end}}</ol>`

	pviews := convertPostsToPostViews(posts)
	templ, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, pviews); err != nil {
		return err
	}

	return nil
}
