package payload

type Status struct {
	Exist *bool `json:"exist"`
}

type CertificateGetRequest struct {
	Type *string `json:"type" form:"type" validate:"required,oneof=p12 mobileconfig sswan"`
}
type CertificateGetResponse struct {
	Certificate *string `json:"certificate"`
}
type CertificateResult struct {
	Result *string `json:"result"`
}
