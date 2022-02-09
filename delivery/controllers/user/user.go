package user

import (
	"net/http"

	"part3/delivery/middlewares"
	"part3/lib/database/user"
	"part3/models/base"
	"part3/models/user/request"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := request.UserRegister{}

		if err := c.Bind(&newUser); err != nil || newUser.Email == "" || newUser.Password == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(nil, "error in request Create", nil))
		}

		res, err := uc.repo.Create(newUser.ToUser())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(nil, "error in access Create", nil))
		}

		return c.JSON(http.StatusCreated, base.Success(http.StatusCreated, "Success Create", res))
	}
}

func (uc *UserController) GetById() echo.HandlerFunc {

	return func(c echo.Context) error {
		userid := int(middlewares.ExtractTokenId(c))

		res, err := uc.repo.GetById(userid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(nil, "error in access Get By id", nil))
		}

		return c.JSON(http.StatusOK, base.Success(http.StatusOK, "Success Get By Id", res))
	}
}

func (uc *UserController) UpdateById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid := int(middlewares.ExtractTokenId(c))
		upUser := request.UserRegister{}

		if err := c.Bind(&upUser); err != nil || upUser.Name == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(http.StatusBadRequest, "error in request Update", nil))
		}

		res, err := uc.repo.UpdateById(userid, upUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(http.StatusInternalServerError, "error in access Update", nil))
		}

		return c.JSON(http.StatusOK, base.Success(http.StatusOK, "Success Update By Id", res))
	}
}

func (uc *UserController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {

		userid := int(middlewares.ExtractTokenId(c))
		upUser := request.UserRegister{}

		if err := c.Bind(&upUser); err != nil || upUser.Name == "" || upUser.Password == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(http.StatusBadRequest, "error in request Delete", nil))
		}

		res, err := uc.repo.DeleteById(userid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(http.StatusInternalServerError, "error in access Delete", nil))
		}

		return c.JSON(http.StatusOK, base.Success(http.StatusOK, "Success Delete By Id", res))
	}
}

func (uc *UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userid := int(middlewares.ExtractTokenId(c))
		admLog := request.Userlogin{}
		res, err := uc.repo.GetAll()

		if err != nil || admLog.Email != "admin" && admLog.Password != "admin" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(nil, "error in request Get", nil))
		}

		return c.JSON(http.StatusOK, base.Success(http.StatusOK, "Success Get All User", res))

	}
}
