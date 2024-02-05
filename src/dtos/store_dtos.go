package dtos

type CreateStoreReq struct {
	UserID      int
	Storename   string
	Description string
}

type CreateStoreRes struct {
}

type UpdateStoreReq struct {
	StoreID     int
	Storename   string
	Description string
}

type UpdateStoreRes struct{}
