package controller

import (
	"github.com/gofiber/fiber/v2"
	"goproj/internal/entities"
	"goproj/internal/result"
	"goproj/utils"
)

// AD 인증
// @Param authBody body entities.IAuth true " "
// @Summary AD 인증
// @Tags auth
// @Description JWT 토큰 발급 response & cookie
// @Accept json
// @Produce json
// @Success 200
// @Router /api/user/auth [post]
func Authenticate(c *fiber.Ctx) error {
	authInfo := new(entities.Auth)
	c.BodyParser(authInfo)

	if authInfo.UserId == "" || authInfo.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(result.Result[interface{}]{
			Success:      false,
			ErrorMessage: "아이디 또는 비밀번호를 입력해주세요.",
		})
	}

	status, err := testAuth(authInfo.UserId, authInfo.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(result.Result[interface{}]{
			Success:      status,
			ErrorMessage: err.Error(),
		})
	}

	if !status {
		return c.Status(fiber.StatusBadRequest).JSON(result.Result[interface{}]{
			Success:      status,
			ErrorMessage: "계정 정보가 올바르지 않습니다.",
		})
	}

	token, err := utils.GenerateNewAccessToken()

	return c.Status(fiber.StatusOK).JSON(result.Result[string]{
		Success: status,
		Result:  token,
	})
}

func testAuth(userid, password string) (bool, error) {

	return true, nil
}
