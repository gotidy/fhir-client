package models

func NewString(v string) *string {
	return &v
}

func ToString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func NewInt(v int) *int {
	return &v
}

func ToInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func NewBool(v bool) *bool {
	return &v
}

func ToBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

type Decimal = float64

func NewDecimal(v float64) *Decimal {
	return &v
}

func ToDecimal(v *Decimal) Decimal {
	if v == nil {
		return 0
	}
	return *v
}

func ToFloat64(v *Decimal) float64 {
	if v == nil {
		return 0
	}
	return *v
}
