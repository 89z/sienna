// Code generated by 'go generate'; DO NOT EDIT.

package msi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modmsi = windows.NewLazySystemDLL("msi.dll")

	procMsiCloseHandle         = modmsi.NewProc("MsiCloseHandle")
	procMsiCreateRecord        = modmsi.NewProc("MsiCreateRecord")
	procMsiDatabaseOpenViewW   = modmsi.NewProc("MsiDatabaseOpenViewW")
	procMsiInstallProductW     = modmsi.NewProc("MsiInstallProductW")
	procMsiOpenDatabaseW       = modmsi.NewProc("MsiOpenDatabaseW")
	procMsiRecordGetFieldCount = modmsi.NewProc("MsiRecordGetFieldCount")
	procMsiRecordGetStringW    = modmsi.NewProc("MsiRecordGetStringW")
	procMsiRecordSetInteger    = modmsi.NewProc("MsiRecordSetInteger")
	procMsiRecordSetStringW    = modmsi.NewProc("MsiRecordSetStringW")
	procMsiViewClose           = modmsi.NewProc("MsiViewClose")
	procMsiViewExecute         = modmsi.NewProc("MsiViewExecute")
	procMsiViewFetch           = modmsi.NewProc("MsiViewFetch")
	procMsiViewGetColumnInfo   = modmsi.NewProc("MsiViewGetColumnInfo")
)

func closeHandle(any Conn) (e error) {
	r0, _, _ := syscall.Syscall(procMsiCloseHandle.Addr(), 1, uintptr(any), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func createRecord(params int) (n int) {
	r0, _, _ := syscall.Syscall(procMsiCreateRecord.Addr(), 1, uintptr(params), 0, 0)
	n = int(r0)
	return
}

func databaseOpenView(database Conn, query string, view *Stmt) (e error) {
	var _p0 *uint16
	_p0, e = syscall.UTF16PtrFromString(query)
	if e != nil {
		return
	}
	return _databaseOpenView(database, _p0, view)
}

func _databaseOpenView(database Conn, query *uint16, view *Stmt) (e error) {
	r0, _, _ := syscall.Syscall(procMsiDatabaseOpenViewW.Addr(), 3, uintptr(database), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(view)))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func InstallProduct(packagePath string, command string) (e error) {
	var _p0 *uint16
	_p0, e = syscall.UTF16PtrFromString(packagePath)
	if e != nil {
		return
	}
	var _p1 *uint16
	_p1, e = syscall.UTF16PtrFromString(command)
	if e != nil {
		return
	}
	return _InstallProduct(_p0, _p1)
}

func _InstallProduct(packagePath *uint16, command *uint16) (e error) {
	r0, _, _ := syscall.Syscall(procMsiInstallProductW.Addr(), 2, uintptr(unsafe.Pointer(packagePath)), uintptr(unsafe.Pointer(command)), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func openDatabase(dbPath string, persist int, database *Conn) (e error) {
	var _p0 *uint16
	_p0, e = syscall.UTF16PtrFromString(dbPath)
	if e != nil {
		return
	}
	return _openDatabase(_p0, persist, database)
}

func _openDatabase(dbPath *uint16, persist int, database *Conn) (e error) {
	r0, _, _ := syscall.Syscall(procMsiOpenDatabaseW.Addr(), 3, uintptr(unsafe.Pointer(dbPath)), uintptr(persist), uintptr(unsafe.Pointer(database)))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func recordGetFieldCount(record int) (n int) {
	r0, _, _ := syscall.Syscall(procMsiRecordGetFieldCount.Addr(), 1, uintptr(record), 0, 0)
	n = int(r0)
	return
}

func recordGetString(record int, field int, buf *uint16, bufSize *int) (e error) {
	r0, _, _ := syscall.Syscall6(procMsiRecordGetStringW.Addr(), 4, uintptr(record), uintptr(field), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(bufSize)), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func recordSetInteger(record int, field int, value int64) (e error) {
	r0, _, _ := syscall.Syscall(procMsiRecordSetInteger.Addr(), 3, uintptr(record), uintptr(field), uintptr(value))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func recordSetString(record int, field int, value string) (e error) {
	var _p0 *uint16
	_p0, e = syscall.UTF16PtrFromString(value)
	if e != nil {
		return
	}
	return _recordSetString(record, field, _p0)
}

func _recordSetString(record int, field int, value *uint16) (e error) {
	r0, _, _ := syscall.Syscall(procMsiRecordSetStringW.Addr(), 3, uintptr(record), uintptr(field), uintptr(unsafe.Pointer(value)))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func viewClose(view Stmt) (e error) {
	r0, _, _ := syscall.Syscall(procMsiViewClose.Addr(), 1, uintptr(view), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func viewExecute(view Stmt, record int) (e error) {
	r0, _, _ := syscall.Syscall(procMsiViewExecute.Addr(), 2, uintptr(view), uintptr(record), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func viewFetch(view Rows, record *int) (e error) {
	r0, _, _ := syscall.Syscall(procMsiViewFetch.Addr(), 2, uintptr(view), uintptr(unsafe.Pointer(record)), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func viewGetColumnInfo(view Rows, columnInfo int, record *int) (e error) {
	r0, _, _ := syscall.Syscall(procMsiViewGetColumnInfo.Addr(), 3, uintptr(view), uintptr(columnInfo), uintptr(unsafe.Pointer(record)))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}
