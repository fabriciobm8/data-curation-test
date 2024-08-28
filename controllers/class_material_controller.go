package controllers

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type ClassMaterialController struct {
    service *service.ClassMaterialService
}

func NewClassMaterialController(service *service.ClassMaterialService) *ClassMaterialController {
    return &ClassMaterialController{service: service}
}

func (c *ClassMaterialController) Create(ctx echo.Context) error {
    var classMaterial models.ClassMaterial
    if err := ctx.Bind(&classMaterial); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Create(context.Background(), &classMaterial)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, classMaterial)
}

func (c *ClassMaterialController) FindAll(ctx echo.Context) error {
    classMaterials, err := c.service.FindAll(context.Background())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterials)
}

func (c *ClassMaterialController) FindByID(ctx echo.Context) error {
    id := ctx.Param("id")
    classMaterial, err := c.service.FindByID(context.Background(), id)
    if err != nil {
        if err.Error() == "id é obrigatório" {
            return ctx.JSON(http.StatusBadRequest, err.Error())
        }
        if err.Error() == "classMaterial não encontrado" {
            return ctx.JSON(http.StatusNotFound, err.Error())
        }
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterial)
}

func (c *ClassMaterialController) Update(ctx echo.Context) error {
    id := ctx.Param("id")
    var classMaterial models.ClassMaterial
    if err := ctx.Bind(&classMaterial); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Update(context.Background(), id, &classMaterial)
    if err != nil {
        if err.Error() == "id é obrigatório" {
            return ctx.JSON(http.StatusBadRequest, err.Error())
        }
        if err.Error() == "classMaterial não encontrado" {
            return ctx.JSON(http.StatusNotFound, err.Error())
        }
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterial)
}

func (c *ClassMaterialController) Delete(ctx echo.Context) error {
    id := ctx.Param("id")
    err := c.service.Delete(context.Background(), id)
    if err != nil {
        if err.Error() == "id é obrigatório" {
            return ctx.JSON(http.StatusBadRequest, err.Error())
        }
        if err.Error() == "classMaterial não encontrado" {
            return ctx.JSON(http.StatusNotFound, err.Error())
        }
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.NoContent(http.StatusNoContent)
}

func (c *ClassMaterialController) UpdateIsSuccessfulClassMaterial(ctx echo.Context) error {
    var classMaterial models.ClassMaterial
    var message struct {
        Message string `json:"message"`
    }

    // Bind JSON request body to ClassMaterial
    if err := ctx.Bind(&classMaterial); err != nil {
        message.Message = err.Error()
        return ctx.JSON(http.StatusBadRequest, message)
    }

    // Check for ID in the URL parameter
    id := ctx.Param("id")
    if id == "" {
        message.Message = "ID é obrigatório"
        return ctx.JSON(http.StatusBadRequest, message)
    }
    classMaterial.ID = id

    // Check for uuidCourse in the query parameter or request body
    uuidCourse := ctx.QueryParam("uuidCourse")
    if uuidCourse == "" {
        uuidCourse = classMaterial.UuidCourse
    }
    if uuidCourse == "" {
        message.Message = "uuidCourse é obrigatório"
        return ctx.JSON(http.StatusBadRequest, message)
    }
    classMaterial.UuidCourse = uuidCourse

    // Call the service method to update evaluation
    if err := c.service.UpdateIsSuccessful(context.Background(), &classMaterial, classMaterial.IsSuccessful); err != nil {
        message.Message = err.Error()
        return ctx.JSON(http.StatusInternalServerError, message)
    }

    message.Message = "Successfully updated"
    return ctx.JSON(http.StatusOK, message)
}
