// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package secexec

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpfSecEvent struct {
	Meta     bpfSecEventMetaT
	Filename [256]uint8
	Buf      [2048]uint8
	Type     uint8
	_        [1]byte
	Conn     bpfConnectionInfoT
	_        [2]byte
}

type bpfSecEventMetaT struct {
	Op       uint8
	_        [3]byte
	Pid      uint32
	Tid      uint32
	Ppid     uint32
	Uid      uint32
	Auid     uint32
	NsPid    uint32
	NsPpid   uint32
	PidNsId  uint32
	_        [4]byte
	TimeNs   uint64
	CapEff   uint64
	CapInh   uint64
	CapPerm  uint64
	CgrpId   uint32
	NetNs    uint32
	CgrpName [128]uint8
	Comm     [16]uint8
}

type bpfSockArgsT struct{ Addr uint64 }

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	KprobeDoTaskDead        *ebpf.ProgramSpec `ebpf:"kprobe_do_task_dead"`
	KprobeSysExecve         *ebpf.ProgramSpec `ebpf:"kprobe_sys_execve"`
	KprobeSysExecveat       *ebpf.ProgramSpec `ebpf:"kprobe_sys_execveat"`
	KprobeTcpConnect        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.ProgramSpec `ebpf:"kprobe_tcp_rcv_established"`
	KprobeWakeUpNewTask     *ebpf.ProgramSpec `ebpf:"kprobe_wake_up_new_task"`
	KretprobeSockAlloc      *ebpf.ProgramSpec `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysConnect     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_connect"`
	SocketHttpFilter        *ebpf.ProgramSpec `ebpf:"socket__http_filter"`
	SyscallEnterExecve      *ebpf.ProgramSpec `ebpf:"syscall_enter_execve"`
	SyscallEnterExecveat    *ebpf.ProgramSpec `ebpf:"syscall_enter_execveat"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	ActiveAcceptArgs    *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActivePids          *ebpf.MapSpec `ebpf:"active_pids"`
	Events              *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections *ebpf.MapSpec `ebpf:"filtered_connections"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	ActiveAcceptArgs    *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.Map `ebpf:"active_connect_args"`
	ActivePids          *ebpf.Map `ebpf:"active_pids"`
	Events              *ebpf.Map `ebpf:"events"`
	FilteredConnections *ebpf.Map `ebpf:"filtered_connections"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActivePids,
		m.Events,
		m.FilteredConnections,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	KprobeDoTaskDead        *ebpf.Program `ebpf:"kprobe_do_task_dead"`
	KprobeSysExecve         *ebpf.Program `ebpf:"kprobe_sys_execve"`
	KprobeSysExecveat       *ebpf.Program `ebpf:"kprobe_sys_execveat"`
	KprobeTcpConnect        *ebpf.Program `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.Program `ebpf:"kprobe_tcp_rcv_established"`
	KprobeWakeUpNewTask     *ebpf.Program `ebpf:"kprobe_wake_up_new_task"`
	KretprobeSockAlloc      *ebpf.Program `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.Program `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysConnect     *ebpf.Program `ebpf:"kretprobe_sys_connect"`
	SocketHttpFilter        *ebpf.Program `ebpf:"socket__http_filter"`
	SyscallEnterExecve      *ebpf.Program `ebpf:"syscall_enter_execve"`
	SyscallEnterExecveat    *ebpf.Program `ebpf:"syscall_enter_execveat"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.KprobeDoTaskDead,
		p.KprobeSysExecve,
		p.KprobeSysExecveat,
		p.KprobeTcpConnect,
		p.KprobeTcpRcvEstablished,
		p.KprobeWakeUpNewTask,
		p.KretprobeSockAlloc,
		p.KretprobeSysAccept4,
		p.KretprobeSysConnect,
		p.SocketHttpFilter,
		p.SyscallEnterExecve,
		p.SyscallEnterExecveat,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_x86.o
var _BpfBytes []byte
