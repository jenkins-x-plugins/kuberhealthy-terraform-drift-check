package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/Comcast/kuberhealthy/v2/pkg/checks/external/checkclient"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cmdrunner"
	"github.com/jenkins-x/jx-helpers/v3/pkg/gitclient"
	"github.com/jenkins-x/jx-helpers/v3/pkg/gitclient/cli"
	"github.com/jenkins-x/jx-helpers/v3/pkg/gitclient/giturl"
)

const (
	gitProviderURL  = "GIT_URL"
	gitUserName     = "GIT_USER"
	gitToken        = "GIT_TOKEN"
	terraformBinary = "terraform"
)

type Options struct {
	gitClient gitclient.Interface
	gitToken  string
	gitUser   string
	gitUrl    string
}

func newOptions() *Options {
	gitToken := os.Getenv(gitToken)
	gitUrl := os.Getenv(gitProviderURL)
	gitUser := os.Getenv(gitUserName)
	return &Options{
		gitClient: cli.NewCLIClient("", nil),
		gitToken:  gitToken,
		gitUrl:    gitUrl,
		gitUser:   gitUser,
	}
}

func main() {

	o := newOptions()
	if khErrors := o.findErrors(); khErrors != nil {
		_ = checkclient.ReportFailure(khErrors)
	}
	_ = checkclient.ReportSuccess()
}

func getOSEnvVars() map[string]string {
	envMap := map[string]string{}

	for _, entry := range os.Environ() {
		pair := strings.SplitN(entry, "=", 2)
		envMap[pair[0]] = pair[1]
	}
	return envMap
}

func (o Options) findErrors() []string {

	gitInfo, err := giturl.ParseGitURL(o.gitUrl)
	if err != nil {
		return []string{fmt.Sprintf("error parsing git url %v", err)}
	}

	gitUrl, _ := url.Parse(o.gitUrl)
	gitUrl.User = url.UserPassword(o.gitUser, o.gitToken)
	parentDir := "/tmp"
	gitDir := path.Join(parentDir, gitInfo.Name)

	_, err = o.gitClient.Command(parentDir, "clone", gitUrl.String())

	if err != nil {
		return []string{fmt.Sprintf("error cloning terraform git repo %v", err)}
	}

	err = terraformInit(gitDir)

	if err != nil {
		return []string{fmt.Sprintf("error initialising terraform module %v", err)}
	}

	plan, err := terraformPlan(gitDir, getOSEnvVars())

	if err != nil {
		return []string{fmt.Sprintf("terraform plan produced error or diff %v - resultant plan = %s", err, plan)}
	}
	return nil
}

func terraformInit(dir string) error {

	command := cmdrunner.NewCommand(dir, terraformBinary, "init")
	_, err := cmdrunner.DefaultCommandRunner(command)
	if err != nil {
		return fmt.Errorf("error initialising terraform module %w", err)
	}
	return nil
}

func terraformPlan(dir string, envVars map[string]string) (string, error) {
	command := cmdrunner.NewCommand(dir, terraformBinary, "plan", "-detailed-exitcode")
	command.Env = envVars
	plan, err := cmdrunner.DefaultCommandRunner(command)
	if err != nil {
		return "", fmt.Errorf("error running plan or diff detected %w", err)
	}
	return plan, nil
}
