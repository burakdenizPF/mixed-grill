package entity

import sq "github.com/Masterminds/squirrel"

const (
	ComparableFilterLt ComparableFilterMode = iota
	ComparableFilterLte
	ComparableFilterGt
	ComparableFilterGte
)

func ComparableFilterPredicate[T comparable](field string, filter ComparableFilter[T]) any {
	switch filter.Mode {
	case ComparableFilterLt:
		return sq.Lt{field: filter.Value}
	case ComparableFilterLte:
		return sq.LtOrEq{field: filter.Value}
	case ComparableFilterGt:
		return sq.Gt{field: filter.Value}
	case ComparableFilterGte:
		return sq.GtOrEq{field: filter.Value}
	}

	return nil
}

type ComparableFilterMode int

type ComparableFilter[T comparable] struct {
	Mode  ComparableFilterMode
	Value T
}

type ComparableSliceFilterMode int

const (
	ComparableSliceFilterIn ComparableSliceFilterMode = iota
	ComparableSliceFilterNotIn
)

type ComparableSliceFilter[T comparable] struct {
	Mode  ComparableSliceFilterMode
	Value []T
}
