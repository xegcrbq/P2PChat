package api

//func (svc *APIService) AuthMiddleware() fiber.Handler {
//	return func(ctx *fiber.Ctx) error {
//		sessionId := ctx.Cookies("session_id")
//		if sessionId == "" {
//			err := ctx.SendStatus(http.StatusUnauthorized)
//			if err != nil {
//				return err
//			}
//			return nil
//		}
//		_, token, err := cC.tknz.ParseDataClaims(sessionId)
//		if err != nil || !token.Valid {
//			err := ctx.SendStatus(http.StatusUnauthorized)
//			if err != nil {
//				return err
//			}
//			return nil
//		}
//		return next(ctx)
//	}
//}
//
//func (svc *APIService) XRequestIDMiddleware() fiber.Handler {
//	return func(ctx *fiber.Ctx) error {
//		xRequestID := ctx.Request().Header.Get(constants.HeaderKeyRequestID)
//		if len(xRequestID) == 0 {
//			xRequestID, err := core.GenUUID()
//			if err != nil {
//				return err
//			}
//			ctx.Request().Header.Set(constants.HeaderKeyRequestID, xRequestID)
//		}
//		return next(ctx)
//	}
//}
