package types

const (
	ModuleName   = "leave"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	StudentId = []byte{0x01}
	LeaveId   = []byte{0x02}
	AdminId   = []byte{0x03}
)

func AdminStoreKey(admin string) []byte {
	key := make([]byte, len(AdminId)+len(admin))
	copy(key, AdminId)
	copy(key[len(AdminId):], []byte(admin))
	return key
}

func StudentStoreKey(student string) []byte {
	key := make([]byte, len(StudentId)+len(student))
	copy(key, StudentId)
	copy(key[len(StudentId):], []byte(student))
	return key
}
