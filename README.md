# terraform-drift-check

Custom kuberhealthy checker for reporting on Terraform drift.

## Prerequisites

[Kuberhealthy](https://github.com/Comcast/kuberhealthy) should be installed in your Kubernetes cluster. 

## Installation

Install the Terraform drift check via the helm chart in `./charts/terraform-drift-check`

## Configuration

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `terraformHealth.enabled` | Enable Terraform drift health check | `true` |
| `terraformHealth.image.repository` | Image repository | `mellardc/terraform-health` |
| `terraformHealth.runInternal` | The interval that Kuberhealthy will run your check on | `300s` |
| `terraformHealth.timeout` | After this much time, Kuberhealthy will kill your check and consider it "failed" | `5m` |
| `terraformHealth.git.url`  | Url of Git repo containing Terraform config  |  |
| `terraformHealth.git.username` | Git user to authenticate to Git repo  containing Terraform config | |
| `terraformHealth.secretEnv.GIT_TOKEN` | Git token to authenticate to Git repo containing Terraform module | |
| `terraformHealth.secretEnv` | Optional secret environment variables that will be made available as environment variables to Terraform | |
| `terraformHealth.env` | Optional environment settings for Terraform | `[]` |

# Build

Built with Go 1.15
