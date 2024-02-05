package model

type Shipment struct{}

func (Shipment) TableName() string {
	return "shipment"
}
