
package cmd

import (
	"os"
	"fmt"
	"log"
	"strings"
	"os/signal"
	cobra "github.com/spf13/cobra"
	nats "github.com/nats-io/go-nats"
	schemas "github.com/polygon-io/go-schemas"
)


var streamCmdOps = struct {
	Channel     string
}{}



var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream a channel from Polygon",
	Run: func(cmd *cobra.Command, args []string) {
		
		if len(streamCmdOps.Channel) == 0 {
			log.Fatal("ERROR: Must specify a channel to stream.. ( T.*, C.*, Q.* )")
		}
		println("Subscribing to:", streamCmdOps.Channel)

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		natsConnection, err := nats.Connect(natsEndpoints, nats.Token( apiKey )); if err != nil {
			log.Fatal( "NATS Connection Error:", err )
		}
		encodedConnection, err := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER); if err != nil {
			log.Fatal("NATS Encoded Connection Error:', err")
		}

		// Currencies:
		if strings.Contains(streamCmdOps.Channel, "C.") {
			encodedConnection.Subscribe(streamCmdOps.Channel, func( tick *schemas.Forex){
				fmt.Println(
					"Pair:", tick.Pair, 
					"\t Exchange:", tick.Exchange, 
					"\t Timestamp:", tick.Timestamp,
					"\t Ask:", tick.Ask, 
					"\t\t Bid:", tick.Bid,
				)
			})
		}
		// Trades:
		if strings.Contains(streamCmdOps.Channel, "T.") {
			encodedConnection.Subscribe(streamCmdOps.Channel, func( tick *schemas.Trade){
				fmt.Println(
					"Symbol:", tick.Symbol, 
					"\t Price:", tick.Price, 
					"\t Size:", tick.Size, 
					"\t Exchange:", tick.Exchange,
					"\t Timestamp:", tick.Timestamp,
					"\t Conditions:", tick.Conditions,
				)
			})
		}
		// Quotes:
		if strings.Contains(streamCmdOps.Channel, "Q.") {
			encodedConnection.Subscribe(streamCmdOps.Channel, func( tick *schemas.Quote){
				fmt.Println(
					"Symbol:", tick.Symbol, 
					"\t Bid:", tick.BidPrice, 
					"\t Ask:", tick.AskPrice, 
					"\t BidExch:", tick.BidExchange, 
					"\t AskExch:", tick.AskExchange,
					"\t BidSize:", tick.BidSize, 
					"\t AskSize:", tick.AskSize,
					"\t Timestamp:", tick.Timestamp,
				)
			})
		}

		for {
			select {
			case <-signals:
				os.Exit(-1)
				return
			}
		}

	},
}


func init() {
	RootCmd.AddCommand(streamCmd)
	streamCmd.Flags().StringVar(&streamCmdOps.Channel, "channel", "", "channel to subscribe to. eg: C.*")
}

