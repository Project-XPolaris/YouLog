package datasource

import (
	"time"
)

type Log interface {
	GetId() string
	GetApplication() string
	GetInstance() string
	GetLevel() int64
	GetScope() string
	GetMessage() string
	GetExtra() interface{}
	GetTime() *time.Time
}
type LogListQueryBuilder struct {
	Page        int
	PageSize    int
	LogLevels   []string
	StartTime   *time.Time
	EndTime     *time.Time
	Application string
	Orders      []string
}

func (b *LogListQueryBuilder) WithPage(page int) *LogListQueryBuilder {
	b.Page = page
	return b
}
func (b *LogListQueryBuilder) WithPageSize(pageSize int) *LogListQueryBuilder {
	b.PageSize = pageSize
	return b
}
func (b *LogListQueryBuilder) InLevels(levels []string) *LogListQueryBuilder {
	b.LogLevels = levels
	return b
}
func (b *LogListQueryBuilder) BeforeTime(beforeTime *time.Time) *LogListQueryBuilder {
	b.EndTime = beforeTime
	return b
}
func (b *LogListQueryBuilder) AfterTime(afterTime *time.Time) *LogListQueryBuilder {
	b.StartTime = afterTime
	return b
}
func (b *LogListQueryBuilder) OfApplication(application string) *LogListQueryBuilder {
	b.Application = application
	return b
}
func (b *LogListQueryBuilder) WithOrder(orders []string) *LogListQueryBuilder {
	b.Orders = orders
	return b
}

type LogDataSource interface {
	ReadLogs(queryBuilder LogListQueryBuilder) (int64, []Log, error)
	Init() error
}
