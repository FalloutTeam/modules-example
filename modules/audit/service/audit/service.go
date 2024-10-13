package audit

import (
	"example/modules/audit/agent"
	"example/modules/audit/repository"
	"fmt"
)

type AuditService interface {
	Log(msg string) error
}

type auditService struct {
	repo       repository.AuditRepository
	auditAgent agent.AuditAgent
}

func NewAuditService(repo repository.AuditRepository, auditAgent agent.AuditAgent) AuditService {
	return &auditService{
		repo:       repo,
		auditAgent: auditAgent,
	}
}

func (s *auditService) Log(msg string) error {
	fmt.Println(msg)

	userString := s.auditAgent.GetUser(1)
	fmt.Println(userString)

	return nil
}
