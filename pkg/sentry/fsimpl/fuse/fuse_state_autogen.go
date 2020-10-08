// automatically generated by stateify.

package fuse

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (c *connection) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.connection"
}

func (c *connection) StateFields() []string {
	return []string{
		"fd",
		"attributeVersion",
		"initialized",
		"initializedChan",
		"connected",
		"connInitError",
		"connInitSuccess",
		"aborted",
		"numWaiting",
		"asyncNum",
		"asyncCongestionThreshold",
		"asyncNumMax",
		"maxRead",
		"maxWrite",
		"maxPages",
		"minor",
		"atomicOTrunc",
		"asyncRead",
		"writebackCache",
		"bigWrites",
		"dontMask",
		"noOpen",
	}
}

func (c *connection) beforeSave() {}

func (c *connection) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	var initializedChanValue bool = c.saveInitializedChan()
	stateSinkObject.SaveValue(3, initializedChanValue)
	stateSinkObject.Save(0, &c.fd)
	stateSinkObject.Save(1, &c.attributeVersion)
	stateSinkObject.Save(2, &c.initialized)
	stateSinkObject.Save(4, &c.connected)
	stateSinkObject.Save(5, &c.connInitError)
	stateSinkObject.Save(6, &c.connInitSuccess)
	stateSinkObject.Save(7, &c.aborted)
	stateSinkObject.Save(8, &c.numWaiting)
	stateSinkObject.Save(9, &c.asyncNum)
	stateSinkObject.Save(10, &c.asyncCongestionThreshold)
	stateSinkObject.Save(11, &c.asyncNumMax)
	stateSinkObject.Save(12, &c.maxRead)
	stateSinkObject.Save(13, &c.maxWrite)
	stateSinkObject.Save(14, &c.maxPages)
	stateSinkObject.Save(15, &c.minor)
	stateSinkObject.Save(16, &c.atomicOTrunc)
	stateSinkObject.Save(17, &c.asyncRead)
	stateSinkObject.Save(18, &c.writebackCache)
	stateSinkObject.Save(19, &c.bigWrites)
	stateSinkObject.Save(20, &c.dontMask)
	stateSinkObject.Save(21, &c.noOpen)
}

func (c *connection) afterLoad() {}

func (c *connection) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.fd)
	stateSourceObject.Load(1, &c.attributeVersion)
	stateSourceObject.Load(2, &c.initialized)
	stateSourceObject.Load(4, &c.connected)
	stateSourceObject.Load(5, &c.connInitError)
	stateSourceObject.Load(6, &c.connInitSuccess)
	stateSourceObject.Load(7, &c.aborted)
	stateSourceObject.Load(8, &c.numWaiting)
	stateSourceObject.Load(9, &c.asyncNum)
	stateSourceObject.Load(10, &c.asyncCongestionThreshold)
	stateSourceObject.Load(11, &c.asyncNumMax)
	stateSourceObject.Load(12, &c.maxRead)
	stateSourceObject.Load(13, &c.maxWrite)
	stateSourceObject.Load(14, &c.maxPages)
	stateSourceObject.Load(15, &c.minor)
	stateSourceObject.Load(16, &c.atomicOTrunc)
	stateSourceObject.Load(17, &c.asyncRead)
	stateSourceObject.Load(18, &c.writebackCache)
	stateSourceObject.Load(19, &c.bigWrites)
	stateSourceObject.Load(20, &c.dontMask)
	stateSourceObject.Load(21, &c.noOpen)
	stateSourceObject.LoadValue(3, new(bool), func(y interface{}) { c.loadInitializedChan(y.(bool)) })
}

func (f *fuseDevice) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.fuseDevice"
}

func (f *fuseDevice) StateFields() []string {
	return []string{}
}

func (f *fuseDevice) beforeSave() {}

func (f *fuseDevice) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *fuseDevice) afterLoad() {}

func (f *fuseDevice) StateLoad(stateSourceObject state.Source) {
}

func (d *DeviceFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.DeviceFD"
}

func (d *DeviceFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"nextOpID",
		"queue",
		"numActiveRequests",
		"completions",
		"writeCursor",
		"writeBuf",
		"writeCursorFR",
		"waitQueue",
		"fullQueueCh",
		"fs",
	}
}

func (d *DeviceFD) beforeSave() {}

func (d *DeviceFD) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	var fullQueueChValue int = d.saveFullQueueCh()
	stateSinkObject.SaveValue(12, fullQueueChValue)
	stateSinkObject.Save(0, &d.vfsfd)
	stateSinkObject.Save(1, &d.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &d.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &d.NoLockFD)
	stateSinkObject.Save(4, &d.nextOpID)
	stateSinkObject.Save(5, &d.queue)
	stateSinkObject.Save(6, &d.numActiveRequests)
	stateSinkObject.Save(7, &d.completions)
	stateSinkObject.Save(8, &d.writeCursor)
	stateSinkObject.Save(9, &d.writeBuf)
	stateSinkObject.Save(10, &d.writeCursorFR)
	stateSinkObject.Save(11, &d.waitQueue)
	stateSinkObject.Save(13, &d.fs)
}

func (d *DeviceFD) afterLoad() {}

func (d *DeviceFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsfd)
	stateSourceObject.Load(1, &d.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &d.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &d.NoLockFD)
	stateSourceObject.Load(4, &d.nextOpID)
	stateSourceObject.Load(5, &d.queue)
	stateSourceObject.Load(6, &d.numActiveRequests)
	stateSourceObject.Load(7, &d.completions)
	stateSourceObject.Load(8, &d.writeCursor)
	stateSourceObject.Load(9, &d.writeBuf)
	stateSourceObject.Load(10, &d.writeCursorFR)
	stateSourceObject.Load(11, &d.waitQueue)
	stateSourceObject.Load(13, &d.fs)
	stateSourceObject.LoadValue(12, new(int), func(y interface{}) { d.loadFullQueueCh(y.(int)) })
}

func (f *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.FilesystemType"
}

func (f *FilesystemType) StateFields() []string {
	return []string{}
}

func (f *FilesystemType) beforeSave() {}

func (f *FilesystemType) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *FilesystemType) afterLoad() {}

func (f *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (f *filesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.filesystemOptions"
}

func (f *filesystemOptions) StateFields() []string {
	return []string{
		"userID",
		"groupID",
		"rootMode",
		"maxActiveRequests",
		"maxRead",
	}
}

func (f *filesystemOptions) beforeSave() {}

func (f *filesystemOptions) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.userID)
	stateSinkObject.Save(1, &f.groupID)
	stateSinkObject.Save(2, &f.rootMode)
	stateSinkObject.Save(3, &f.maxActiveRequests)
	stateSinkObject.Save(4, &f.maxRead)
}

func (f *filesystemOptions) afterLoad() {}

func (f *filesystemOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.userID)
	stateSourceObject.Load(1, &f.groupID)
	stateSourceObject.Load(2, &f.rootMode)
	stateSourceObject.Load(3, &f.maxActiveRequests)
	stateSourceObject.Load(4, &f.maxRead)
}

func (f *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.filesystem"
}

func (f *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
		"conn",
		"opts",
		"umounted",
	}
}

func (f *filesystem) beforeSave() {}

func (f *filesystem) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Filesystem)
	stateSinkObject.Save(1, &f.devMinor)
	stateSinkObject.Save(2, &f.conn)
	stateSinkObject.Save(3, &f.opts)
	stateSinkObject.Save(4, &f.umounted)
}

func (f *filesystem) afterLoad() {}

func (f *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Filesystem)
	stateSourceObject.Load(1, &f.devMinor)
	stateSourceObject.Load(2, &f.conn)
	stateSourceObject.Load(3, &f.opts)
	stateSourceObject.Load(4, &f.umounted)
}

func (i *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.inode"
}

func (i *inode) StateFields() []string {
	return []string{
		"inodeRefs",
		"InodeAttrs",
		"InodeDirectoryNoNewChildren",
		"InodeNoDynamicLookup",
		"InodeNotSymlink",
		"OrderedChildren",
		"dentry",
		"fs",
		"metadataMu",
		"nodeID",
		"locks",
		"size",
		"attributeVersion",
		"attributeTime",
		"version",
		"link",
	}
}

func (i *inode) beforeSave() {}

func (i *inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.inodeRefs)
	stateSinkObject.Save(1, &i.InodeAttrs)
	stateSinkObject.Save(2, &i.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(3, &i.InodeNoDynamicLookup)
	stateSinkObject.Save(4, &i.InodeNotSymlink)
	stateSinkObject.Save(5, &i.OrderedChildren)
	stateSinkObject.Save(6, &i.dentry)
	stateSinkObject.Save(7, &i.fs)
	stateSinkObject.Save(8, &i.metadataMu)
	stateSinkObject.Save(9, &i.nodeID)
	stateSinkObject.Save(10, &i.locks)
	stateSinkObject.Save(11, &i.size)
	stateSinkObject.Save(12, &i.attributeVersion)
	stateSinkObject.Save(13, &i.attributeTime)
	stateSinkObject.Save(14, &i.version)
	stateSinkObject.Save(15, &i.link)
}

func (i *inode) afterLoad() {}

func (i *inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.inodeRefs)
	stateSourceObject.Load(1, &i.InodeAttrs)
	stateSourceObject.Load(2, &i.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(3, &i.InodeNoDynamicLookup)
	stateSourceObject.Load(4, &i.InodeNotSymlink)
	stateSourceObject.Load(5, &i.OrderedChildren)
	stateSourceObject.Load(6, &i.dentry)
	stateSourceObject.Load(7, &i.fs)
	stateSourceObject.Load(8, &i.metadataMu)
	stateSourceObject.Load(9, &i.nodeID)
	stateSourceObject.Load(10, &i.locks)
	stateSourceObject.Load(11, &i.size)
	stateSourceObject.Load(12, &i.attributeVersion)
	stateSourceObject.Load(13, &i.attributeTime)
	stateSourceObject.Load(14, &i.version)
	stateSourceObject.Load(15, &i.link)
}

func (i *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.inodeRefs"
}

func (i *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (i *inodeRefs) beforeSave() {}

func (i *inodeRefs) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.refCount)
}

func (i *inodeRefs) afterLoad() {}

func (i *inodeRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.refCount)
}

func (r *requestList) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.requestList"
}

func (r *requestList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (r *requestList) beforeSave() {}

func (r *requestList) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.head)
	stateSinkObject.Save(1, &r.tail)
}

func (r *requestList) afterLoad() {}

func (r *requestList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.head)
	stateSourceObject.Load(1, &r.tail)
}

func (r *requestEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.requestEntry"
}

func (r *requestEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (r *requestEntry) beforeSave() {}

func (r *requestEntry) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.next)
	stateSinkObject.Save(1, &r.prev)
}

func (r *requestEntry) afterLoad() {}

func (r *requestEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.next)
	stateSourceObject.Load(1, &r.prev)
}

func (r *Request) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.Request"
}

func (r *Request) StateFields() []string {
	return []string{
		"requestEntry",
		"id",
		"hdr",
		"data",
		"payload",
		"async",
		"noReply",
	}
}

func (r *Request) beforeSave() {}

func (r *Request) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.requestEntry)
	stateSinkObject.Save(1, &r.id)
	stateSinkObject.Save(2, &r.hdr)
	stateSinkObject.Save(3, &r.data)
	stateSinkObject.Save(4, &r.payload)
	stateSinkObject.Save(5, &r.async)
	stateSinkObject.Save(6, &r.noReply)
}

func (r *Request) afterLoad() {}

func (r *Request) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.requestEntry)
	stateSourceObject.Load(1, &r.id)
	stateSourceObject.Load(2, &r.hdr)
	stateSourceObject.Load(3, &r.data)
	stateSourceObject.Load(4, &r.payload)
	stateSourceObject.Load(5, &r.async)
	stateSourceObject.Load(6, &r.noReply)
}

func (f *futureResponse) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.futureResponse"
}

func (f *futureResponse) StateFields() []string {
	return []string{
		"opcode",
		"ch",
		"hdr",
		"data",
		"async",
	}
}

func (f *futureResponse) beforeSave() {}

func (f *futureResponse) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.opcode)
	stateSinkObject.Save(1, &f.ch)
	stateSinkObject.Save(2, &f.hdr)
	stateSinkObject.Save(3, &f.data)
	stateSinkObject.Save(4, &f.async)
}

func (f *futureResponse) afterLoad() {}

func (f *futureResponse) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.opcode)
	stateSourceObject.Load(1, &f.ch)
	stateSourceObject.Load(2, &f.hdr)
	stateSourceObject.Load(3, &f.data)
	stateSourceObject.Load(4, &f.async)
}

func (r *Response) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.Response"
}

func (r *Response) StateFields() []string {
	return []string{
		"opcode",
		"hdr",
		"data",
	}
}

func (r *Response) beforeSave() {}

func (r *Response) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.opcode)
	stateSinkObject.Save(1, &r.hdr)
	stateSinkObject.Save(2, &r.data)
}

func (r *Response) afterLoad() {}

func (r *Response) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.opcode)
	stateSourceObject.Load(1, &r.hdr)
	stateSourceObject.Load(2, &r.data)
}

func init() {
	state.Register((*connection)(nil))
	state.Register((*fuseDevice)(nil))
	state.Register((*DeviceFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystemOptions)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*inode)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*requestList)(nil))
	state.Register((*requestEntry)(nil))
	state.Register((*Request)(nil))
	state.Register((*futureResponse)(nil))
	state.Register((*Response)(nil))
}
