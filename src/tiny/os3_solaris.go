// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

//go:cgo_export_dynamic runtime.end _end
//go:cgo_export_dynamic runtime.etext _etext
//go:cgo_export_dynamic runtime.edata _edata

//go:cgo_import_dynamic libc____errno ___errno "libc.so"
//go:cgo_import_dynamic libc_clock_gettime clock_gettime "libc.so"
//go:cgo_import_dynamic libc_close close "libc.so"
//go:cgo_import_dynamic libc_exit exit "libc.so"
//go:cgo_import_dynamic libc_fstat fstat "libc.so"
//go:cgo_import_dynamic libc_getcontext getcontext "libc.so"
//go:cgo_import_dynamic libc_getrlimit getrlimit "libc.so"
//go:cgo_import_dynamic libc_kill kill "libc.so"
//go:cgo_import_dynamic libc_madvise madvise "libc.so"
//go:cgo_import_dynamic libc_malloc malloc "libc.so"
//go:cgo_import_dynamic libc_mmap mmap "libc.so"
//go:cgo_import_dynamic libc_munmap munmap "libc.so"
//go:cgo_import_dynamic libc_open open "libc.so"
//go:cgo_import_dynamic libc_pthread_attr_destroy pthread_attr_destroy "libc.so"
//go:cgo_import_dynamic libc_pthread_attr_getstack pthread_attr_getstack "libc.so"
//go:cgo_import_dynamic libc_pthread_attr_init pthread_attr_init "libc.so"
//go:cgo_import_dynamic libc_pthread_attr_setdetachstate pthread_attr_setdetachstate "libc.so"
//go:cgo_import_dynamic libc_pthread_attr_setstack pthread_attr_setstack "libc.so"
//go:cgo_import_dynamic libc_pthread_create pthread_create "libc.so"
//go:cgo_import_dynamic libc_raise raise "libc.so"
//go:cgo_import_dynamic libc_read read "libc.so"
//go:cgo_import_dynamic libc_select select "libc.so"
//go:cgo_import_dynamic libc_sched_yield sched_yield "libc.so"
//go:cgo_import_dynamic libc_sem_init sem_init "libc.so"
//go:cgo_import_dynamic libc_sem_post sem_post "libc.so"
//go:cgo_import_dynamic libc_sem_reltimedwait_np sem_reltimedwait_np "libc.so"
//go:cgo_import_dynamic libc_sem_wait sem_wait "libc.so"
//go:cgo_import_dynamic libc_setitimer setitimer "libc.so"
//go:cgo_import_dynamic libc_sigaction sigaction "libc.so"
//go:cgo_import_dynamic libc_sigaltstack sigaltstack "libc.so"
//go:cgo_import_dynamic libc_sigprocmask sigprocmask "libc.so"
//go:cgo_import_dynamic libc_sysconf sysconf "libc.so"
//go:cgo_import_dynamic libc_usleep usleep "libc.so"
//go:cgo_import_dynamic libc_write write "libc.so"
//go:cgo_import_dynamic libc_brk brk "libc.so"

//go:linkname libc____errno libc____errno
//go:linkname libc_clock_gettime libc_clock_gettime
//go:linkname libc_close libc_close
//go:linkname libc_exit libc_exit
//go:linkname libc_fstat libc_fstat
//go:linkname libc_getcontext libc_getcontext
//go:linkname libc_getrlimit libc_getrlimit
//go:linkname libc_kill libc_kill
//go:linkname libc_madvise libc_madvise
//go:linkname libc_malloc libc_malloc
//go:linkname libc_mmap libc_mmap
//go:linkname libc_munmap libc_munmap
//go:linkname libc_open libc_open
//go:linkname libc_pthread_attr_destroy libc_pthread_attr_destroy
//go:linkname libc_pthread_attr_getstack libc_pthread_attr_getstack
//go:linkname libc_pthread_attr_init libc_pthread_attr_init
//go:linkname libc_pthread_attr_setdetachstate libc_pthread_attr_setdetachstate
//go:linkname libc_pthread_attr_setstack libc_pthread_attr_setstack
//go:linkname libc_pthread_create libc_pthread_create
//go:linkname libc_raise libc_raise
//go:linkname libc_read libc_read
//go:linkname libc_select libc_select
//go:linkname libc_sched_yield libc_sched_yield
//go:linkname libc_sem_init libc_sem_init
//go:linkname libc_sem_post libc_sem_post
//go:linkname libc_sem_reltimedwait_np libc_sem_reltimedwait_np
//go:linkname libc_sem_wait libc_sem_wait
//go:linkname libc_setitimer libc_setitimer
//go:linkname libc_sigaction libc_sigaction
//go:linkname libc_sigaltstack libc_sigaltstack
//go:linkname libc_sigprocmask libc_sigprocmask
//go:linkname libc_sysconf libc_sysconf
//go:linkname libc_usleep libc_usleep
//go:linkname libc_write libc_write
//go:linkname libc_brk libc_brk

var (
	libc____errno,
	libc_clock_gettime,
	libc_close,
	libc_exit,
	libc_fstat,
	libc_getcontext,
	libc_getrlimit,
	libc_kill,
	libc_madvise,
	libc_malloc,
	libc_mmap,
	libc_munmap,
	libc_open,
	libc_pthread_attr_destroy,
	libc_pthread_attr_getstack,
	libc_pthread_attr_init,
	libc_pthread_attr_setdetachstate,
	libc_pthread_attr_setstack,
	libc_pthread_create,
	libc_raise,
	libc_read,
	libc_sched_yield,
	libc_select,
	libc_sem_init,
	libc_sem_post,
	libc_sem_reltimedwait_np,
	libc_sem_wait,
	libc_setitimer,
	libc_sigaction,
	libc_sigaltstack,
	libc_sigprocmask,
	libc_sysconf,
	libc_usleep,
	libc_write,
	libc_brk libcFunc
)

//go:nosplit
func closefd(fd int32) int32 {
	return int32(sysvicall1(&libc_close, uintptr(fd)))
}

//go:nosplit
func exit(r int32) {
	sysvicall1(&libc_exit, uintptr(r))
}

//go:nosplit
func getcontext(context *ucontext) /* int32 */ {
	sysvicall1(&libc_getcontext, uintptr(unsafe.Pointer(context)))
}

//go:nosplit
func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
	sysvicall3(&libc_madvise, uintptr(addr), uintptr(n), uintptr(flags))
}

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) unsafe.Pointer {
	p, err := doMmap(uintptr(addr), n, uintptr(prot), uintptr(flags), uintptr(fd), uintptr(off))
	if p == ^uintptr(0) {
		return unsafe.Pointer(err)
	}
	return unsafe.Pointer(p)
}

//go:nosplit
func doMmap(addr, n, prot, flags, fd, off uintptr) (uintptr, uintptr) {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(&libc_mmap))
	libcall.n = 6
	libcall.args = uintptr(noescape(unsafe.Pointer(&addr)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(&libcall))
	return libcall.r1, libcall.err
}

//go:nosplit
func munmap(addr unsafe.Pointer, n uintptr) {
	sysvicall2(&libc_munmap, uintptr(addr), uintptr(n))
}

func nanotime1()

//go:nosplit
func nanotime() int64 {
	return int64(sysvicall0((*libcFunc)(unsafe.Pointer(funcPC(nanotime1)))))
}

//go:nosplit
func open(path *byte, mode, perm int32) int32 {
	return int32(sysvicall3(&libc_open, uintptr(unsafe.Pointer(path)), uintptr(mode), uintptr(perm)))
}

func pthread_attr_destroy(attr *pthreadattr) int32 {
	return int32(sysvicall1(&libc_pthread_attr_destroy, uintptr(unsafe.Pointer(attr))))
}

func pthread_attr_getstack(attr *pthreadattr, addr unsafe.Pointer, size *uint64) int32 {
	return int32(sysvicall3(&libc_pthread_attr_getstack, uintptr(unsafe.Pointer(attr)), uintptr(addr), uintptr(unsafe.Pointer(size))))
}

func pthread_attr_init(attr *pthreadattr) int32 {
	return int32(sysvicall1(&libc_pthread_attr_init, uintptr(unsafe.Pointer(attr))))
}

func pthread_attr_setdetachstate(attr *pthreadattr, state int32) int32 {
	return int32(sysvicall2(&libc_pthread_attr_setdetachstate, uintptr(unsafe.Pointer(attr)), uintptr(state)))
}

func pthread_attr_setstack(attr *pthreadattr, addr uintptr, size uint64) int32 {
	return int32(sysvicall3(&libc_pthread_attr_setstack, uintptr(unsafe.Pointer(attr)), uintptr(addr), uintptr(size)))
}

func pthread_create(thread *pthread, attr *pthreadattr, fn uintptr, arg unsafe.Pointer) int32 {
	return int32(sysvicall4(&libc_pthread_create, uintptr(unsafe.Pointer(thread)), uintptr(unsafe.Pointer(attr)), uintptr(fn), uintptr(arg)))
}

//go:nosplit
//go:nowritebarrierrec
func raise(sig int32) /* int32 */ {
	sysvicall1(&libc_raise, uintptr(sig))
}

func raiseproc(sig int32) /* int32 */ {
	pid := sysvicall0(&libc_getpid)
	sysvicall2(&libc_kill, pid, uintptr(sig))
}

//go:nosplit
func read(fd int32, buf unsafe.Pointer, nbyte int32) int32 {
	return int32(sysvicall3(&libc_read, uintptr(fd), uintptr(buf), uintptr(nbyte)))
}

//go:nosplit
func sem_init(sem *semt, pshared int32, value uint32) int32 {
	return int32(sysvicall3(&libc_sem_init, uintptr(unsafe.Pointer(sem)), uintptr(pshared), uintptr(value)))
}

//go:nosplit
func sem_post(sem *semt) int32 {
	return int32(sysvicall1(&libc_sem_post, uintptr(unsafe.Pointer(sem))))
}

//go:nosplit
func sem_reltimedwait_np(sem *semt, timeout *timespec) int32 {
	return int32(sysvicall2(&libc_sem_reltimedwait_np, uintptr(unsafe.Pointer(sem)), uintptr(unsafe.Pointer(timeout))))
}

//go:nosplit
func sem_wait(sem *semt) int32 {
	return int32(sysvicall1(&libc_sem_wait, uintptr(unsafe.Pointer(sem))))
}

func setitimer(which int32, value *itimerval, ovalue *itimerval) /* int32 */ {
	sysvicall3(&libc_setitimer, uintptr(which), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(ovalue)))
}

//go:nosplit
//go:nowritebarrierrec
func sigaction(sig int32, act *sigactiont, oact *sigactiont) /* int32 */ {
	sysvicall3(&libc_sigaction, uintptr(sig), uintptr(unsafe.Pointer(act)), uintptr(unsafe.Pointer(oact)))
}

//go:nosplit
//go:nowritebarrierrec
func sigaltstack(ss *sigaltstackt, oss *sigaltstackt) /* int32 */ {
	sysvicall2(&libc_sigaltstack, uintptr(unsafe.Pointer(ss)), uintptr(unsafe.Pointer(oss)))
}

//go:nosplit
//go:nowritebarrierrec
func sigprocmask(how int32, set *sigset, oset *sigset) /* int32 */ {
	sysvicall3(&libc_sigprocmask, uintptr(how), uintptr(unsafe.Pointer(set)), uintptr(unsafe.Pointer(oset)))
}

func sysconf(name int32) int64 {
	return int64(sysvicall1(&libc_sysconf, uintptr(name)))
}

func usleep1(uint32)

//go:nosplit
func usleep(µs uint32) {
	usleep1(µs)
}

//go:nosplit
func write(fd uintptr, buf unsafe.Pointer, nbyte int32) int32 {
	return int32(sysvicall3(&libc_write, uintptr(fd), uintptr(buf), uintptr(nbyte)))
}

func osyield1()

//go:nosplit
func osyield() {
	_g_ := getg()

	// Check the validity of m because we might be called in cgo callback
	// path early enough where there isn't a m available yet.
	if _g_ != nil && _g_.m != nil {
		sysvicall0(&libc_sched_yield)
		return
	}
	osyield1()
}

func brk(addr uintptr) uintptr {
	return sysvicall1(&libc_brk, addr)
}