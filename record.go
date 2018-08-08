package matchlight

type RecordType uint

// Possible values for RecordType.
const (
	RecordTypeDocument   = RecordType(1)
	RecordTypeSourceCode = RecordType(2)
	RecordTypePII        = RecordType(3)
)

type RecordService interface {
	List() ([]Record, error)
	Edit(id, name, desc string) (Record, error)
	Delete(id string) error
	CreateDocument(rec CreateRecordReq) (Record, error)
	CreateSourceCode(rec CreateRecordReq) (Record, error)
	CreatePII(rec PIIRecord) (Record, error)
	// TODO: Seemingly the same as CreatePII...?
	CreateBulkPII(rec PIIRecord) (Record, error)
}

type Record struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Type         RecordType     `json:"type"`
	Ctime        uint           `json:"ctime"`
	Mtime        uint           `json:"mtime"`
	UnseenAlerts uint           `json:"number_unseen_alerts"`
	Metadata     RecordMetadata `json:"metadata"`
}

// TODO: Seems to just be an arbitrary map. Can't really make type assumptions.
type RecordMetadata map[string]interface{}

// type RecordMetadata struct {
// 	Email string `json:"email"`
// 	First string `json:"first"`
// 	Last  string `json:"last"`
// 	// TODO: What type is this?
// 	UserRecordID interface{} `json:"user_record_id"`
// }

// A PIIRecord represents a PII record in Matchlight.
type PIIRecord struct {
	Name                     string   `json:"name"`
	Description              string   `json:"desc"`
	UserRecordID             string   `json:"user_record_id"`
	BlindedEmail             string   `json:"blinded_email"`
	BlindedFirst             string   `json:"blinded_first"`
	BlindedLast              string   `json:"blinded_last"`
	EmailFingerprints        []string `json:"email_fingerprints"`
	NameFingerprints         []string `json:"name_fingerprints"`
	SSNFingerprints          []string `json:"ssn_fingerprints"`
	StreetFingerprints       []string `json:"street_address_fingerprints"`
	CityStateZipFingerprints []string `json:"city_state_zip_fingerprints"`
	PhoneFingerprints        []string `json:"phone_fingerprints"`
}

type DocumentRecord struct {
}

type CreateRecordReq struct {
	Name         string   `json:"name"`
	Description  string   `json:"desc"`
	Fingerprints []string `json:"fingerprints"`
	Metadata     RecordMetadata
}
