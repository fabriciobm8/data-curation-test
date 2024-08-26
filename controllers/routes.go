package controllers

import (
    "data-curation-test/service"
    "github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, classMaterialService service.ClassMaterialService, transcriptTimeService service.TranscriptTimeService) {
    // Inicializando os controladores
    classMaterialController := NewClassMaterialController(classMaterialService)
    transcriptTimeController := NewTranscriptTimeController(transcriptTimeService)

    // Rotas para ClassMaterial
    e.POST("/class-material", classMaterialController.Create)
    e.GET("/class-material", classMaterialController.FindAll)
    e.GET("/class-material/:id", classMaterialController.FindByID)
    e.PUT("/class-material/:id", classMaterialController.Update)
    e.DELETE("/class-material/:id", classMaterialController.Delete)

    // Rotas para TranscriptTime
    e.POST("/transcript-time", transcriptTimeController.Create)
    e.GET("/transcript-time", transcriptTimeController.FindAll)
    e.GET("/transcript-time/:id", transcriptTimeController.FindByID)
    e.PUT("/transcript-time/:id", transcriptTimeController.Update)
    e.DELETE("/transcript-time/:id", transcriptTimeController.Delete)
}
