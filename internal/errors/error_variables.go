package errors_package

import "errors"

var (
	ErrorScanningUser  = errors.New("error scanning user from database")
	ErrorScanningUsers = errors.New("error scanning users from database")

	ErrorScanningadmin  = errors.New("error scanning admin from database")
	ErrorScanningadmins = errors.New("error scanning admins from database")

	ErrorWhileInserting = errors.New("errors inserting books in database")
	ErrorWhileRemoveing = errors.New("error removing book from database")

	ErrorGettingBooks = errors.New("error getting books from database")

	ErrorUnauthorized    = errors.New("unauthorized")
	ErrorInvalidUser     = errors.New("invalid user id")
	ErrorInvalidPassword = errors.New("invalid password")

	ErrorUserAlreadyExist  = errors.New("user already exists")
	ErrorCanNotRemoveBooks = errors.New("can not remove books")
	ErrorCanNotBorrowBooks = errors.New("can not borrow books")

	ErrorAdminNotAllowed = errors.New("account not approved by admin")

	ErrorReturningBooks  = errors.New("error returning book")
	ErrorAlreadyBorrowed = errors.New("error book already borrowed")

	ErrorNoBooksAvailable = errors.New("error no books available")
)
