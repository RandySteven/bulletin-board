package enums

type UserRole int

const (
	Admin UserRole = iota + 1
	User
	SuperUser
)

//func WhitelistedRoleAuth() map[UserRole][]string {
//	userRoleMap := make(map[UserRole][]string)
//	userRoleMap[Admin] = []string{
//		"",
//		"",
//	}
//	userRoleMap[User] = []string{
//		"",
//		"",
//	}
//	userRoleMap[SuperUser] = []string{
//		"",
//		"",
//	}
//	return map[UserRole][]string{}
//}
