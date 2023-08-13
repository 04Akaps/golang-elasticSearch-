package types

type CollectionOne struct {
	Name string `bson:"name"`
	Age  int64  `bson:"age"`
}

type CollectionTwo struct {
	Address string `bson:"address"`
	Price   int64  `bson:"price"`
	Inner   Inner  `bson:"inner"`
	Array   Array  `bson:"array"`
}

type Inner struct {
	Owner string `bson:"owner"`
}

type Array struct {
	Array []string `bson:"array"`
}
