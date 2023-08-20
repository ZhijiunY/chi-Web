package handlers

import (
	"log"
	"net/http"

	"github.com/ZhijiunY/chi-web/internal/config"
	"github.com/ZhijiunY/chi-web/internal/forms"
	"github.com/ZhijiunY/chi-web/internal/render"
	"github.com/ZhijiunY/chi-web/models"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// GET
////////////////////////////////

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	m.App.Session.Put(r.Context(), "userid", "derekbanas")
	render.RenderTemplate(w, r, "home.tmpl", &models.PageData{})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "about.tmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "login.tmpl", &models.PageData{StrMap: strMap})

}

func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {

	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	render.RenderTemplate(w, r, "make-post.tmpl", &models.PageData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {

	// // create a empty article
	// var emptyArticle models.Article
	// data := make(map[string]interface{})
	// data["article"] = emptyArticle

	// render.RenderTemplate(w, r, "make-post.tmpl",
	// 	&models.PageData{
	// 		Form: forms.New(nil),
	// 		Data: data,
	// 	})

	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "page.tmpl", &models.PageData{StrMap: strMap})

}

// POST
////////////////////////////////

func (m *Repository) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "PostMakePostHandler Error", http.StatusInternalServerError)
		return
	}

	article := models.Article{
		BlogTitle:   r.Form.Get("blog_title"),
		BlogArticle: r.Form.Get("blog_article"),
	}

	form := forms.New(r.PostForm)
	// 檢查表單不同欄位
	form.HasRequired("blog_title", "blog_article")

	// 驗證表單不同字段
	if !form.MinLength("blog_title", 5, r) {
		log.Fatal("can't get blog title")
	}
	if !form.MinLength("blog_article", 5, r) {
		log.Fatal("can't set min length")
	}

	// // check if email is valid
	// form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article

		render.RenderTemplate(w, r, "make-post.tmpl", &models.PageData{
			Form: form,
			Data: data,
		})
		return
	}
	// use session to store data
	m.App.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received", http.StatusSeeOther)
}

// receive article page
func (m *Repository) ArticleReceived(w http.ResponseWriter, r *http.Request) {

	article, ok := m.App.Session.Get(r.Context(), "article").(models.Article)
	if !ok {
		log.Println("Can't get data from session")
		http.Error(w, "ArticleReceived Error", http.StatusInternalServerError)

		m.App.Session.Put(r.Context(), "error", "Can't get data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	// store the article data
	data := make(map[string]interface{})
	data["article"] = article

	render.RenderTemplate(w, r, "article-received.tmpl", &models.PageData{
		Data: data,
	})
}
