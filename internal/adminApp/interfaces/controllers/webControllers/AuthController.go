package webcontrollers

import (
	"fmt"

	"github.com/alwismt/selectify/internal/adminApp/domains/Auth/services"
	transferobjects "github.com/alwismt/selectify/internal/adminApp/interfaces/transferObjects"
	"github.com/alwismt/selectify/internal/infrastructure/session"
	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	AuthCheck(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type authController struct {
	authService services.AuthService
	session     session.Session
}

func NewAuthController(authService services.AuthService, store session.Session) AuthController {
	if store == nil {
		store = session.NewSession(nil)
	}
	return &authController{authService: authService, session: store}

}

func (cc *authController) Login(c *fiber.Ctx) error {
	// set seesion keys and values
	// cc.session.NewSession(c)
	redirect := c.Query("redirect") // get redirect url

	return c.Render("auth/login", fiber.Map{
		"title":    "Login | Selectify Admin",
		"email":    c.Query("email"),
		"error":    c.Query("error"),
		"redirect": redirect,
	}, "layouts/sign")
}

func (cc *authController) AuthCheck(c *fiber.Ctx) error {
	data := new(transferobjects.SignInDTO)
	// Checking received data from JSON body.
	if err := c.BodyParser(data); err != nil {
		return c.Redirect("/controlpanel/login?error=requiredfields")
	}
	// Create a new validator and validate the request.
	validate := utils.NewValidator()
	if err := validate.Struct(data); err != nil {
		return c.Redirect("/controlpanel/login?error=requiredfields")
	}
	ctx, cancel := utils.CreateContextWithTimeout()
	defer cancel()

	eventData := new(transferobjects.EventAuthDTO)
	eventData.IP = c.IP()
	eventData.UserAgent = c.Get("User-Agent")
	eventData.SessionID = cc.session.ID(c)

	redirect := data.Redirect
	id, name, user_type, err := cc.authService.AuthCheck(ctx, data, eventData)

	if err != nil {
		return c.Redirect(fmt.Sprintf("/controlpanel/login?error=%v&email=%v&redirect=%v", err, data.Email, redirect))
	}

	// Save user details to session
	cc.session.Set(c, "id", id.String())
	cc.session.Set(c, "type", user_type)
	cc.session.Set(c, "name", name)
	cc.session.Set(c, "email", data.Email)

	if redirect != "" {
		return c.Redirect(fmt.Sprintf("/%s", redirect))
	}

	return c.Redirect("/controlpanel/dashboard")
}

func (cc *authController) Logout(c *fiber.Ctx) error {
	cc.session.Destroy(c)
	return c.Redirect("/controlpanel/login")
}
