package topics

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"

	"gitlab.com/teserakt/c2/internal/cli"
	"gitlab.com/teserakt/c2/pkg/pb"
)

type listCommand struct {
	cobraCmd        *cobra.Command
	flags           listCommandFlags
	c2ClientFactory cli.APIClientFactory
}

type listCommandFlags struct {
	Offset int64
	Count  int64
}

var _ cli.Command = &listCommand{}

// NewListCommand creates a new command allowing to
// list existing topics
func NewListCommand(c2ClientFactory cli.APIClientFactory) cli.Command {
	listCmd := &listCommand{
		c2ClientFactory: c2ClientFactory,
	}

	cobraCmd := &cobra.Command{
		Use:   "list",
		Short: "List existing topics",
		RunE:  listCmd.run,
	}

	cobraCmd.Flags().SortFlags = false
	cobraCmd.Flags().Int64Var(&listCmd.flags.Offset, "offset", 0, "The offset to start listing topics from")
	cobraCmd.Flags().Int64Var(&listCmd.flags.Count, "count", 0, "The maximum number of topics to return, values <= 0 means all")

	listCmd.cobraCmd = cobraCmd

	return listCmd
}

func (c *listCommand) CobraCmd() *cobra.Command {
	return c.cobraCmd
}

func (c *listCommand) run(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c2Client, err := c.c2ClientFactory.NewClient(cmd)
	if err != nil {
		return fmt.Errorf("cannot create c2 api client: %v", err)
	}
	defer c2Client.Close()

	countResp, err := c2Client.CountTopics(ctx, &pb.CountTopicsRequest{})
	if err != nil {
		return fmt.Errorf("failed to get topic count: %v", err)
	}

	totalCount := countResp.Count
	if c.flags.Count > 0 {
		// Will fetch as many as requested by user, up to maximum number of available topics
		totalCount = int64(math.Min(float64(totalCount), float64(c.flags.Count)))
	}

	currentOffset := c.flags.Offset
	for totalCount > 0 {
		count := int64(math.Min(float64(cli.MaxPageSize), float64(totalCount)))
		req := &pb.GetTopicsRequest{
			Count:  count,
			Offset: currentOffset,
		}

		resp, err := c2Client.GetTopics(ctx, req)
		if err != nil {
			return fmt.Errorf("failed to fetch topics: %v", err)
		}

		currentOffset += count
		totalCount -= count
		for _, topic := range resp.Topics {
			fmt.Println(topic)
		}
	}

	return nil
}
