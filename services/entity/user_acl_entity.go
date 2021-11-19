package entity

import (
	"sort"
)

// UserAccessControlLevel DataStructure
type UserACL struct {
	Permissions []string
	Roles       []string
}

//NewUserACl returns roles and permission, each being sorted
func NewUserACL(roles []*Role, permissions []*Permission) (*UserACL, error) {

	perms := []string{}
	rls := []string{}

	for _, r := range roles {
		rls = append(rls, r.Name)
	}

	for _, p := range permissions {
		perms = append(perms, p.Path)
	}

	sort.Strings(rls)   //sort the roles
	sort.Strings(perms) //sort permission

	userACL := &UserACL{
		Permissions: perms,
		Roles:       rls,
	}
	return userACL, nil
}

//HasPermission function check if the user has the given permission
//returns true if has the given permission
func (uc *UserACL) HasPermission(perm string) bool {
	for _, v := range uc.Permissions {
		if v == perm {
			return true
		}
	}
	return false
	//res := sort.SearchStrings(uc.Permissions, perm)
	//fmt.Printf("\n\n\nPermission:index = %d, length = %d\n\n", res, len(uc.Permissions))
	//return res < len(uc.Permissions)
}

//HasRole function check if user has the given role
//returns true if the user has the given role
func (uc *UserACL) HasRole(role string) bool {
	for _, v := range uc.Permissions {
		if v == role {
			return true
		}
	}
	return false
	/*res := sort.SearchStrings(uc.Roles, role)
	fmt.Printf("\n\n\n Roles:index = %d, length = %d\n\n", res, len(uc.Permissions))

	return res < len(uc.Roles)*/
}
