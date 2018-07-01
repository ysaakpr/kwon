package orders

import "time"

type Interval struct {
	StartTime *time.Time
	EndTime   *time.Time
}

// Order defeneition
type Order struct {
	BaseModel

	// Status         OrderStatus `json:"status,omitempty"`
	// Parent         *Order      `json:"parent,omitempty"`
	// OrderType      OrderType   `json:"order_type,omitempty"`
	Timing          *[]Interval        `json:"timing,omitempty"`
	Buyer           *Account           `json:"buyer,omitempty"`
	Seller          *Account           `json:"seller,omitempty"`
	Items           *[]Item            `json:"items,omitempty"`
	ItemRefs        *[]ItemRef         `json:"item_refs,omitempty"`
	ShipTo          string             `json:"ship_to,omitempty"`
	BillTo          string             `json:"bill_to,omitempty"`
	TotalItemValue  float64            `json:"total_item_value,omitempty"`
	TotalOrderValue float64            `json:"total_order_value,omitempty"`
	TotalDiscount   float64            `json:"total_discount,omitempty"`
	TotalPayable    float64            `json:"total_payable,omitempty"`
	MetaData        *map[string]string `json:"meta_data,omitempty"`
}

// Status base status model
type Status struct {
	BaseModel
}

// OrderStatus defines the orders status
type OrderStatus struct {
	Status
}

// Item denotes the kind of items in the Order
type Item struct {
	BaseModel

	Ref          string             `json:"ref,omitempty"`
	RefExt       string             `json:"ref_ext,omitempty"`
	Discription  string             `json:"discription,omitempty"`
	ImageURL     string             `json:"image_url,omitempty"`
	Unit         string             `json:"unit,omitempty"`
	Quantity     int                `json:"quantity,omitempty"`
	UnitPrice    float64            `json:"unit_price,omitempty"`
	UnitDiscount float64            `json:"unit_discount,omitempty"`
	MetaData     *map[string]string `json:"meta_data,omitempty"`
}

// OrderType defines the type of the order
type OrderType struct {
	BaseModel
	Name string `json:"name,omitempty"`
}

// ItemRef is the real item that getting shiped as part of the order
type ItemRef struct {
	BaseModel

	Item     *Item      `json:"item,omitempty"`
	Quantity int        `json:"quantity,omitempty"`
	Status   ItemStatus `json:"status,omitempty"`
	LinkedBy *Link      `json:"linked_by,omitempty"`
}

// ItemStatus Status of the item
type ItemStatus struct {
	Status
}

// Link defines the link of an item or order to any other entity
type Link struct {
	Type string `json:"type,omitempty"`
}

//Account refers to buyer and seller
type Account struct {
	BaseModel

	Orders *[]Order `json:"orders,omitempty"`
}
