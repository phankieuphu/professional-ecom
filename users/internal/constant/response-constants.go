package user_constant

// Response message success
const (
	MessageUserCreatedSuccess            = "User create successfully"
	MessageUserUpdatedSuccess            = "User updated successfully"
	MessageUserDeletedSuccess            = "User deleted successfully"
	MessageUserFetchedSuccess            = "User fetched successfully"
	MessageUserListFetchedSuccess        = "User list fetched successfully"
	MessageUserLoginSuccess              = "User login successfully"
	MessageUserLogoutSuccess             = "User logout successfully"
	MessageUserPasswordChangedSuccess    = "User password changed successfully"
	MessageUserProfileUpdatedSuccess     = "User profile updated successfully"
	MessageUserEmailVerifiedSuccess      = "User email verified successfully"
	MessageUserRoleUpdatedSuccess        = "User role updated successfully"
	MessageUserPermissionsUpdatedSuccess = "User permissions updated successfully"
	MessageUserAccountDeactivatedSuccess = "User account deactivated successfully"
	MessageUserAccountReactivatedSuccess = "User account reactivated successfully"
	MessageUserAccountLockedSuccess      = "User account locked successfully"
	MessageUserAccountUnlockedSuccess    = "User account unlocked successfully"
	MessageUserAccountSuspendedSuccess   = "User account suspended successfully"
	MessageUserAccountRestoredSuccess    = "User account restored successfully"
	MessageUserAccountDeletedSuccess     = "User account deleted successfully"
)

// Response message error
const (
	MessageValidateError           = "Validation error"
	MessageUserNotFoundError       = "User not found"
	MessageUserAlreadyExistsError  = "User already exists"
	MessageMissingInformationError = "Missing required information"
	MessageInvalidCredentialsError = "Invalid credentials"
	MessageUnauthorizedError       = "Unauthorized access"
	MessageForbiddenError          = "Forbidden access"
	MessageInternalServerError     = "Internal server error"
	MessageTimeoutError            = "Request timeout"

	MessageUserCreatedError         = "Failed to create user"
	MessageUserUpdatedError         = "Failed to update user"
	MessageUserDeletedError         = "Failed to delete user"
	MessageUserFetchedError         = "Failed to fetch user"
	MessageUserListFetchedError     = "Failed to fetch user list"
	MessageUserLoginError           = "Failed to login user"
	MessageUserLogoutError          = "Failed to logout user"
	MessageUserPasswordChangedError = "Failed to change user password"
	MessageUserProfileUpdatedError  = "Failed to update user profile"
	MessageUserEmailVerifiedError   = "Failed to verify user email"
	MessageUserRoleUpdatedError     = "Failed to update user role"
)

// Internal code
const (
	// Success codes
	CodeUserCreatedSuccess            = 1000
	CodeUserUpdatedSuccess            = 1001
	CodeUserDeletedSuccess            = 1002
	CodeUserFetchedSuccess            = 1003
	CodeUserListFetchedSuccess        = 1004
	CodeUserLoginSuccess              = 1005
	CodeUserLogoutSuccess             = 1006
	CodeUserPasswordChangedSuccess    = 1007
	CodeUserProfileUpdatedSuccess     = 1008
	CodeUserEmailVerifiedSuccess      = 1009
	CodeUserRoleUpdatedSuccess        = 1010
	CodeUserPermissionsUpdatedSuccess = 1011
	CodeUserAccountDeactivatedSuccess = 1012
	CodeUserAccountReactivatedSuccess = 1013
	CodeUserAccountLockedSuccess      = 1014
	CodeUserAccountUnlockedSuccess    = 1015
	CodeUserAccountSuspendedSuccess   = 1016
	CodeUserAccountRestoredSuccess    = 1017
	CodeUserAccountDeletedSuccess     = 1018

	// Error codes
	CodeUserCreatedError            = 2000
	CodeUserUpdatedError            = 2001
	CodeUserDeletedError            = 2002
	CodeUserFetchedError            = 2003
	CodeUserListFetchedError        = 2004
	CodeUserLoginError              = 2005
	CodeUserLogoutError             = 2006
	CodeUserPasswordChangedError    = 2007
	CodeUserProfileUpdatedError     = 2008
	CodeUserEmailVerifiedError      = 2009
	CodeUserRoleUpdatedError        = 2010
	CodeUserPermissionsUpdatedError = 2011
	CodeUserNotFoundError           = 2012
	CodeUserAlreadyExistsError      = 2013
	CodeMissingInformationError     = 2014
	CodeInvalidCredentialsError     = 2015
	CodeUnauthorizedError           = 2016
	CodeForbiddenError              = 2017
	CodeInternalServerError         = 2018
	CodeTimeoutError                = 2019
	CodeValidateError               = 2020
)
