package task

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/common/model"
	"github.com/joewilson27/common/response"
	"github.com/joewilson27/config/db"
)

func AddTask(c *fiber.Ctx) error {
	var newTask TaskAdd

	if err := c.BodyParser(&newTask); err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	svc := Service{
		DB: db.PG,
	}

	resp, errAdd := svc.Add(newTask)
	if errAdd != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.DataAddedFailed + ": " + errAdd.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusAccepted).JSON(model.AppResponse{
		Data: resp,
		Meta: model.Meta{Message: response.DataAddedSuccessfully},
	})
}

func GetTasks(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTasks()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.GetDataFailed + " : " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.GetDataSuccessfully},
	})
}

func GetTasksPaginate(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	// Parse query parameters for pagination and search
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "5"))
	offset := (page - 1) * size
	search := c.Query("search", "")

	// get data
	result, err := svc.GetTasksPaginate(size, offset, search)
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.GetDataFailed + " : " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	// get total data
	totalData, _ := svc.GetTotalData(search)

	// calculate total pages
	totalPages := int(math.Ceil(float64(totalData) / float64(size)))

	// Create the response
	responseData := PaginateResponse{
		TotalData:   totalData,
		TotalPage:   totalPages,
		PageSize:    size,
		CurrentPage: page,
		Data:        result,
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: responseData,
		Meta: model.Meta{Message: response.GetDataSuccessfully},
	})
}

func GetTaskById(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTaskById()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.DataNotFound + " " + err.Error()},
		}

		return c.Status(fiber.StatusNotFound).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.StatusSuccess},
	})
}

func DeleteTask(c *fiber.Ctx) error {
	idp := c.Params("id")
	id, _ := strconv.Atoi(idp)

	svc := Service{
		DB: db.PG,
	}
	fmt.Println(idp)
	if idp == "" {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailToDelete},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	if err := svc.DeleteByID(uint(id)); err != nil {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailToDelete + " " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Meta: model.Meta{Message: response.SuccessDeleted},
	})
}

func UpdateTask(c *fiber.Ctx) error {
	var dataInput TaskAdd
	idp := c.Params("id")
	id, _ := strconv.Atoi(idp)

	if idp == "" {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailedToUpdate},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	if err := c.BodyParser(&dataInput); err != nil {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.ErrorInput + " " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}
	fmt.Println("dataa ", dataInput)
	if err := svc.UpdateByID(uint(id), dataInput); err != nil {
		dataReturnError := model.AppResponse{
			Meta: model.Meta{Message: response.FailedToUpdate + " " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: dataInput,
		Meta: model.Meta{Message: response.SuccessUpdate},
	})
}

func SamplePaginate(c *fiber.Ctx) error {
	// Simulated data
	data := []string{"item1", "item2", "item3", "item4", "item5", "item6", "item7", "item8", "item9", "item10"}

	// Get query parameters for pagination
	page := c.Query("page", "1")
	size := c.Query("size", "5")

	// Convert query parameters to integers
	pageNum, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(size)

	fmt.Println(data)

	// Calculate total pages
	totalPages := int(math.Ceil(float64(len(data)) / float64(pageSize)))

	fmt.Println("totalPages: ", totalPages)

	// Extract data for current page
	start := (pageNum - 1) * pageSize
	end := start + pageSize
	if end > len(data) {
		end = len(data)
	}
	fmt.Println("start: ", start)
	fmt.Println("end: ", end)
	currentPageData := data[start:end]

	// Create the response
	response := DataResponse{
		TotalPage:   totalPages,
		PageSize:    pageSize,
		CurrentPage: pageNum,
		Data:        currentPageData,
	}

	// Return the response as JSON
	return c.JSON(response)
}
