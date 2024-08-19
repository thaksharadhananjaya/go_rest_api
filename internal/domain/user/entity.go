package user

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"size:255;not null"`
    Email string `gorm:"size:255;unique;not null"`
    Age   int    `gorm:"not null"`
}
