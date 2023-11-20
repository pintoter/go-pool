package transport

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

type articlePage struct {
	Id      int
	Title   string
	Content template.HTML
}

type homePage struct {
	Articles    []articlePage
	HasPrevPage bool
	PrevPage    int
	HasNextPage bool
	NextPage    int
	Empty       bool
}

func (h *Handler) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	if !h.checkLoginState(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" || content == "" {
		http.Error(w, "params can't be empty", http.StatusBadRequest)
		return
	}

	_, err := h.service.CreateArticle(context.Background(), title, content)
	if err != nil {
		http.Error(w, "failed to create article", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusCreated)
}

func (h *Handler) articleHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	articleId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	article, err := h.service.GetArticle(context.Background(), articleId)
	if err != nil {
		http.Error(w, "article not found", http.StatusNotFound)
		return
	}

	renderTemplate(w, "internal/templates/article.html", &articlePage{
		Id:      article.ID,
		Title:   article.Title,
		Content: template.HTML(blackfriday.Run([]byte(article.Content))),
	})
}

func (h *Handler) articlesHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("page")

	page := 0
	var err error

	if param != "" {
		page, err = strconv.Atoi(param)
		if err != nil {
			http.Error(w, "uncorrect page", http.StatusBadRequest)
		}
	}

	articles, total, err := h.service.GetArticles(context.Background(), page*3)
	if err != nil {
		http.Error(w, "Failed to get articles", http.StatusInternalServerError)
		return
	}

	empty := false
	if page == 0 && len(articles) < 1 {
		empty = true
	}

	if page < 0 {
		http.Error(w, "Page is out of range", http.StatusInternalServerError)
		return
	}

	articlePages := make([]articlePage, 0)
	for _, article := range articles {
		articlePages = append(articlePages, articlePage{
			Id:      article.ID,
			Title:   article.Title,
			Content: template.HTML(blackfriday.Run([]byte(article.Content))),
		})
	}

	hasPrevPage := page > 0
	hasNextPage := (total-1)/(page+1)/3 > 0

	renderTemplate(w, "internal/templates/articles.html", &homePage{
		Articles:    articlePages,
		HasPrevPage: hasPrevPage,
		PrevPage:    page - 1,
		HasNextPage: hasNextPage,
		NextPage:    page + 1,
		Empty:       empty,
	})
}
