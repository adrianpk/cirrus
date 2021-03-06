package cirrus

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"time"
)

type (
	// App description
	App struct {
		name     string
		revision string

		// Service
		Service Service

		// CQRS
		Commands CommandReg
		//QueryHandlers  QueryHandlerReg

		// Misc
		cancel context.CancelFunc
	}
)

func NewApp(name string) *App {
	name = genName(name, "app")

	return &App{
		name:     name,
		Commands: newCommandReg(),
	}
}

func (app *App) Name() string {
	return app.name
}

func genName(name, defName string) string {
	if strings.Trim(name, " ") == "" {
		return fmt.Sprintf("%s-%s", defName, nameSufix())
	}
	return name
}

func nameSufix() string {
	digest := hash(time.Now().String())
	return digest[len(digest)-8:]
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}

func (a *App) AddCommand(command Command) {
	a.Commands.Add(command)
}
