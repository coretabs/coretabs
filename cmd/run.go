package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/TChi91/coretabs/config"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run your servers",
	Long: `Run you Front-End and Back-End servers with 
one command only: coretabs run.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()

		// err := viper.Unmarshal(&config)
		// if err != nil {
		// 	log.Fatalf("unable to decode into struct, %v", err)
		// }

		var server int
		fmt.Print(`What server you want to start:
(1) ==> Front-End server
(2) ==> Back-End server: `)
		if _, err := fmt.Scanf("%d", &server); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		var port int
		fmt.Print(`Which port you want to use: `)

		switch server {
		case 1:
			fmt.Print(`Default is 8080: `)
			inputPort, err := readPort(&port)
			if err != nil {
				return
			}

			if inputPort != 0 {
				config.FrontEnd.Port = port
				frontEndServer(config)
			} else {
				frontEndServer(config)
			}

		case 2:
			fmt.Print(`Default is 8000: `)
			inputPort, err := readPort(&port)
			if err != nil {
				return
			}

			if inputPort != 0 {
				config.BackEnd.Port = port
				backEndServer(config)
			} else {
				backEndServer(config)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func backEndServer(config config.AppConfig) error {
	backCmd := fmt.Sprintf("source venv/bin/activate; python manage.py runserver %v", config.BackEnd.Port)

	execBackCmd := exec.Command("bash", "-c", backCmd)

	execBackCmd.Stdout = os.Stdout
	execBackCmd.Stderr = os.Stderr

	must(execBackCmd.Start())

	must(execBackCmd.Wait())

	return nil

}

func frontEndServer(config config.AppConfig) error {
	frontCmd := fmt.Sprintf("npm run serve -- --port %v", config.FrontEnd.Port)

	execfrontCmd := exec.Command("bash", "-c", frontCmd)

	execfrontCmd.Stdout = os.Stdout
	execfrontCmd.Stderr = os.Stderr

	err := execfrontCmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = execfrontCmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
	return nil

}

func readPort(port *int) (int, error) {
	// if _, err := fmt.Scanf("%d", port); err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	return err
	// }
	// return nil

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSuffix(input, "\n")
	if input == "" {
		input = "0"
	}

	if *port, err = strconv.Atoi(input); err != nil {
		return *port, err
	}

	return *port, nil
}
