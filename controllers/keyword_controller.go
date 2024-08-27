package controllers

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type KeywordController struct {
    service service.KeywordService
}

func NewKeywordController(service service.KeywordService) *KeywordController {
    return &KeywordController{service: service}
}

func (c *KeywordController) Create(ctx echo.Context) error {
    var keyword models.Keyword
    if err := ctx.Bind(&keyword); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Create(context.Background(), &keyword)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, keyword)
}

func (c *KeywordController) FindAll(ctx echo.Context) error {
    keywords, err := c.service.FindAll(context.Background())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, keywords)
}

func (c *KeywordController) FindByID(ctx echo.Context) error {
    id := ctx.Param("id")
    keyword, err := c.service.FindByID(context.Background(), id)
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.JSON(http.StatusOK, keyword)
}

func (c *KeywordController) Update(ctx echo.Context) error {
    id := ctx.Param("id")
    var keyword models.Keyword
    if err := ctx.Bind(&keyword); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Update(context.Background(), id, &keyword)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, keyword)
}

func (c *KeywordController) Delete(ctx echo.Context) error {
    id := ctx.Param("id")
    err := c.service.Delete(context.Background(), id)
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.NoContent(http.StatusNoContent)
}
