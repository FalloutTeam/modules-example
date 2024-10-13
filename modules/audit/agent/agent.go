package agent

import authContract "example/modules/auth/contract"

type AuditAgent interface {
	GetUser(int64) string
}

type auditAgent struct {
	authContract authContract.AuthContract
}

func NewAuditAgent(authContract authContract.AuthContract) *auditAgent {
	return &auditAgent{authContract: authContract}
}

func (a *auditAgent) GetUser(id int64) string {
	return a.authContract.GetUser(id)
}
