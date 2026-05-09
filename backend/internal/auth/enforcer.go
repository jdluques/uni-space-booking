package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/jdluques/uni-space-booking/internal/associations/user_organization"
)

func InitCasbin(userOrganizationService *user_organization.UserOrganizationService) (*casbin.Enforcer, error) {
	adapter := fileadapter.NewAdapter("../../config/casbin/policies.csv")

	enforcer, err := casbin.NewEnforcer("../../config/casbin/model.conf", adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Casbin: %w", err)
	}

	enforcer.AddFunction("isOrgMember", isOrgMember(userOrganizationService))

	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("failed to load Casbin policy: %w", err)
	}

	return enforcer, nil
}

func isOrgMember(service *user_organization.UserOrganizationService) func(args ...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		userID, ok1 := args[0].(string)
		orgID, ok2 := args[1].(string)
		if !ok1 || !ok2 {
			return false, fmt.Errorf("invalid arguments passed to isOrgMember")
		}

		return service.UserIsMemberOfOrganization(userID, orgID), nil
	}
}
