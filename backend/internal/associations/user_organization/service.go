package user_organization

type UserOrganizationService struct {
	userOrganizationRepository UserOrganizationRepository
}

func NewUserOrganizationService(userOrganizationRepository UserOrganizationRepository) *UserOrganizationService {
	return &UserOrganizationService{
		userOrganizationRepository: userOrganizationRepository,
	}
}

func (s *UserOrganizationService) UserIsMemberOfOrganization(userID string, orgID string) bool {
	return s.userOrganizationRepository.IsUserInOrganization(userID, orgID)
}
