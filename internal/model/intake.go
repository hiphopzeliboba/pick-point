package model

import "time"

/* Сущность Intake описывает Приемку */
type Intake struct {
	ID          int
	PickPointId int
	EmployeeId  int
	Products    *[]Product
	CreatedAt   time.Time
	Status      string
}
