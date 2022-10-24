package controller

import (
	"goproj/internal/db"
	"goproj/internal/entities"
	"goproj/internal/result"
	"goproj/utils"

	"github.com/gofiber/fiber/v2"
)

// 사용자 1명 조회
// @Summary 사용자 1명 조회
// @Tags users
// @Description
// @Success 200
// @Param userId path string false "사용자 아이디"
// @Router /api/user [get]
func GetUser(c *fiber.Ctx) error {
	query := utils.ConvertUrlParamToBson(c)
	if len(query) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(result.Result[interface{}]{
			ErrorMessage: "사용자 아이디를 입력하세요.",
			Success:      false,
		})
	}

	dbResult, err := db.FindOne[entities.User]("user", query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(result.Result[interface{}]{
			ErrorMessage: "사용자 정보가 없습니다.",
			Success:      false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result.Result[entities.User]{
		Success: true,
		Result:  dbResult,
	})
}

// 전체 사용자 조회
// @Summary 전체 사용자 조회
// @Tags users
// @Description JWT 토큰 필수
// @Success 200
// @Security ApiKeyAuth
// @Param userId path string false "사용자 아이디"
// @Param name path string false "사용자 이름"
// @Param limit path string false "페이징(Limit)"
// @Param skip path string false "페이징(Skip)"
// @Router /api/admin/users [get]
func GetUsers(c *fiber.Ctx) error {
	query := utils.ConvertUrlParamToBson(c)
	option := utils.ConvertUrlParamToBsonOption(c)

	dbResult, err := db.FindAll[entities.User]("user", query, option)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(result.Result[interface{}]{
			ErrorMessage: err.Error(),
			Success:      false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result.Result[*result.PageResult[[]entities.User]]{
		Success: true,
		Result:  dbResult,
	})
}
