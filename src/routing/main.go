package routing

import (
    "fmt"
    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "src/databases"
)

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Age      int `json:"age"`
}

func (u User) String() string {
    return fmt.Sprintf("Name:%s \n Age:%d \n ",
        u.Name,
        u.Age)
}

// ユーザー情報取得
func BaseAPI_GET() echo.HandlerFunc {
    return func(c echo.Context) error {
        db := databases.GormConnect()
        defer db.Close()

        text1 := search(db)
        return c.JSON(200, text1)
    }
}

// ユーザーを登録，更新
func BaseAPI_POST() echo.HandlerFunc {
    return func(c echo.Context) error {
        db := databases.GormConnect()
        defer db.Close()

        //追加・更新
        user := new(User)
        if err := c.Bind(user); err != nil {
            return err
        }

        user1 := User{
            Name:     user.Name,
            Age:      user.Age,
        }

        insertUsers := []User{user1}
        insert(insertUsers, db)

        // update(user1, db)

        return c.JSON(200, "追加完了")
    }
}

func insert(users []User, db *gorm.DB) {
    for _, user := range users {
        db.NewRecord(user)
        db.Create(&user)
    }
}

func update(users User, db *gorm.DB) {
    var user User
    db.Model(&user).Where("id = ?", 1).Update(map[string]interface{}{"name": users.Name, "age": users.Age})
}

func search(db *gorm.DB) []User {
    var user []User
    // 同じ位置にいるリストを取得
    db.Raw("SELECT * FROM users").Scan(&user)
    return user
}
