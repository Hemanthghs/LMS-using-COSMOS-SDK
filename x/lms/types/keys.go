package types

import "strconv"

const (
	ModuleName   = "leave"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	StudentKey      = []byte{0x01}
	LeaveKey        = []byte{0x02}
	AdminKey        = []byte{0x03}
	AcceptLeaveKey  = []byte{0x04}
	LeaveCounterKey = []byte{0x05}
	LeaveStatusKey  = []byte{0x06}
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

func LeaveStoreKey(studentId string, leaveId int) []byte {
	leaveID := strconv.Itoa(leaveId)
	key := make([]byte, len(LeaveKey)+len(studentId)+len(leaveID))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(studentId))
	copy(key[len(LeaveKey)+len(studentId):], []byte(leaveID))
	return key
}

func LeaveCounterStoreKey(studentId string) []byte {
	key := make([]byte, len(LeaveCounterKey)+len(studentId))
	copy(key, LeaveCounterKey)
	copy(key[len(LeaveCounterKey):], studentId)
	return key
}

func LeaveStatusStoreKey(studentAddress string, leaveId int) []byte {
	leaveID := strconv.Itoa(leaveId)
	key := make([]byte, len(LeaveStatusKey)+len(studentAddress)+len(leaveID))
	copy(key, LeaveStatusKey)
	copy(key[len(LeaveStatusKey):], []byte(studentAddress))
	copy(key[len(LeaveStatusKey)+len(studentAddress):], []byte(leaveID))
	return key
}

func AcceptLeaveStoreKey(admin string, leaveId string) []byte {
	key := make([]byte, len(AcceptLeaveKey)+len(admin)+len(leaveId))
	copy(key, AcceptLeaveKey)
	copy(key[len(AcceptLeaveKey):], []byte(admin))
	copy(key[len(key):], []byte(leaveId))
	return key
}
