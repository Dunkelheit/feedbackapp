package database

import (
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/jinzhu/gorm"
)

// DB is the database instance
var DB *gorm.DB

// OpenDB opens the database
func OpenDB() {
	var err error
	DB, err = gorm.Open("postgres", "sslmode=disable dbname=feedbackapp host=localhost user=arturo.martinez")
	if err != nil {
		panic("Failed to connect database")
	}
	// defer DB.Close()

	DB.LogMode(true)

	DB.DropTable(&model.Review{}, &model.Card{}, &model.User{})
	DB.AutoMigrate(&model.Card{}, &model.User{}, &model.Review{})

	// DB.Model(&model.Review{}).Related(&model.User{}, "Reviewer")
	// DB.Model(&model.Review{}).Related(&model.User{}, "Reviewee")
	// DB.Model(&model.Review{}).Related(&model.User{}, "reviewer_id")
	// DB.Model(&model.Review{}).Related(&model.User{}, "reviewee_id")

	DB.Model(&model.Review{}).AddForeignKey("reviewer_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Review{}).AddForeignKey("reviewee_id", "users(id)", "CASCADE", "CASCADE")

	DB.Create(&model.Card{Title: "Great moves", Category: model.CardCategoryPositive})
	DB.Create(&model.Card{Title: "Computer hacking skills", Category: model.CardCategoryPositive})
	DB.Create(&model.Card{Title: "Lazy eye", Category: model.CardCategoryNegative})
	DB.Create(&model.Card{Title: "Stinky feet", Category: model.CardCategoryNegative})

	DB.Create(&model.User{Name: "Arturo Martinez", Email: "arturo@icemobile.com", JobTitle: "Node.js Developer", Avatar: "none"})
	DB.Create(&model.User{Name: "Rosa van Colmjon", Email: "rosa@icemobile.com", JobTitle: "Scrum Master Extraordinaire", Avatar: "none"})
	DB.Create(&model.User{Name: "Caio Borges", Email: "caio@icemobile.com", JobTitle: "Test Engineer", Avatar: "none"})
	DB.Create(&model.User{Name: "Marco Silva", Email: "marco@icemobile.com", JobTitle: "Visual Designer", Avatar: "none"})

	DB.Create(&model.Review{
		UUID:      "1234",
		Remark:    "Lorem ipsum",
		Completed: false,
		Reviewee:  model.User{Name: "Willem Verker", Email: "willem@icemobile.com", JobTitle: "UX Designer", Avatar: "none"},
		Reviewer:  model.User{Name: "Miika Kossi", Email: "miika@icemobile.com", JobTitle: "Associate Creative Director", Avatar: "none"},
		Cards:     []model.Card{{Title: "Bad english", Category: model.CardCategoryNegative}},
	})
}
