package datastore

import "io"

type Datastore interface {
	io.Closer

	Property() PropertyRepo
}
