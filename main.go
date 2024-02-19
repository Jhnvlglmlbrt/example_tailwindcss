package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Task struct {
	Title string
	Done  bool
}

type TaskPageData struct {
	PageTitle string
	Tasks     []Task
}

func tasks(c *fiber.Ctx) error {
	data := TaskPageData{
		PageTitle: "Todays Tasks",
		Tasks: []Task{
			{Title: "Make tiktok video about SSR and Go", Done: false},
			{Title: "Integrate golang SSR and Tailwinds CSS", Done: true},
			{Title: "Carry whittington on fortnite", Done: true},
		},
	}
	return c.Render("tasks", data)
}

func main() {

	engine := html.New("./web/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./web/styles")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Name": "Programming with God",
		})
	})

	app.Get("/tasks", tasks)

	app.Listen(":8080")
}

// TODO:
// add Parcel
// add templ +
// add htmx +
