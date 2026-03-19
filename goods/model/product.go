package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);comment:用户名"`
	Password string `gorm:"type:char(64);comment:密码"`
	Phone    string `gorm:"type:int(11);unique;index;comment:手机号"`
	Email    string `gorm:"type:varchar(30);comment:邮箱"`
}

func (u *User) FindUserById(db *gorm.DB, id int64) error {
	return db.Where("id=?", id).First(&u).Error
}

type Product struct {
	gorm.Model
	Title       string  `gorm:"type:varchar(30);comment:商品名称"`
	Price       float64 `gorm:"type:decimal(10,2);comment:商品价格"`
	InventoryId int     `gorm:"type:int;comment:库存id"`
	TypeId      int     `gorm:"type:varchar(30);index;comment:商品分类id"`
	Img         string  `gorm:"type:varchar(255);comment:商品图片"`
}

func (p *Product) FindTitle(db *gorm.DB, title string) error {
	return db.Where("title=?", title).First(&p).Error
}

func (p *Product) ProductCreate(db *gorm.DB) error {
	return db.Create(&p).Error
}

func (p *Product) FindProductById(db *gorm.DB, id int64) error {
	return db.Where("id=?", id).First(&p).Error
}

type Order struct {
	gorm.Model
	OrderSn   string  `gorm:"type:varchar(30);comment:订单号"`
	ProductId int     `gorm:"type:int;comment:商品id"`
	Quantity  int     `gorm:"type:int;comment:购买数量"`
	Total     float64 `gorm:"type:decimal(10,2);comment:总价"`
	Status    string  `gorm:"type:varchar(30);comment:状态"`
}

func (o *Order) CreateOrder(db *gorm.DB) error {
	return db.Create(&o).Error
}

type Inventory struct {
	gorm.Model
	ProductId int    `gorm:"type:int;comment:商品id"`
	Spec      string `gorm:"type:varchar(30);comment:规格"`
	Stock     int    `gorm:"int;comment:数量"`
}

func (i *Inventory) FindProductByStock(db *gorm.DB, id int64) error {
	return db.Where("product_id=?", id).First(&i).Error
}

type ProductExtra struct {
	gorm.Model
	Color string `gorm:"type:varchar(30);comment:颜色"`
	Size  string `gorm:"type:varchar(50);comment:尺寸"`
}
