package models

import "errors"

var (
	ErrorScanningUser  = errors.New("error scanning user from database")
	ErrorScanningUsers = errors.New("error scanning users from database")

	ErrorScanningadmin  = errors.New("error scanning admin from database")
	ErrorScanningadmins = errors.New("error scanning admins from database")

	ErrorWhileInserting = errors.New("errors inserting books in database")
	ErrorWhileRemoveing = errors.New("error removing book from database")

	ErrorGettingBooks = errors.New("error getting books from database")
)
