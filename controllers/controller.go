package controllers


import (
	"net/http"
	"github.com/gin-gonic/gin"

	model "backend_community_rest/models"
	response "backend_community_rest/responses"
	exception "backend_community_rest/exceptions"
	validation "backend_community_rest/validations"
	repository "backend_community_rest/repositories"
)


type response_data struct {
	Message string `json:"message"`
}


// GetBooks      godoc
// @Summary      Join Community Request
// @Description  Join Community Request
// @Tags         join
// @Produce      json
// @Success      200  {object}  models.JoinRequest
// @Router       /join [post]
func AddJoinRequest(c *gin.Context) {

	// Declare Payload Struct
	var user_request model.JoinRequest

	// Default Value
	http_response, status, message := response.AddJoinResponse(
		http.StatusOK,
	)

	// Bind JSON Payload
	err := c.BindJSON(&user_request)
	if err != nil {
		http_response, status, message = response.AddJoinResponse(
			http.StatusBadRequest,
		)
	}

	exception.TryCatchError(err)

	// Github Username Validation
	valid, err := validation.ValidateGitHubUsername(user_request.UsernameGithub)
	if err != nil {
		http_response, status, message = response.AddJoinResponse(
			http.StatusInternalServerError,
		)
	}

	exception.TryCatchError(err)

	if !valid {
		http_response, status, message = response.GithubNotFoundResponse(
			http.StatusBadRequest,
		)

	} else {
		// Call Repository
		if !repository.AddJoinRequestRepo(user_request) {
			http_response, status, message = response.AddJoinResponse(
				http.StatusUnauthorized,
			)
		}
	}

	// Response JSON Struct
	response := model.JoinResponse {
		HttpResponse: http_response,
		Status: status,
		Data: response_data{

			Message: message,
		},
	}

	c.JSON(http_response, response)
}