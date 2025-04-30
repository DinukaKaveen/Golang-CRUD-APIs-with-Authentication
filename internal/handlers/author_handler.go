package handlers

import (
	"database/sql"
	"goapp/internal/database"
	"goapp/internal/database/sqlc"

	"github.com/gofiber/fiber/v2"
)

var db = database.ConnectDB()
var queries = sqlc.New(db)

type Author struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

func CreateAuthor(c *fiber.Ctx) error {
	var author Author
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newAuthor, err := queries.CreateAuthor(c.Context(), sqlc.CreateAuthorParams{
		FirstName: author.FirstName,
		LastName : author.LastName,
		Bio      : sql.NullString{String: author.Bio, Valid: author.Bio != ""},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create author",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newAuthor,
		"message": "Successfully created the Author",
	})
}

func UpdateAuthor(c *fiber.Ctx) error {
	var author Author
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	_, err = queries.GetAuthor(c.Context(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Author not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve author",
		})
	}

	updatedAuthor, err := queries.UpdateAuthor(c.Context(), sqlc.UpdateAuthorParams{
		ID       : int64(id),
		FirstName: author.FirstName,
		LastName : author.LastName,
		Bio      : sql.NullString{String: author.Bio, Valid: author.Bio != ""},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update author",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": updatedAuthor,
		"message": "Successfully updated the Author",
	})
}

func GetAuthor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	author, err := queries.GetAuthor(c.Context(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Author not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve author",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": author,
		"message": "Successfully retrieved the Author",
	})
}

func ListAuthors(c *fiber.Ctx) error {
	authors, err := queries.ListAuthors(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": authors,
		"message": "Successfully retrieved the Authors",
	})
}

func DeleteAuthor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	_, err = queries.GetAuthor(c.Context(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Author not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve author",
		})
	}

	err = queries.DeleteAuthor(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete author",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Author deleted successfully",
	})
}
