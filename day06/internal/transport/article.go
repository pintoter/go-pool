package transport

import (
	"context"
	"day06/internal/entity"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

const (
	articlesOnPage = 3
	articlesHTML   = "./internal/templates/articles.html"
	articleHTML    = "./internal/templates/article.html"
	minPage        = 1
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
		http.Error(w, entity.ErrEmptyParams.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.service.CreateArticle(context.Background(), title, content)
	if err != nil {
		http.Error(w, entity.ErrFailedCreateArticle.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusCreated)
}

func (h *Handler) articleHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	articleId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, entity.ErrInvalidId.Error(), http.StatusBadRequest)
		return
	}

	article, err := h.service.GetArticle(context.Background(), articleId)
	if err != nil {
		http.Error(w, entity.ErrArticleNotFound.Error(), http.StatusNotFound)
		return
	}

	renderTemplate(w, articleHTML, &articlePage{
		Id:      article.ID,
		Title:   article.Title,
		Content: template.HTML(blackfriday.Run([]byte(article.Content))),
	})
}

func (h *Handler) articlesHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("page")

	var page int = 1
	var err error

	if param != "" {
		page, err = strconv.Atoi(param)
		if err != nil || page < minPage {
			http.Error(w, entity.ErrInvalidPage.Error(), http.StatusBadRequest)
			return
		}
	}

	currentOffset := (page - 1) * 3

	articles, total, err := h.service.GetArticles(context.Background(), articlesOnPage, currentOffset)
	if err != nil {
		http.Error(w, entity.ErrArticlesNotFound.Error(), http.StatusInternalServerError)
		return
	}

	var empty bool
	if len(articles) == 0 {
		renderTemplate(w, articlesHTML, homePage{
			Empty: !empty,
		})
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

	renderTemplate(w, articlesHTML, homePage{
		Articles:    articlePages,
		HasPrevPage: hasPrevPage,
		PrevPage:    page - 1,
		HasNextPage: hasNextPage,
		NextPage:    page + 1,
		Empty:       empty,
	})
}
