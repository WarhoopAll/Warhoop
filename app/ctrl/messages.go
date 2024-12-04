package ctrl

import "github.com/gofiber/fiber/v2"

const (
	MsgNoRows                 = "No rows found in the database."
	MsgNoData                 = "No data available."
	MsgHostResponse           = "Received an error from the host."
	MsgIncorrectPassword      = "The password must be at least 6 characters long."
	MsgIncorrectEmail         = "The email address is invalid."
	MsgIOAction               = "An I/O error occurred."
	MsgAccountExists          = "An account with this information already exists."
	MsgAccountAlreadyVerified = "The account has already been verified."
	MsgAccountNotFound        = "The account could not be found."
	MsgInternal               = "An internal error occurred. Please try again later."
	MsgDataBase               = "A database error occurred."
	MsgNotFound               = "The requested resource was not found."
	MsgNotAllowed             = "This action is not allowed."
	MsgForbidden              = "Access is forbidden."
	MsgBadConfig              = "The configuration is invalid."
	MsgBadToken               = "The token is invalid."
	MsgBadSession             = "The session is invalid."
	MsgInvalidation           = "Validation failed."
	MsgUnauthorized           = "You are not authorized to perform this action."
	MsgSendEmail              = "An error occurred while sending the email."
	MsgBody                   = "There was an error with the input body."
	MsgValidate               = "Input validation failed."
	MsgSignIn                 = "authorization success"
	MsgSignUp                 = "registration success"
	MsgLogout                 = "logout complete"
	MsgAvatarUpdate           = "avatar update"
	MsgSuccess                = "success"
)

var ErrorMapping = map[string]int{
	MsgNoRows:                 fiber.StatusNotFound,
	MsgNoData:                 fiber.StatusNotFound,
	MsgHostResponse:           fiber.StatusInternalServerError,
	MsgIncorrectPassword:      fiber.StatusUnauthorized,
	MsgIncorrectEmail:         fiber.StatusBadRequest,
	MsgIOAction:               fiber.StatusInternalServerError,
	MsgAccountExists:          fiber.StatusConflict,
	MsgAccountAlreadyVerified: fiber.StatusConflict,
	MsgAccountNotFound:        fiber.StatusNotFound,
	MsgInternal:               fiber.StatusInternalServerError,
	MsgDataBase:               fiber.StatusInternalServerError,
	MsgNotFound:               fiber.StatusNotFound,
	MsgNotAllowed:             fiber.StatusMethodNotAllowed,
	MsgForbidden:              fiber.StatusForbidden,
	MsgBadConfig:              fiber.StatusBadRequest,
	MsgBadToken:               fiber.StatusForbidden,
	MsgBadSession:             fiber.StatusUnauthorized,
	MsgInvalidation:           fiber.StatusBadRequest,
	MsgUnauthorized:           fiber.StatusUnauthorized,
	MsgSendEmail:              fiber.StatusInternalServerError,
	MsgBody:                   fiber.StatusBadRequest,
	MsgValidate:               fiber.StatusBadRequest,
	MsgSignIn:                 fiber.StatusOK,
	MsgAvatarUpdate:           fiber.StatusOK,
	MsgSignUp:                 fiber.StatusCreated,
	MsgLogout:                 fiber.StatusOK,
	MsgSuccess:                fiber.StatusOK,
}
