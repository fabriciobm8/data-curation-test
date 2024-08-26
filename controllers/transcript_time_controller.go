package controllers

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type TranscriptTimeController struct {
    service service.TranscriptTimeService
}

func NewTranscriptTimeController(service service.TranscriptTimeService) *TranscriptTimeController {
    return &TranscriptTimeController{service: service}
}

func (c *TranscriptTimeController) Create(ctx echo.Context) error {
    var transcriptTime models.TranscriptTime
    if err := ctx.Bind(&transcriptTime); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Create(context.Background(), &transcriptTime)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, transcriptTime)
}

func (c *TranscriptTimeController) FindAll(ctx echo.Context) error {
    transcriptTime, err := c.service.FindAll(context.Background())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, transcriptTime)
}

func (c *TranscriptTimeController) FindByID(ctx echo.Context) error {
    id := ctx.Param("id")
    transcriptTime, err := c.service.FindByID(context.Background(), id)
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.JSON(http.StatusOK, transcriptTime)
}

func (c *TranscriptTimeController) Update(ctx echo.Context) error {
    id := ctx.Param("id")
    var transcriptTime models.TranscriptTime
    if err := ctx.Bind(&transcriptTime); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Update(context.Background(), id, &transcriptTime)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, transcriptTime)
}

func (c *TranscriptTimeController) Delete(ctx echo.Context) error {
    id := ctx.Param("id")
    err := c.service.Delete(context.Background(), id)
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.NoContent(http.StatusNoContent)
}
