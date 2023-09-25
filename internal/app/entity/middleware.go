package entity


type InputData struct {
	FlowID string `json:"flow_id" validate:"required,max=50,alphanum"`
    UserID string `json:"user_id" validate:"required,max=50,alphanum"`
    Type  string `json:"type" validate:"required,max=20,alphanum"`
	Lang string `json:"lan"`
}