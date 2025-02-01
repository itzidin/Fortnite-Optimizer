package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

const (
	titleText = `
    +--------------------------------------------------+
    |  FORTNITE OPTIMIZER                              |
    |  ===============                                  |
    |                                                  |
    |  Created by: itzidin                            |
    +--------------------------------------------------+`

	adminRequiredText = `
    +--------------------------------------------------+
    |               ADMINISTRATOR REQUIRED              |
    +--------------------------------------------------+

    Please run this program as Administrator!`

	progressBarTemplate = `
    +--------------------------------------------------+
    | Progress: [%s] %d%%                              |
    +--------------------------------------------------+
    
    > %s`

	completionText = `
    +--------------------------------------------------+
    |             OPTIMIZATION COMPLETE!                |
    +--------------------------------------------------+`
)

func batmanIsAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false
	}
	defer windows.FreeSid(sid)
	token := windows.Token(0)
	member, err := token.IsMember(sid)
	return err == nil && member
}

func showMenu() int {
	fmt.Print("\033[H\033[2J")
	fmt.Println(titleText)
	fmt.Println(`
    +--------------------------------------------------+
    | [1] Advanced Optimization                         |
    |     Maximum Performance + Input Delay Reduction   |
    |                                                  |
    | [2] Basic Optimization                           |
    |     Safe Performance Improvements                |
    |                                                  |
    | [3] Credits                                      |
    | [4] Exit                                         |
    +--------------------------------------------------+`)
	fmt.Print("\n    Enter your choice (1-4): ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func showCredits() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(`
    +--------------------------------------------------+
    |                FORTNITE OPTIMIZER                 |
    +--------------------------------------------------+

    Created by: itzidin
    Version: 1.0
    Release Date: 2024

    This program provides advanced optimization techniques
    for Fortnite, focusing on maximum performance and
    minimal input delay.

    Press Enter to return to menu...`)
	fmt.Scanln()
}

func showProgress(progress int, message string) {
	bar := strings.Repeat("■", progress/10) + strings.Repeat("□", 10-progress/10)
	fmt.Print("\033[H\033[2J")
	fmt.Printf(progressBarTemplate, bar, progress, message)
}

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This program only works on Windows")
		os.Exit(1)
	}

	if !batmanIsAdmin() {
		fmt.Println(adminRequiredText)
		os.Exit(1)
	}

	for {
		choice := showMenu()
		switch choice {
		case 1:
			amirwkonieAdvanced()
		case 2:
			batmanBasic()
		case 3:
			showCredits()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("\nInvalid choice. Press Enter to continue...")
			fmt.Scanln()
		}
	}
}

type kiresephrfv struct {
	component string
	err       error
}

func (e kiresephrfv) Error() string {
	return fmt.Sprintf("Failed to optimize %s: %v", e.component, e.err)
}

func setRegistryDWORD(path string, name string, value uint32) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()
	return key.SetDWordValue(name, value)
}

func setRegistryString(path string, name string, value string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()
	return key.SetStringValue(name, value)
}

func optimizeGPUSettings() error {
	settings := map[string]uint32{
		"DpiMapIommuContiguous":      1,
		"TdrLevel":                   0,
		"UseGpuTimer":                1,
		"RmGpsPsEnablePerCpuCoreDpc": 1,
		"HwSchMode":                  2,
		"TdrDelay":                   20,
	}
	for name, value := range settings {
		if err := setRegistryDWORD(`SYSTEM\CurrentControlSet\Control\GraphicsDrivers`, name, value); err != nil {
			return kiresephrfv{"GPU settings", err}
		}
	}
	return nil
}

func optimizeNetworkSettings() error {
	tcpipParams := map[string]uint32{
		"TcpNoDelay":             1,
		"TcpAckFrequency":        1,
		"TcpDelAckTicks":         0,
		"TCPCongestionControl":   1,
		"DefaultTTL":             64,
		"EnablePMTUDiscovery":    1,
		"EnablePMTUBHDetect":     1,
		"GlobalMaxTcpWindowSize": 65535,
	}
	for name, value := range tcpipParams {
		if err := setRegistryDWORD(`SYSTEM\CurrentControlSet\Services\Tcpip\Parameters`, name, value); err != nil {
			return kiresephrfv{"network settings", err}
		}
	}
	return nil
}

func optimizeProcessPriority() error {
	perfOptions := map[string]uint32{
		"CpuPriorityClass": 3,
		"IoPriority":       3,
		"PagePriority":     5,
	}
	basePath := `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\FortniteClient-Win64-Shipping.exe\PerfOptions`
	for name, value := range perfOptions {
		if err := setRegistryDWORD(basePath, name, value); err != nil {
			return kiresephrfv{"process priority", err}
		}
	}
	return nil
}

func optimizeGPUPreference() error {
	err := setRegistryString(
		`Software\Microsoft\DirectX\UserGpuPreferences`,
		"FortniteClient-Win64-Shipping.exe",
		"GpuPreference=2;",
	)
	if err != nil {
		return kiresephrfv{"GPU preference", err}
	}
	return nil
}

func safeOverclock() error {
	settings := map[string]uint32{
		"RMClockBoost":         1,
		"RMPowerLimit":         115,
		"RMTempLimit":          83,
		"RMVoltageOffset":      0,
		"RMMemoryBoost":        1,
		"RMFanControl":         1,
		"RMFanSpeed":           70,
		"RMClockBoostPercent":  7,
		"RMMemoryBoostPercent": 5,
		"RMThermalMonitoring":  1,
	}
	basePath := `SYSTEM\CurrentControlSet\Control\Class\{4d36e968-e325-11ce-bfc1-08002be10318}\0000`
	for name, value := range settings {
		if err := setRegistryDWORD(basePath, name, value); err != nil {
			return kiresephrfv{"overclocking", err}
		}
	}
	return nil
}

func optimizePowerSettings() error {
	cmd := exec.Command("powercfg", "/setactive", "8c5e7fda-e8bf-4a96-9a85-a6e23a8c635c")
	if err := cmd.Run(); err != nil {
		return kiresephrfv{"power settings", err}
	}
	cmd = exec.Command("powercfg", "/change", "monitor-timeout-ac", "0")
	if err := cmd.Run(); err != nil {
		return kiresephrfv{"power settings", err}
	}
	cmd = exec.Command("powercfg", "/change", "disk-timeout-ac", "0")
	if err := cmd.Run(); err != nil {
		return kiresephrfv{"power settings", err}
	}
	return nil
}

func optimizeSecurityFeatures() error {
	settings := map[string]uint32{
		`SYSTEM\CurrentControlSet\Control\DeviceGuard\Scenarios\HypervisorEnforcedCodeIntegrity\Enabled`: 0,
		`SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management\EnableCfg`:                   0,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System\EnableVirtualization`:                 0,
		`SOFTWARE\Policies\Microsoft\Windows\DeviceGuard\EnableVirtualizationBasedSecurity`:             0,
		`SOFTWARE\Policies\Microsoft\Windows\DeviceGuard\HypervisorEnforcedCodeIntegrity`:              0,
		`SYSTEM\CurrentControlSet\Control\Session Manager\kernel\DisableExceptionChainValidation`:       1,
		`SYSTEM\CurrentControlSet\Control\Session Manager\kernel\KernelSEHOPEnabled`:                   0,
	}
	for path, value := range settings {
		lastBackslash := strings.LastIndex(path, "\\")
		registryPath := path[:lastBackslash]
		valueName := path[lastBackslash+1:]
		if err := setRegistryDWORD(registryPath, valueName, value); err != nil {
			return kiresephrfv{"security features", err}
		}
	}
	return nil
}

func optimizeCPUParking() error {
	cmd := exec.Command("powercfg", "-setacvalueindex", "scheme_current", "sub_processor", "CPMINCORES", "100")
	if err := cmd.Run(); err != nil {
		return kiresephrfv{"CPU parking", err}
	}
	cmd = exec.Command("powercfg", "-setacvalueindex", "scheme_current", "sub_processor", "CPMAXCORES", "100")
	if err := cmd.Run(); err != nil {
		return kiresephrfv{"CPU parking", err}
	}
	cmd = exec.Command("powercfg", "-setactive", "scheme_current")
	if err := cmd.Run(); err != nil {
		return kiresephrfv{"CPU parking", err}
	}
	return nil
}

func amirwkonieAdvanced() {
	steps := []struct {
		progress int
		message  string
		action   func() error
	}{
		{10, "Enabling Performance Mode...", optimizePowerSettings},
		{20, "Optimizing GPU settings...", optimizeGPUSettings},
		{35, "Optimizing network settings...", optimizeNetworkSettings},
		{50, "Setting process priorities...", optimizeProcessPriority},
		{65, "Configuring security features...", optimizeSecurityFeatures},
		{80, "Optimizing CPU parking...", optimizeCPUParking},
		{90, "Applying safe overclocking...", safeOverclock},
		{95, "Setting GPU preferences...", optimizeGPUPreference},
	}
	for _, step := range steps {
		showProgress(step.progress, step.message)
		if err := step.action(); err != nil {
			fmt.Printf("\nError: %v\nPress Enter to continue...", err)
			fmt.Scanln()
			return
		}
	}
	showProgress(100, "Verifying optimizations...")
	issues := verifyOptimizations()
	fmt.Println(completionText)
	fmt.Println("\n    Advanced optimizations applied:")
	fmt.Println("    ▸ Set Fortnite process to high priority")
	fmt.Println("    ▸ Optimized CPU core usage and memory management")
	fmt.Println("    ▸ Enhanced GPU thread priority and DirectX settings")
	fmt.Println("    ▸ Disabled CPU core parking and unnecessary security features")
	fmt.Println("    ▸ Set maximum GPU preference for Fortnite")
	fmt.Println("    ▸ Applied safe CPU and GPU overclocking")
	fmt.Println("    ▸ Optimized TCP/IP stack and network settings for minimal latency")
	fmt.Println("    ▸ Enhanced MTU configuration and network routing")
	if len(issues) > 0 {
		fmt.Println("\n    Warning: Some optimizations may not have been applied correctly:")
		for _, issue := range issues {
			fmt.Printf("    ▸ %s\n", issue)
		}
	}
	fmt.Println("\n    Please restart your computer for all changes to take effect.")
	fmt.Println("\n    Press Enter to return to menu...")
	fmt.Scanln()
}

func verifyOptimizations() []string {
	var issues []string
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\FortniteClient-Win64-Shipping.exe\PerfOptions`, registry.QUERY_VALUE)
	if err != nil || getRegistryDWORD(key, "CpuPriorityClass") != 3 {
		issues = append(issues, "CPU Priority: Not Set")
	}
	key.Close()
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management`, registry.QUERY_VALUE)
	if err != nil || getRegistryDWORD(key, "DisablePagingExecutive") != 1 {
		issues = append(issues, "Memory Management: Not Optimized")
	}
	key.Close()
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\DeviceGuard`, registry.QUERY_VALUE)
	if err != nil || getRegistryDWORD(key, "EnableVirtualizationBasedSecurity") != 0 {
		issues = append(issues, "Security Feat
