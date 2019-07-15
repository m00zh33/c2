package topics

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"gitlab.com/teserakt/c2/internal/cli"
	"gitlab.com/teserakt/c2/pkg/pb"
)

type removeCommand struct {
	cobraCmd        *cobra.Command
	flags           removeCommandFlags
	c2ClientFactory cli.APIClientFactory
}

type removeCommandFlags struct {
	Name string
}

var _ cli.Command = &removeCommand{}

// NewRemoveCommand returns a new command to remove topics
func NewRemoveCommand(c2ClientFactory cli.APIClientFactory) cli.Command {
	removeCmd := &removeCommand{
		c2ClientFactory: c2ClientFactory,
	}

	cobraCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a topic",
		RunE:  removeCmd.run,
	}

	cobraCmd.Flags().StringVar(&removeCmd.flags.Name, "name", "", "The topic name")

	removeCmd.cobraCmd = cobraCmd

	return removeCmd
}

func (c *removeCommand) CobraCmd() *cobra.Command {
	return c.cobraCmd
}

func (c *removeCommand) run(cmd *cobra.Command, args []string) error {
	if len(c.flags.Name) == 0 {
		return fmt.Errorf("flag --name is required")

	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c2Client, err := c.c2ClientFactory.NewClient(cmd)
	if err != nil {
		return fmt.Errorf("cannot create c2 api client: %v", err)
	}
	defer c2Client.Close()

	removeTopicReq := &pb.RemoveTopicRequest{
		Topic: c.flags.Name,
	}

	_, err = c2Client.RemoveTopic(ctx, removeTopicReq)
	if err != nil {
		return fmt.Errorf("failed to remove client: %v", err)
	}

	fmt.Printf("Topic %s removed successfully\n", c.flags.Name)

	return nil
}
