package sender

import (
	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"github.com/mpuzanov/otus-go/calendar/internal/sender"
	"github.com/mpuzanov/otus-go/calendar/pkg/logger"

	"log"

	"github.com/spf13/cobra"
)

var cfgPath string

func init() {
	ServerCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
}

var (
	// ServerCmd .
	ServerCmd = &cobra.Command{
		Use:   "sender",
		Short: "Run sender",
		Run:   senderServerStart,
	}
)

func senderServerStart(cmd *cobra.Command, args []string) {

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить %s: %s", cfgPath, err)
	}
	logger := logger.NewLogger(cfg.Log)

	mq, err := sender.NewSender(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	if err := mq.ReadSendMessage(); err != nil {
		log.Fatal(err)
	}

}
