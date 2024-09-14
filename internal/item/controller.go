package item

import (
	"fmt"
	"midterm-api/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Service: NewService(db),
	}
}

//Validation
type ApiError struct {
	Field  string
	Reason string
}

func msgForTag(tag, param string) string {
	switch tag {
	case "required":
		return "Please fill out required field"
	case "gt":
		return fmt.Sprintf("Number must greater than %v", param)
	}
	return ""
}

func getValidationErrors(err error) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), msgForTag(fe.Tag(), fe.Param())}
		}
		return out
	}
	return nil
}


func (controller Controller) CreateItem(ctx *gin.Context) {

	var request model.RequestItem

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": getValidationErrors(err),
		})
		return
	}

	item, err := controller.Service.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
}

func (controller Controller) FindItems(ctx *gin.Context) {
	var (
		request model.RequestFindItem
	)
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": err,
		})
		return
	}
	items, err := controller.Service.Find(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"data": items,
	})
}

func (controller Controller) UpdateItemStatus(ctx *gin.Context) {

	var (
		request model.RequestUpdateItem
	)
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": err,
		})
		return
	}

	// id, _ := strconv.Atoi(ctx.Param("id"))
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	item, err := controller.Service.UpdateStatus(uint(id), request.Status)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": err,
        })
        return
    }
	
    ctx.JSON(http.StatusOK, gin.H{
        "data": item,
    })
}

func (controller Controller) FindItemById(ctx *gin.Context) {
	var (
		request model.RequestFindItem
	)
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": err,
		})
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	items, err := controller.Service.FindbyId(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"data": items,
	})
}

func (controller Controller) DeleteItemById(ctx *gin.Context) {

	var (
		request model.RequestFindItem
	)
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": err,
		})
		return
	}

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	item, err := controller.Service.DeleteByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": err,
        })
        return
    }
	
    ctx.JSON(http.StatusOK, gin.H{
        "data": item,
    })
}

func (controller Controller) ReplaceItem(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var request model.RequestItem

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": err,
		})
		return
	}

	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}

	item, err := controller.Service.UpdateItem(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"data": item,
	})
}