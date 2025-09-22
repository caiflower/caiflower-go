package gomock

//go:generate mockgen -destination ./mock_people.go --package gomock -source interface.go
type People interface {
	DescribeName() string
	DescribeAge() string
	SaySame(string) string
}
