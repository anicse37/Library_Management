package router

import (
	"context"
	"net/http"
	"time"

	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
)

func Router(dns string, SuperAdmin models.User) {
	db := models.Database{}
	db.DB = librarySQL.StartDB(dns)
	defer db.DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()
	librarySQL.InsertSuperAdmin(ctx, db, SuperAdmin)
	router := RouterEndpoints(ctx, db)
	http.ListenAndServe(":5050", router)

}
