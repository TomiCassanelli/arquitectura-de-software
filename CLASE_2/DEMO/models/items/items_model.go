package items

type ItemModel struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Title string  `bson:"title" json:"title"`
	Price float64 `bson:"price" json:"price"`
}
