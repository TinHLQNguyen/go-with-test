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

// Create a dedicated view model type, such as postViewModel with exactly the data we need to render
type postViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
	HtmlBody                                 template.HTML
}

func newPostView(p blogposts.Post, r *PostRenderer) postViewModel {
	htmlBody := template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, r.mdRenderer))
	return postViewModel{
		Title:          p.Title,
		SanitisedTitle: strings.ToLower(strings.ReplaceAll(p.Title, " ", "-")),
		Description:    p.Description,
		Body:           p.Body,
		Tags:           p.Tags,
		HtmlBody:       htmlBody,
	}
}

func convertPostsToPostViews(posts []blogposts.Post, r *PostRenderer) []postViewModel {
	postviews := make([]postViewModel, len(posts))
	for i, p := range posts {
		postviews[i] = newPostView(p, r)
	}
	return postviews
}

type PostRenderer struct {
	templ      *template.Template
	mdParser   *parser.Parser
	mdRenderer *html.Renderer
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	r := html.NewRenderer(opts)

	return &PostRenderer{templ: templ, mdParser: p, mdRenderer: r}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	pv := newPostView(p, r)
	return r.templ.ExecuteTemplate(w, "blog.gohtml", pv)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	pviews := convertPostsToPostViews(posts, r)
	return r.templ.ExecuteTemplate(w, "index.gohtml", pviews)
}
