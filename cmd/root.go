package cmd
// taken from github.com/awslabs/fargatecli for now to get started

import (
  "os"
  "runtime"
  
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/garretthoffman/sagecli/console"
  "github.com/spf13/cobra"
  "golang.org/x/crypto/ssh/terminal"
)

const (
	version = "0.0.1"
	
	defaultRegion	= "us-east-1"

	runtimeMacOS = "darwin"
)

var (
	noColor	bool
	noEmoji	bool
	output	ConsoleOutput
	region 	string
	sess 	*session.Session
	verbose	bool
)

var rootCmd = &cobra.Command{
	Use: "sage",
	Short: "Deploy machine learning training jobs, endpoints and notebooks from your command line",
	Long: `Deploy machine learning training jobs, endpoints and notebooks from your command line
	
sage is a command-line interface to manage hosted notebook environments, launch data
labeling jobs, train machine learning models, deploy machine leanring model enpoints using 
AWS Sagemaker.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		output = ConsoleOutput{}

		if cmd.Parent().Name() == "sage" {
			return
		}

		if verbose {
			verbose = true
			console.Verbose = true
			output.Verbose = true
		}

		if terminal.IsTerminal(int(os.Stdout.Fd())) {
			if !noColor {
				console.Color = true
				output.Color = true
			}

			if runtime.GOOS == runtimeMacOS && !noEmoji {
				output.Emoji = true
			}
		}

		envAwsDeafultRegion := os.Getenv("AWS_DEFAULT_REGION")
		envAwsRegion := os.Getenv("AWS_REGION")

		if region == "" {
			if envAwsDeafultRegion != "" {
				region = envAwsDeafultRegion
			} else if envAwsRegion != "" {
				region = envAwsRegion
			} else {
			  	if sess = session.Must(session.NewSession()); *sess.Config.Region != "" {
					region = *sess.Config.Region
				} else {
					region = defaultRegion
				}
			}
		}

		config := &aws.Config{
			Region: aws.String(region),
		}
			
		if verbose {
			config.LogLevel = aws.LogLevel(aws.LogDebugWithHTTPBody)
		}

		sess = session.Must(
			session.NewSession(config),
		)

		_, err := sess.Config.Credentials.Get()

		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NoCredentialProviders":
				console.Issue("Could not find your AWS credentials")
				console.Info("Your AWS credentials could not be found. Please configure your environment with your access key")
				console.Info("   ID and secret access key using either the shared configuration file or environment variables.")
				console.Info("   See http://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials")
				console.Info("   for more details.")
				console.Exit(1)
			default:
				console.ErrorExit(err, "Could not create AWS session")
			}
		}
	},
}

func Execute() {
	rootCmd.Version = version
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().StringVar(&region, "region", "", `AWS region (default "us-east-1")`)
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable color output")

	if runtime.GOOS == runtimeMacOS {
		rootCmd.PersistentFlags().BoolVar(&noEmoji, "no-emoji", false, "Disable emoji output")
	}
}