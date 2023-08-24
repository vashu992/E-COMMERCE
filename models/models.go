package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              *string
	First_Name      *string
	Last_Name       *string
	Password        *string
	Email           *string
	Phone           *string
	Token           *string
	Refresh_Token   *string
	Created_At      time.Time
	Updated_At      time.Time
	User_ID         string
	UserCart        []ProductUser
	Address_Details []Address
	Order_Status    []Order
}

type Product struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        *string
	Rating       *uint8
	Image        *string
}

type ProductUser struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        int
	Rating       *uint
	Image        *string
}

type Address struct {
	Address_id primitive.ObjectID
	House      *string
	Street     *string
	City       *string
	Pincode    int
}

type Order struct {
	Order_id       primitive.ObjectID
	Order_Cart     []ProductUser
	Ordered_At     time.Time
	Price          int
	Discount       *int
	Payment_Method Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
