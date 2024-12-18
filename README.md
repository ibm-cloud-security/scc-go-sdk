[![main](https://github.com/IBM/scc-go-sdk/actions/workflows/main.yaml/badge.svg)](https://github.com/IBM/scc-go-sdk/actions/workflows/main.yaml)
[![test](https://github.com/IBM/scc-go-sdk/actions/workflows/ct-check.yml/badge.svg?branch=main)](https://github.com/IBM/scc-go-sdk/actions/workflows/ct-check.yml)
[![Release](https://img.shields.io/github/v/release/IBM/scc-go-sdk)](https://img.shields.io/github/v/release/IBM/scc-go-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/IBM/scc-go-sdk/v5.svg)](https://pkg.go.dev/github.com/IBM/scc-go-sdk/v5)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/scc-go-sdk?filename=v5%2Fgo.mod)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ad2d585c763ad627e0cb/test_coverage)](https://codeclimate.com/github/IBM/scc-go-sdk/test_coverage)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)


# IBM Cloud Security and Compliance Center Go SDK Version v5.5.1

Go client library to interact with various
[IBM Cloud Security and Compliance Center APIs](https://cloud.ibm.com/apidocs/security-compliance).

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [IBM Cloud Security and Compliance Center Go SDK Version v5.5.1](#ibm-cloud-security--compliance-center-go-sdk-version-v500)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Go modules](#go-modules)
    - [`go get` command](#go-get-command)
  - [Using the SDK](#using-the-sdk)
    - [Example](#example)
  - [Testing and Development](#testing-and-development)
  - [Questions](#questions)
  - [Issues](#issues)
  - [Open source @ IBM](#open-source--ibm)
  - [Contributing](#contributing)
  - [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Security and Compliance Center Go SDK allows developers to programmatically interact with the following IBM Cloud Security and Compliance Center.

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one
[here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: v5.5.1

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `securityandcompliancecenterapiv3` part of the import path is the package name
associated with the Findings service.
See the service table above to find the appropriate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/scc-go-sdk/v5
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
For general SDK usage information, please see
[this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

### Example
```go
package main

import (
  "github.com/IBM/go-sdk-core/v5/core"

  scc "github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)


func main() {
  // Initialize Client
  apiClient, err := scc.NewSecurityAndComplianceCenterApiV3(
    &scc.SecurityAndComplianceCenterApiV3Options{
      URL: "https://{region}.compliance.cloud.ibm.com",
      Authenticator: &core.IamAuthenticator{
        ApiKey: "{ibmcloud_api_key}",
      },
    },
  )
  ...
  // Example: Grab a list of profiles
  lpo := scc.ListProfilesOptions{
    InstanceID: "{instance_id}",
  }
  profiles, res, err := apiClient.ListProfiles(&lpo)
  ...
}
```

## Testing and Development
To test out any changes to SDK locally, `security_and_compliance_center_api_v3.env` should be in the base directory of the current version.

```
SECURITY_AND_COMPLIANCE_CENTER_API_URL=https://us-south.compliance.cloud.ibm.com
SECURITY_AND_COMPLIANCE_CENTER_API_IAM_APIKEY_URL=https://iam.cloud.ibm.com/identity/token
SECURITY_AND_COMPLIANCE_CENTER_API_IAM=<INSERT_SECRET_API_KEY>
SECURITY_AND_COMPLIANCE_CENTER_API_SERVICENAME=SECURITY AND COMPLIANCE CENTER
SECURITY_AND_COMPLIANCE_CENTER_API_ACCOUNTID=<INSERT_IBM_ACCOUNT_ID>
SECURITY_AND_COMPLIANCE_CENTER_API_INSTANCEID=<INSERT_SCC_INSTANCE_ID>
SECURITY_AND_COMPLIANCE_CENTER_API_ATTACHMENTID=<INSERT_SCC_PROFILE_ATTACHMENT_ID>
SECURITY_AND_COMPLIANCE_CENTER_API_PROFILEID=<INSERT_SCC_PROFILE_ID
SECURITY_AND_COMPLIANCE_CENTER_API_AUTHTYPE=noauth
```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/scc-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

The IBM Cloud Security and Compliance Center Go SDK is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
