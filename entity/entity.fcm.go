package entity

type UpdateFcmRequest struct {
	UserId   string `json:"user_id" validate:"required"`
	FcmToken string `json:"fcm_token" validate:"required"`
}
