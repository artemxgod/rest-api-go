/* package model keeps all structures that are
represented in our database*/
package model


// struct User contains all fields of table "users"
type User struct {
	ID                int
	Email             string
	EncryptedPassowrd string
}


