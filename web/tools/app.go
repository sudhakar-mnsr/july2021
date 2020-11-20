package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var (
	changeLogURL                    = "https://github.com/urfave/cli/blob/master/CHANGELOG.md"
	appActionDeprecationURL         = fmt.Sprintf("%s#deprecated-cli-app-action-signature", changeLogURL)
	runAndExitOnErrorDeprecationURL = fmt.Sprintf("%s#deprecated-cli-app-runandexitonerror", changeLogURL)

	contactSysadmin = "This is an error in the application.  Please contact the distributor of this application if this is not you."

	errInvalidActionType = NewExitError("ERROR invalid Action type. "+
		fmt.Sprintf("Must be `func(*Context`)` or `func(*Context) error).  %s", contactSysadmin)+
		fmt.Sprintf("See %s", appActionDeprecationURL), 2)
)

// App is the main structure of a cli application. It is recommended that
// an app be created with the cli.NewApp() function
type App struct {
	// The name of the program. Defaults to path.Base(os.Args[0])
	Name string
	// Full name of command for help, defaults to Name
	HelpName string
	// Description of the program.
	Usage string
	// Text to override the USAGE section of help
	UsageText string
