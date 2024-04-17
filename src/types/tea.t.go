package types

type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}
