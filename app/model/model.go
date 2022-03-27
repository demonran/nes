package model

import (
	"gorm.io/gorm"
	"time"
)

type Entity interface {
	SetId(id string)
	GetId() string
}

type BModel struct {
	ID        string `gorm:"primarykey;size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *BModel) SetId(id string) {
	m.ID = id
}

func (m *BModel) GetId() string {
	return m.ID
}

type Customer struct {
	BModel
	Company string `json:"company"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}

type GeneralInfo struct {
	BModel
	LiveCount            int      `json:"liveCount"`
	MaxClientCount       int      `json:"maxClientCount"`
	OperationExperience  bool     `json:"operationExperience"`
	TypesOfViewer        ViewType `json:"typesOfViewer"`
	VendorSupport        bool     `json:"vendorSupport"`
	OutfieldScene        string   `json:"outfieldScene"`
	VendorSupportModules string   `json:"vendorSupportModules"`
	LiveDevices          string   `json:"liveDevices"`
	ObservationStyles    string   `json:"observationStyles"`
	EvaluationId         string   `json:"evaluation_id"`
}

type ViewType string

const (
	Internal ViewType = "内部员工"
	External ViewType = "外部用户"
)

type SiteInfo struct {
	BModel
	Location                 string        `json:"location"`
	LiveCount                int           `json:"liveCount"`
	LoadBalanceSupported     bool          `json:"loadBalanceSupported"`
	WirelessClientProportion float32       `json:"wirelessClientProportion"`
	JoinUpCountPerAP         int           `json:"joinUpCountPerAP"`
	WirelessIssue            string        `json:"wirelessIssue"`
	Bandwidth                int           `json:"bandwidth"`
	EvaluationId             string        `json:"evaluationId"`
	CircuitInfos             []CircuitInfo `json:"circuitInfos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CircuitInfo struct {
	BModel
	Bandwidth  int         `json:"bandwidth"`
	Operator   Operator    `json:"operator"`
	Type       CircuitType `json:"type"`
	SiteInfoId string      `json:"site_info_id"`
}

type Operator string

const (
	ChinaTelecom Operator = "电信"
	ChinaMobile  Operator = "移动"
	ChinaUnicom  Operator = "联通"
	Other        Operator = ""
)

type CircuitType string

const (
	Internet CircuitType = "互联网专线"
	ADSL     CircuitType = "ADSL"
	MPLS     CircuitType = "私网专线"
	P2P      CircuitType = "裸光"
)

type Evaluation struct {
	BModel
	GeneralInfo *GeneralInfo `json:"generalInfo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SiteInfos   []SiteInfo   `json:"siteInfos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
