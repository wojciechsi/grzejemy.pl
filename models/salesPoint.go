package models

// Region @todo enumeration or sth
type Region string

type SalesPoint struct {
	Region
	Vendor
}

func NewSalesPoint(region Region, vendor Vendor) SalesPoint {
	return SalesPoint{
		region,
		vendor,
	}
}

func (s SalesPoint) GetRegion() string {
	return string(s.Region)
}
