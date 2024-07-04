package handler

import (
	"fmt"
	"net/http"
	"strings"

	"kos-barokah-api/features/users"
	"kos-barokah-api/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.UserServiceInterface
	// sc  customer.CustomerServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service users.UserServiceInterface /* sci customer.CustomerServiceInterface,*/, jwt helper.JWTInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: service,
		/*sc:  sci,*/
		jwt: jwt,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(RegisterInput)

		if err := c.Bind(input); err != nil {
			fmt.Println("Errornya disini:", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid user input", nil))
		}

		isValid, errors := helper.ValidateForm(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation(false, "Invalid format request", errors))
		}

		if !helper.PasswordWithCombination(input.Password) {
			errors := []string{"Password must contain a combination of letters, symbols, and numbers"}
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation(false, "Invalid format request", errors))
		}

		var serviceInput = new(users.User)
		serviceInput.Name = input.Name
		serviceInput.Email = input.Email
		serviceInput.Password = input.Password
		serviceInput.Role = input.Role
		serviceInput.DateOfBirth = input.DateOfBirth
		serviceInput.PhoneNumber = input.PhoneNumber

		result, err := uh.s.Register(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(RegisterResponse)
		response.Email = result.Email
		response.PhoneNumber = result.PhoneNumber
		response.DateOfBirth = result.DateOfBirth
		response.Name = result.Name
		response.Role = result.Role
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "register success", response))
	}
}

func (uh *UserHandler) RegisterCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(RegisterInputCustomer)

		if err := c.Bind(input); err != nil {
			fmt.Println("Errornya disini:", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid user input", nil))
		}

		isValid, errors := helper.ValidateForm(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation(false, "Invalid format request", errors))
		}

		if !helper.PasswordWithCombination(input.Password) {
			errors := []string{"Password must contain a combination of letters, symbols, and numbers"}
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation(false, "Invalid format request", errors))
		}

		var serviceInput = new(users.User)
		serviceInput.Name = input.Name
		serviceInput.Email = input.Email
		serviceInput.Password = input.Password
		serviceInput.DateOfBirth = input.DateOfBirth
		serviceInput.PhoneNumber = input.PhoneNumber

		result, err := uh.s.RegisterCustomer(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		// var serviceAddCustomer = new(customer.Customer)
		// serviceAddCustomer.UserID = result.ID
		// serviceAddCustomer.FullName = input.Name
		// serviceAddCustomer.Position = input.Position
		// serviceAddCustomer.CompanyName = input.CompanyName
		// serviceAddCustomer.CompanyEmail = input.CompanyEmail
		// serviceAddCustomer.CompanyAddress = input.CompanyAddress

		// fmt.Println("coba liat userid: ", result.ID)

		// resultAddCustomer, err := uh.sc.CreateCustomer(*serviceAddCustomer)

		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		// }

		var response = new(RegisterCustomerResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Role = result.Role
		// response.CompanyName = resultAddCustomer.CompanyName
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Register customer success", response))
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid user input", nil))
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			c.Logger().Info("Handler: Login process error: ", err.Error())
			if strings.Contains(err.Error(), "Not Found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "User not found", nil))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(LoginResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success login", response))
	}
}

func (uh *UserHandler) LoginCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid user input", nil))
		}

		result, err := uh.s.LoginCustomer(input.Email, input.Password)

		if err != nil {
			c.Logger().Info("Handler: Login process error: ", err.Error())
			if strings.Contains(err.Error(), "Not Found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "Customer not found", nil))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(LoginResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success login", response))
	}
}

func (uh *UserHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := uh.jwt.CheckID(c)
		userIdInt := int(id.(float64))
		var input = new(UpdateProfile)
		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Fail to bind input", nil))
		}

		var serviceUpdate = new(users.UpdateProfile)
		serviceUpdate.Name = input.Name
		serviceUpdate.Email = input.Email
		serviceUpdate.Password = input.Password

		result, err := uh.s.UpdateProfile(userIdInt, *serviceUpdate)

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Fail to input process", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success update profile", result))
	}
}

func (uh *UserHandler) RefreshToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RefreshInput)
		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler : Bind Refresh Token Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid User Input", nil))
		}

		var currentToken = c.Get("user").(*jwt.Token)
		fmt.Println(currentToken)

		result, err := uh.jwt.RefreshJWT(input.Token, currentToken)
		if err != nil {
			c.Logger().Info("Handler : Refresh JWT Process Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Refresh JWT Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success refresh token", result))
	}
}

func (uh *UserHandler) ForgetPasswordVerify() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token = c.QueryParam("token_reset_password")
		if token == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Token not found", nil))
		}

		_, err := uh.s.TokenResetVerify(token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Token is valid", nil))
	}
}

func (uh *UserHandler) ForgetPasswordWeb() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(ForgetPasswordInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Failed to bind input", nil))
		}

		isValid, err := helper.ValidateJSON(input)
		if !isValid {
			c.Logger().Info("Handler: Bind input error: ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Validation error", err))
		}

		fmt.Println("Email: ", input.Email)

		result := uh.s.ForgetPasswordWeb(input.Email)

		if result != nil {
			fmt.Println("This is the error: ", result)
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Failed to send email", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success to send email", nil))

	}
}

func (uh *UserHandler) ResetPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token = c.QueryParam("token_reset_password")
		if token == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Token not found", nil))
		}

		dataToken, err := uh.s.TokenResetVerify(token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		var input = new(ResetPasswordInput)
		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Failed to bind input", nil))
		}

		isValid, errorMsg := helper.ValidateJSON(input)
		if !isValid {
			c.Logger().Info("Handler: Bind input error: ", err)
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Validation error", errorMsg))
		}

		if input.Password != input.PasswordConfirm {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Password not match", nil))
		}

		result := uh.s.ResetPassword(dataToken.Code, dataToken.Email, input.Password)

		if result != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Failed to reset password", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success to reset password", result))
	}
}

func (uh *UserHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		getID, err := uh.jwt.GetID(c)

		if err != nil {
			c.Logger().Error("Handler: Error getting profile by ID JWT: ", err.Error())
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Data not found", nil))
		}

		result, err := uh.s.GetProfile(int(getID))

		if err != nil {
			c.Logger().Error("Handler: Error getting profile by ID JWT: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		if result.ID == 0 {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "Data not found", nil))

		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success fetch data", result))
	}
}
