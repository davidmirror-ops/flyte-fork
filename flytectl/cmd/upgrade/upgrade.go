package upgrade

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/flyteorg/flytectl/pkg/util"
	stdlibversion "github.com/flyteorg/flytestdlib/version"

	"github.com/flyteorg/flytectl/pkg/util/githubutil"

	"github.com/flyteorg/flytestdlib/logger"
	"github.com/mouuff/go-rocket-update/pkg/updater"

	cmdCore "github.com/flyteorg/flytectl/cmd/core"
	"github.com/flyteorg/flytectl/pkg/util/platformutil"
	"github.com/spf13/cobra"
)

type Goos string

// Long descriptions are whitespace sensitive when generating docs using sphinx.
const (
	upgradeCmdShort = `Used for upgrade/rollback flyte version`
	upgradeCmdLong  = `
Upgrade flytectl
::

 bin/flytectl upgrade
	
Rollback flytectl binary
::

 bin/flytectl upgrade rollback

Note: Upgrade is not available on windows
`
	rollBackSubCommand = "rollback"
)

var (
	goos = platformutil.Platform(runtime.GOOS)
)

// SelfUpgrade will return self upgrade command
func SelfUpgrade(rootCmd *cobra.Command) map[string]cmdCore.CommandEntry {
	getResourcesFuncs := map[string]cmdCore.CommandEntry{
		"upgrade": {CmdFunc: selfUpgrade, Aliases: []string{"upgrade"}, ProjectDomainNotRequired: true,
			Short: upgradeCmdShort,
			Long:  upgradeCmdLong},
	}
	return getResourcesFuncs
}

func selfUpgrade(ctx context.Context, args []string, cmdCtx cmdCore.CommandContext) error {
	// Check if it's a rollback
	if len(args) == 1 {
		if args[0] == rollBackSubCommand && !isRollBackSupported(goos) {
			return nil
		}
		ext, err := githubutil.FlytectlReleaseConfig.GetExecutable()
		if err != nil {
			return err
		}
		backupBinary := fmt.Sprintf("%s.old", ext)
		if _, err := os.Stat(backupBinary); err != nil {
			return errors.New("flytectl backup doesn't exist. Rollback is not possible")
		}
		return githubutil.FlytectlReleaseConfig.Rollback()
	}

	if isSupported, err := isUpgradeSupported(goos); err != nil {
		return err
	} else if !isSupported {
		return nil
	}

	if message, err := upgrade(githubutil.FlytectlReleaseConfig); err != nil {
		return err
	} else if len(message) > 0 {
		logger.Info(ctx, message)
	}
	return nil
}

func upgrade(u *updater.Updater) (string, error) {
	updateStatus, err := u.Update()
	if err != nil {
		return "", err
	}

	if updateStatus == updater.Updated {
		latestVersion, err := u.GetLatestVersion()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Successfully updated to version %s", latestVersion), nil
	}
	return "", u.Rollback()
}

func isUpgradeSupported(goos platformutil.Platform) (bool, error) {
	latest, err := githubutil.FlytectlReleaseConfig.GetLatestVersion()
	if err != nil {
		return false, err
	}

	if isGreater, err := util.IsVersionGreaterThan(latest, stdlibversion.Version); err != nil {
		return false, err
	} else if !isGreater {
		fmt.Println("You have already latest version of flytectl")
		return false, nil
	}

	message, err := githubutil.GetUpgradeMessage(latest, goos)
	if err != nil {
		return false, err
	}
	if goos.String() == platformutil.Windows.String() || strings.Contains(message, "brew") {
		if len(message) > 0 {
			fmt.Println(message)
		}
		return false, nil
	}
	return true, nil
}

func isRollBackSupported(goos platformutil.Platform) bool {
	if goos.String() == platformutil.Windows.String() {
		fmt.Printf("Flytectl rollback is not available on %s \n", goos.String())
		return false
	}
	return true
}
