package types

const (
	ModuleName   = "leave"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	StudentKey = []byte{0x01}
	LeaveKey   = []byte{0x02}
	AdminKey   = []byte{0x03}
)

func AdminStoreKey(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

func StudentStoreKey(student string) []byte {
	key := make([]byte, len(StudentKey)+len(student))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], []byte(student))
	return key
}

func LeaveStoreKey(leaveCounter string) []byte {
	key := make([]byte, len(LeaveKey)+len(leaveCounter))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(leaveCounter))
	return key
}
