package main

import (
	"net/http"
	"strconv"

	"api-go-gerenciamento-eventos/internal/database"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := app.models.Events.Insert(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao criar novo evento"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro interno ao buscar o evento"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Id de evento inválido"})
	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Id de evento não encontrado"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro interno ao buscar o evento"})
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Id de evento inválido"})
		return
	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Id de evento não encontrado"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro interno ao buscar o evento"})
	}

	updatedEvent := &database.Event{}
	if err := c.ShouldBindJSON(updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	updatedEvent.Id = id
	if err := app.models.Events.Update(updatedEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao atualizar evento"})
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Id de evento inválido"})
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao deletar o evento"})
	}

	c.JSON(http.StatusNoContent, nil)
}
