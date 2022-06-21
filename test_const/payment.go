//структура сущностей - добавить json теги
package payment

type Payment struct {
	Id       int     `json:"id" db:"id"`
	IdUser   int     `json:"id_user" `
	Email    string  `json:"email" `
	Sum      float32 `json:"sum" `
	Currency int     `json:"cur" `
	Create   string  `json:"create" `
	Change   string  `json:"change" `
	Status   int     `json:"status" `
}
