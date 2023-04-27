package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/qascade/dcr/lib/service"
)

var (
	runner    string
	dRef      string
	tRef      string
	destOwner string
	pkgPath   string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Transformation and Destination",
	Long:  `CLI Command to run mentioned transformation and destination`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fetch the flags.
		runner = cmd.Flag("runner").Value.String()
		tRef = cmd.Flag("transformation").Value.String()
		dRef = cmd.Flag("destination").Value.String()
		destOwner = cmd.Flag("destinationOwner").Value.String()
		pkgPath = cmd.Flag("pkgpath").Value.String()
		if runner == "" || tRef == "" || dRef == "" || destOwner == "" || pkgPath == "" {
			err := fmt.Errorf("one or more flags are empty")
			return err
		}

		service, err := service.NewService(pkgPath, runner, destOwner, tRef, dRef)
		if err != nil {
			err = fmt.Errorf("err creating new service with package path: %s", pkgPath)
			log.Error(err)
			return err
		}
		err = service.RunCollaborationEvent()
		if err != nil {
			err = fmt.Errorf("err running collaboration event: %s", err)
			log.Error(err)
			return err
		}
		return nil
	},
}

// dcr run --runner Media --transformation t -p package_location -d destinationOwner -dref destination_ref

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("runner", "r", "", "name of the transformation runner")
	runCmd.Flags().StringP("transformation", "t", "", "reference of the transformation")
	runCmd.Flags().StringP("destinationOwner", "o", "", "name of the destination owner")
	runCmd.Flags().StringP("destination", "d", "", "reference of the destination")
	runCmd.Flags().StringP("pkgpath", "p", "", "reference of the destination")
}
