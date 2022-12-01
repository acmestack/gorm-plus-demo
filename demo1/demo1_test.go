package demo1

import (
	"fmt"
	"github.com/zouchangfu/gorm-plus/gormplus"
	"github.com/zouchangfu/gorm-plus/mapper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func init() {
	dsn := "root:root-abcd-1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	gormplus.Init(gormDb)
}

type WarnVulnerability struct {
	gorm.Model
	CveId            string
	CnvdId           string
	CnnvdId          string
	AvdId            string
	ThreatenLevel    int
	Score            string
	Name             string
	Utilization      string
	Patch            string
	PublishedDate    string
	Remark           string
	Proposal         string
	AttackPath       string
	AttackComplexity string
	Permission       string
	Influence        string
	ExpMaturity      string
	Confidentiality  string
	Integrity        string
	ServerHarm       string
	Type             string
}

var SelectUser = `
   select * from user #{}
`

func (WarnVulnerability) TableName() string {
	return "warn_vulnerability"
}

func TestSelect(t *testing.T) {
	q := mapper.Query[WarnVulnerability]{}
	q.Lt("threaten_level", 3).Like("name", "Apache")
	_, results := mapper.SelectCount(&q)
	fmt.Println(results)
}

func TestSelectCount(t *testing.T) {
	q := mapper.Query[WarnVulnerability]{}
	_, results := mapper.SelectCount(&q)
	fmt.Println(results)
}
