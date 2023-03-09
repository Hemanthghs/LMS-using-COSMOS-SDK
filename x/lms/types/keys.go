package types

const (
	ModuleName   = "leave"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	StudentKey     = []byte{0x01}
	LeaveKey       = []byte{0x02}
	AdminKey       = []byte{0x03}
	AcceptLeaveKey = []byte{0x04}
	CounterKey     = []byte{0x05}
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

func LeaveStoreKey(sid string, studentLeavesCount string) []byte {
	key := make([]byte, len(LeaveKey)+len(sid)+len(studentLeavesCount))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(sid))
	copy(key[len(LeaveKey)+len(sid):], []byte(studentLeavesCount))
	return key
}

func StudentLeavesCounterKey(sid string) []byte {
	key := make([]byte, len(CounterKey)+len(sid))
	copy(key, CounterKey)
	copy(key[len(CounterKey):], []byte(sid))
	return key
}

func AcceptLeaveStoreKey(leaveid string) []byte {
	key := make([]byte, len(AcceptLeaveKey)+len(leaveid))
	copy(key, leaveid)
	copy(key[len(AcceptLeaveKey):], []byte(leaveid))
	return key
}
