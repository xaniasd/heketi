package cmds

import (
	"errors"

	client "github.com/heketi/heketi/client/api/go-client"
	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/spf13/cobra"
)

var (
	slaveHost, slaveVolume string
	volumeID               string
)

func initGeoRepCommand(command *cobra.Command) {
	volumeCommand.AddCommand(geoReplicationCommand)
	geoReplicationCommand.AddCommand(
		geoReplicationCreateCommand,
		geoReplicationConfigCommand,
	)

	// Flags
	geoReplicationCommand.PersistentFlags().StringVar(&slaveHost, "slave-host", "", "The host of the slave volume")
	geoReplicationCommand.PersistentFlags().StringVar(&slaveVolume, "slave-volume", "", "The volume name of the geo-replication target")
	geoReplicationCreateCommand.Flags().Bool("force", false, "\n\tForce creation")
	geoReplicationCreateCommand.Flags().String("option", "no-verify", "\n\tSet option for create command")
	geoReplicationCreateCommand.Flags().Int("ssh-port", 0, "\n\tThe gluster SSH port on the slave host")
}

var geoReplicationCommand = &cobra.Command{
	Use:   "georep",
	Short: "Volume GeoReplication Management",
	Long:  "Heketi Volume GeoReplication Management",
}

var geoReplicationCreateCommand = &cobra.Command{
	Use:     "create",
	Short:   "Create session",
	Long:    "Create GeoReplication session",
	Example: "  $ heketi-cli volume 886a86a868711bef83001 georep --slave-host=blah --slave-volume=23423423 create",
	RunE: func(cmd *cobra.Command, args []string) error {
		//ensure proper number of args
		if len(cmd.Flags().Args()) < 1 {
			return errors.New("Volume id missing")
		}

		volumeID := cmd.Flags().Arg(0)
		if volumeID == "" {
			return errors.New("Volume id missing")
		}

		forceFlag, _ := cmd.Flags().GetBool("force")
		optionFlag, _ := cmd.Flags().GetString("option")
		actionParams := make(map[string]string)
		actionParams["option"] = optionFlag
		sshPortFlag, _ := cmd.Flags().GetInt("ssh-port")

		// Create a client
		heketi := client.NewClient(options.Url, options.User, options.Key)
		req := api.GeoReplicationRequest{
			Action:       api.GeoReplicationActionCreate,
			ActionParams: actionParams,
			GeoReplicationInfo: api.GeoReplicationInfo{
				SlaveHost:    slaveHost,
				SlaveVolume:  slaveVolume,
				SlaveSSHPort: sshPortFlag,
			},
			Force: forceFlag,
		}
		heketi.GeoReplicationCreate(volumeID, &req)
		return nil
	},
}

var geoReplicationConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "Configure session",
	Long:  "Configure GeoReplication session",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
