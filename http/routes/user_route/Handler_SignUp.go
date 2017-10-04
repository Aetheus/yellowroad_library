package user_route

import (
	"yellowroad_library/containers"
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/auth_serv/user_registration_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_response"
	"net/http"
)

type signUpForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

/*TODO try to think of another way to generate dynamic dependencies. Or at least push the responsibility
of making the Container create an instance of the dependency higher up the tree */
func SignUpHandler(container containers.Container) gin.HandlerFunc {
	return func (c *gin.Context){
		userRegistrationService := container.UserRegistrationService(nil,true);
		SignUp(c,userRegistrationService)
	}
}

func SignUp(c *gin.Context, userRegistrationService user_registration_serv.UserRegistrationService) {
	form := signUpForm{}

	if err := c.BindJSON(&form) ; err != nil {
		var err = app_error.Wrap(err).SetHttpCode(http.StatusUnprocessableEntity)
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	user, err := userRegistrationService.Run(form.Username,form.Password,form.Email)
	if (err != nil) {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	c.JSON(api_response.SuccessWithCode(
		gin.H {
			"user" : user,
		},
	))
	return
}