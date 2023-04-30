package aggregate

import (
	"github.com/ikadgzl/ddd-in-go/entity"
	"github.com/ikadgzl/ddd-in-go/valueobject"
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}
