package model

type MasterProvince struct {
	ID   uint
	Name string
}

func (MasterProvince) TableName() string {
	return "master_province"
}

type MasterCity struct {
	ID   uint
	Name string
}

func (MasterCity) TableName() string {
	return "master_city"
}

type MasterKelurahan struct {
	ID   uint
	Name string
}

func (MasterKelurahan) TableName() string {
	return "master_kelurahan"
}
