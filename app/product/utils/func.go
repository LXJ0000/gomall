package utils

import "context"

func GetUserIDFromCtx(ctx context.Context) uint32 {
	userID := ctx.Value(SessionUserID)
	if userID == nil {
		return 0
	}
	return userID.(uint32)
}
