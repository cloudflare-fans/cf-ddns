package tray

import (
	"cf-ddns/cloudflare"
	"cf-ddns/tray/icon"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"os"
	"time"
)

func onReady() {
	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("cfDDNS")
		systray.SetTooltip("Cloudflare DDNS Tool")

		mPortal := systray.AddMenuItem("Control Panel", "open control panel")
		mPortal.Hide()
		mTaskCenter := systray.AddMenuItem("Edit DDNS Tasks", "open task center")
		mTaskCenter.Hide()
		mDoUpdateCFDDNS := systray.AddMenuItem("Update DDNS", "manually update Cloudflare DDNS now")
		systray.AddSeparator()
		mRunning := systray.AddMenuItem("Stop Service [Started]", "stop backend config service, and all DDNS tasks")
		mRunning.Hide()
		systray.AddSeparator()
		mShowDDNSLogs := systray.AddMenuItem("Show DDNS Logs", "show DDNS logs in read-only mode")
		mShowDDNSLogs.Hide()
		mEditDDNSTaskConfigFile := systray.AddMenuItem("Edit DDNS Task Config File", "show DDNS task config file in read-write mode")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Exit", "exit")

		for {
			select {
			case <-mPortal.ClickedCh:
				open.Run("http://localhost:7879/#/")
			case <-mTaskCenter.ClickedCh:
				open.Run("http://localhost:7879/#/tasks")
			case <-mDoUpdateCFDDNS.ClickedCh:
				mDoUpdateCFDDNS.SetTitle("Update DDNS [Executing...]")
				mDoUpdateCFDDNS.SetTooltip("Executing DDNS updating, please wait...")
				mDoUpdateCFDDNS.Disable()
				cloudflare.DoUpdateCFDDNS(
					func() {
						mDoUpdateCFDDNS.SetTitle("Update DDNS [Success!]")
						mDoUpdateCFDDNS.Enable()
						time.Sleep(10 * time.Second)
						mDoUpdateCFDDNS.SetTitle("Update DDNS")
					},
					func(err error) {
						mDoUpdateCFDDNS.SetTitle("Update DDNS [Failed!]")
						mDoUpdateCFDDNS.SetTooltip(err.Error())
						mDoUpdateCFDDNS.Enable()
						time.Sleep(10 * time.Second)
						mDoUpdateCFDDNS.SetTitle("Update DDNS")
					},
				)
			case <-mRunning.ClickedCh:
				mRunning.SetTitle("Start Service [Stopped]")
				mRunning.SetTooltip("start backend config service, and all DDNS tasks")
			case <-mEditDDNSTaskConfigFile.ClickedCh:
				open.Start("./config.yaml")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	now := time.Now()
	os.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
}

func InitTray() {
	systray.Run(onReady, onExit)
}
