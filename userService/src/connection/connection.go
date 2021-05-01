package connection

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func InitializeConnection() *pg.DB {
	opt, err := pg.ParseURL("postgres://ghostchaser:ghostchaser@test-db.cjwtduue94lp.eu-central-1.rds.amazonaws.com:5432/user_service")
	db := pg.Connect(opt)
	defer db.Close()

	// err = createSchema(db)
	// if err != nil {
	// 	panic(err)
	// }

	user1 := &User{
		Name:   "kaxelo2",
		Emails: []string{"kaxelo1@admin", "kaxelo2@admin"},
	}
	_, err = db.Model(user1).Insert()
	if err != nil {

	}

	// _, err = db.Model(&User{
	// 	Name:   "jojo",
	// 	Emails: []string{"root1@root", "root2@root"},
	// }).Insert()
	// if err != nil {
	// 	panic(err)
	// }

	// story1 := &Story{
	// 	Title:    "Cool story",
	// 	AuthorId: user1.Id,
	// }
	// _, err = db.Model(story1).Insert()
	// if err != nil {
	// 	panic(err)
	// }

	// Select user by primary key.
	// user := &User{Id: user1.Id}
	// err = db.Model(user).WherePK().Select()
	// if err != nil {
	// 	panic(err)
	// }

	// Select all users.
	// var users []User
	// err = db.Model(&users).Select()
	// if err != nil {
	// 	panic(err)
	// }

	// Select story and associated author in one query.
	// story := new(Story)
	// err = db.Model(story).
	// 	Relation("Author").
	// 	Where("story.id = ?", story1.Id).
	// 	Select()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)
	// fmt.Println(users)
	// fmt.Println(story)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>

	return db
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Story)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
