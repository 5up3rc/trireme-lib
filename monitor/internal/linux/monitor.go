package linuxmonitor

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aporeto-inc/trireme-lib/common"
	"github.com/aporeto-inc/trireme-lib/monitor/config"
	"github.com/aporeto-inc/trireme-lib/monitor/registerer"
	"github.com/aporeto-inc/trireme-lib/utils/cgnetcls"
)

// LinuxMonitor captures all the monitor processor information
// It implements the EventProcessor interface of the rpc monitor
type LinuxMonitor struct {
	proc *linuxProcessor
}

// New returns a new implmentation of a monitor implmentation
func New() *LinuxMonitor {

	return &LinuxMonitor{
		proc: &linuxProcessor{},
	}
}

// Run implements Implementation interface
func (l *LinuxMonitor) Run(ctx context.Context) error {

	if err := l.proc.config.IsComplete(); err != nil {
		return fmt.Errorf("linux %t: %s", l.proc.host, err)
	}

	if err := l.ReSync(ctx); err != nil {
		return err
	}

	return nil
}

// SetupConfig provides a configuration to implmentations. Every implmentation
// can have its own config type.
func (l *LinuxMonitor) SetupConfig(registerer registerer.Registerer, cfg interface{}) error {

	defaultConfig := DefaultConfig(false)
	if cfg == nil {
		cfg = defaultConfig
	}

	linuxConfig, ok := cfg.(*Config)
	if !ok {
		return fmt.Errorf("Invalid configuration specified")
	}

	if registerer != nil {
		if err := registerer.RegisterProcessor(common.LinuxProcessPU, l.proc); err != nil {
			return err
		}
	}

	// Setup defaults
	linuxConfig = SetupDefaultConfig(linuxConfig)

	// Setup config
	l.proc.host = linuxConfig.Host
	l.proc.netcls = cgnetcls.NewCgroupNetController(common.TriremeCgroupPath, linuxConfig.ReleasePath)

	l.proc.regStart = regexp.MustCompile("^[a-zA-Z0-9_].{0,11}$")
	l.proc.regStop = regexp.MustCompile("^/trireme/[a-zA-Z0-9_].{0,11}$")

	l.proc.metadataExtractor = linuxConfig.EventMetadataExtractor
	if l.proc.metadataExtractor == nil {
		return fmt.Errorf("Unable to setup a metadata extractor")
	}

	return nil
}

// SetupHandlers sets up handlers for monitors to invoke for various events such as
// processing unit events and synchronization events. This will be called before Start()
// by the consumer of the monitor
func (l *LinuxMonitor) SetupHandlers(m *config.ProcessorConfig) {

	l.proc.config = m
}

// ReSync instructs the monitor to do a resync.
func (l *LinuxMonitor) ReSync(ctx context.Context) error {

	return l.proc.ReSync(ctx, nil)
}
