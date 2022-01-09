package monitor

import "github.com/mosn/holmes"

func Init(path string) {
	h, _ := holmes.New(
		holmes.WithCollectInterval("5s"),
		holmes.WithCoolDown("1m"),
		holmes.WithDumpPath(path),
		holmes.WithTextDump(),
		holmes.WithGoroutineDump(10, 25, 100),
		holmes.WithMemDump(3, 25, 80),
		holmes.WithCPUDump(20, 25, 80),
		holmes.WithThreadDump(10, 25, 100),
		holmes.WithCGroup(true),
	)
	h.EnableGoroutineDump()
	h.EnableMemDump()
	h.EnableCPUDump()
	h.EnableThreadDump()
	h.Start()
}
