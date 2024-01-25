package convert

import "github.com/jinzhu/copier"

// Return a new variable copied from source
func To[DstType any, SrcType any](src SrcType) (dst DstType) {
	_ = copier.Copy(&dst, &src)
	return
}

func WithErr[DstType any, SrcType any](src SrcType, e error) (dst DstType, err error) {
	if e != nil {
		err = e
	} else {
		err = copier.Copy(&dst, &src)
	}
	return
}

func Update[Model any, Schema any](model *Model, schema Schema) error {
	return copier.Copy(model, schema)
}
