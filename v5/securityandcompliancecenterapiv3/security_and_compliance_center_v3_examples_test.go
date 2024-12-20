//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package securityandcompliancecenterapiv3_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	securityandcompliancecenterv3 "github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Security and Compliance Center service.
//
// The following configuration properties are assumed to be defined:
// SECURITY_AND_COMPLIANCE_CENTER_URL=<service base url>
// SECURITY_AND_COMPLIANCE_CENTER_AUTH_TYPE=iam
// SECURITY_AND_COMPLIANCE_CENTER_APIKEY=<IAM apikey>
// SECURITY_AND_COMPLIANCE_CENTER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`SecurityAndComplianceCenterV3 Examples Tests`, func() {
	const externalConfigFile = "../security_and_compliance_center_v3.env"

	var (
		securityAndComplianceCenterService *securityandcompliancecenterv3.SecurityAndComplianceCenterApiV3
		config                             map[string]string
		serviceURL                         string

		// Variables to hold link values
		accountIDForReportLink                     string
		attachmentIDForReportLink                  string
		attachmentIDLink                           string
		controlLibraryIDLink                       string
		eTagLink                                   string
		eventNotificationsCRNForUpdateSettingsLink string
		groupIDForReportLink                       string
		instanceIDForLink                          string
		objectStorageBucketForUpdateSettingsLink   string
		objectStorageCRNForUpdateSettingsLink      string
		objectStorageLocationForUpdateSettingsLink string
		oldProfileIDForReportLink                  string
		profileIDForReportLink                     string
		profileIDLink                              string
		providerTypeIDLink                         string
		providerTypeInstanceIDLink                 string
		reportIDForReportLink                      string
		ruleIDLink                                 string
		scanIDforReportLink                        string
		scopeIDLink                                string
		scopeIDforReportLink                       string
		subScopeIDLink                             string
		targetIDLink                               string
		typeForReportLink                          string
		workloadProtectionCRNLink                  string
	)

	shouldSkipTest := func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(securityandcompliancecenterv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}
			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}

			// Manual: Adding the ability to customize identifiers using environmental variables
			accountIDForReportLink = config["ACCOUNTID"]
			if accountIDForReportLink == "" {
				Skip("Unable to load accountIDForReportLink configuration property, skipping tests")
			}

			scopeIDforReportLink = config["SCOPEID"]
			if scopeIDforReportLink == "" {
				Skip("Unable to load scopeIDforReportLink configuration property, skipping tests")
			}

			scanIDforReportLink = config["SCANJOBID"]
			if scanIDforReportLink == "" {
				Skip("Unable to load scanreportID configuration property, skipping tests")
			}

			reportIDForReportLink = config["REPORTID"]
			if reportIDForReportLink == "" {
				Skip("Unable to load reportID configuration property, skipping tests")
			}

			instanceIDForLink = config["INSTANCEID"]
			if instanceIDForLink == "" {
				Skip("Unable to load instanceID configuration property, skipping tests")
			}

			profileIDForReportLink = config["PROFILEID"]
			if profileIDForReportLink == "" {
				Skip("Unable to load profileID configuration property, skipping tests")
			}

			oldProfileIDForReportLink = config["OLDPROFILEID"]
			if oldProfileIDForReportLink == "" {
				Skip("Unable to load oldprofileID configuration property, skipping tests")
			}

			attachmentIDForReportLink = config["ATTACHMENTID"]
			if attachmentIDForReportLink == "" {
				Skip("Unable to load attachmentID configuration property, skipping tests")
			} else {
				attachmentIDLink = attachmentIDForReportLink
			}

			workloadProtectionCRNLink = config["WORKLOADPROTECTIONCRN"]
			if workloadProtectionCRNLink == "" {
				Skip("Unable to load workloadProtectCRNLInk configuration property, skipping tests")
			}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			securityAndComplianceCenterServiceOptions := &securityandcompliancecenterv3.SecurityAndComplianceCenterApiV3Options{}

			securityAndComplianceCenterService, err = securityandcompliancecenterv3.NewSecurityAndComplianceCenterV3(securityAndComplianceCenterServiceOptions)
			if err != nil {
				panic(err)
			}

			// end-common

			Expect(securityAndComplianceCenterService).ToNot(BeNil())
		})
	})

	Describe(`SecurityAndComplianceCenterV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := securityAndComplianceCenterService.NewGetSettingsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			settings, response, err := securityAndComplianceCenterService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())

			eventNotificationsCRNForUpdateSettingsLink = *settings.EventNotifications.InstanceCRN
			fmt.Fprintf(GinkgoWriter, "Saved eventNotificationsCRNForUpdateSettingsLink value: %v\n", eventNotificationsCRNForUpdateSettingsLink)
			objectStorageCRNForUpdateSettingsLink = *settings.ObjectStorage.InstanceCRN
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageCRNForUpdateSettingsLink value: %v\n", objectStorageCRNForUpdateSettingsLink)
			objectStorageBucketForUpdateSettingsLink = *settings.ObjectStorage.Bucket
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageBucketForUpdateSettingsLink value: %v\n", objectStorageBucketForUpdateSettingsLink)
			objectStorageLocationForUpdateSettingsLink = *settings.ObjectStorage.BucketLocation
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageLocationForUpdateSettingsLink value: %v\n", objectStorageLocationForUpdateSettingsLink)
		})
		It(`CreateCustomControlLibrary request example`, func() {
			fmt.Println("\nCreateCustomControlLibrary() result:")
			// begin-create_custom_control_library

			assessmentPrototypeModel := &securityandcompliancecenterv3.AssessmentPrototype{
				AssessmentID:          core.StringPtr("rule-d1bd9f3f-bee1-46c5-9533-da8bba9eed4e"),
				AssessmentDescription: core.StringPtr("This rule will check on regulation"),
			}

			controlSpecificationPrototypeModel := &securityandcompliancecenterv3.ControlSpecificationPrototype{
				ComponentID: core.StringPtr("apprapp"),
				Environment: core.StringPtr("ibm-cloud"),
				Description: core.StringPtr("This field is used to describe a control specification"),
				Assessments: []securityandcompliancecenterv3.AssessmentPrototype{*assessmentPrototypeModel},
			}

			controlDocModel := &securityandcompliancecenterv3.ControlDoc{}

			controlPrototypeModel := &securityandcompliancecenterv3.ControlPrototype{
				ControlName:           core.StringPtr("security"),
				ControlDescription:    core.StringPtr("This is a description of a control"),
				ControlCategory:       core.StringPtr("test-control"),
				ControlRequirement:    core.BoolPtr(true),
				ControlSpecifications: []securityandcompliancecenterv3.ControlSpecificationPrototype{*controlSpecificationPrototypeModel},
				ControlDocs:           controlDocModel,
				Status:                core.StringPtr("disabled"),
			}

			createCustomControlLibraryOptions := securityAndComplianceCenterService.NewCreateCustomControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"custom control library from SDK",
				"This is a custom control library made from the SDK test framework",
				"custom",
				"0.0.1",
				[]securityandcompliancecenterv3.ControlPrototype{*controlPrototypeModel},
			)

			controlLibrary, response, err := securityAndComplianceCenterService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-create_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(controlLibrary).ToNot(BeNil())

			controlLibraryIDLink = *controlLibrary.ID
			fmt.Fprintf(GinkgoWriter, "Saved controlLibraryIDLink value: %v\n", controlLibraryIDLink)
		})
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile

			profileControlsPrototypeModel := &securityandcompliancecenterv3.ProfileControlsPrototype{
				ControlLibraryID: core.StringPtr("51ca566e-c559-412b-8d64-f05b57044c32"),
				ControlID:        core.StringPtr("2ce21ba3-0548-49a3-88e2-1122632218f4"),
			}

			defaultParametersPrototypeModel := &securityandcompliancecenterv3.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr(`["1.2","1.3"]`),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			createProfileOptions := securityAndComplianceCenterService.NewCreateProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			createProfileOptions.SetProfileName("Profile Example")
			createProfileOptions.SetProfileDescription("This is a profile")
			createProfileOptions.SetProfileVersion("1.0.0")
			createProfileOptions.SetLatest(true)
			createProfileOptions.SetControls([]securityandcompliancecenterv3.ProfileControlsPrototype{*profileControlsPrototypeModel})
			createProfileOptions.SetDefaultParameters([]securityandcompliancecenterv3.DefaultParametersPrototype{*defaultParametersPrototypeModel})

			profile, response, err := securityAndComplianceCenterService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())

			profileIDLink = *profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIDLink value: %v\n", profileIDLink)
		})
		It(`CreateProfileAttachment request example`, func() {
			fmt.Println("\nCreateProfileAttachment() result:")
			// begin-create_profile_attachment

			parameterModel1 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:        core.StringPtr("tls_version"),
				ParameterDisplayName: core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:        core.StringPtr("string_list"),
				ParameterValue:       core.StringPtr("['1.2', '1.3']"),
			}
			parameterModel2 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"),
				ParameterName:        core.StringPtr("ssh_port"),
				ParameterDisplayName: core.StringPtr("Network ACL rule for allowed IPs to SSH port"),
				ParameterType:        core.StringPtr("numeric"),
				ParameterValue:       core.StringPtr("22"),
			}
			parameterModel3 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
				ParameterName:        core.StringPtr("rdp_port"),
				ParameterValue:       core.StringPtr("3389"),
				ParameterDisplayName: core.StringPtr("Security group rule RDP allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel4 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
				ParameterName:        core.StringPtr("ssh_port"),
				ParameterValue:       core.StringPtr("22"),
				ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel5 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
				ParameterName:        core.StringPtr("rdp_port"),
				ParameterValue:       core.StringPtr("3389"),
				ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel6 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-96527f89-1867-4581-b923-1400e04661e0"),
				ParameterName:        core.StringPtr("exclude_default_security_groups"),
				ParameterValue:       core.StringPtr("['Update the parameter']"),
				ParameterDisplayName: core.StringPtr("Exclude the default security groups"),
				ParameterType:        core.StringPtr("string_list"),
			}

			attachmentParameters := []securityandcompliancecenterv3.Parameter{
				*parameterModel1,
				*parameterModel2,
				*parameterModel3,
				*parameterModel4,
				*parameterModel5,
				*parameterModel6,
			}

			attachmentNotificationsControlsModel := &securityandcompliancecenterv3.AttachmentNotificationsControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentNotificationsModel := &securityandcompliancecenterv3.AttachmentNotifications{
				Enabled:  core.BoolPtr(true),
				Controls: attachmentNotificationsControlsModel,
			}

			multiCloudScopePayloadModel := &securityandcompliancecenterv3.MultiCloudScopePayload{
				ID: core.StringPtr("8baad3b5-2e69-4027-9967-efac19508e1c"),
			}

			profileAttachmentBaseModel := &securityandcompliancecenterv3.ProfileAttachmentBase{
				AttachmentParameters: attachmentParameters,
				Description:          core.StringPtr("This is a profile attachment targeting IBM CIS Foundation using a SDK"),
				Name:                 core.StringPtr("Profile Attachment for IBM CIS Foundation SDK test"),
				Notifications:        attachmentNotificationsModel,
				Schedule:             core.StringPtr("daily"),
				Scope:                []securityandcompliancecenterv3.MultiCloudScopePayload{*multiCloudScopePayloadModel},
				Status:               core.StringPtr("disabled"),
			}

			createProfileAttachmentOptions := securityAndComplianceCenterService.NewCreateProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
			)
			createProfileAttachmentOptions.SetAttachments([]securityandcompliancecenterv3.ProfileAttachmentBase{*profileAttachmentBaseModel})

			profileAttachmentResponse, response, err := securityAndComplianceCenterService.CreateProfileAttachment(createProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profileAttachmentResponse).ToNot(BeNil())

			attachmentIDLink = *profileAttachmentResponse.Attachments[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIDLink value: %v\n", attachmentIDLink)
		})
		It(`CreateScope request example`, func() {
			fmt.Println("\nCreateScope() result:")
			// begin-create_scope

			scopeIDPropertyModel := &securityandcompliancecenterv3.ScopePropertyScopeID{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("ff88f007f9ff4622aac4fbc0eda36255"),
			}

			scopeTypePropertyModel := &securityandcompliancecenterv3.ScopePropertyScopeType{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			createScopeOptions := securityAndComplianceCenterService.NewCreateScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			createScopeOptions.SetName("Sample Scope")
			createScopeOptions.SetDescription("The scope that is defined for IBM resources.")
			createScopeOptions.SetEnvironment("ibm-cloud")
			createScopeOptions.SetProperties([]securityandcompliancecenterv3.ScopePropertyIntf{scopeIDPropertyModel, scopeTypePropertyModel})

			scope, response, err := securityAndComplianceCenterService.CreateScope(createScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-create_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())

			scopeIDLink = *scope.ID
			fmt.Fprintf(GinkgoWriter, "Saved scopeIDLink value: %v\n", scopeIDLink)
		})
		It(`CreateSubscope request example`, func() {
			fmt.Println("\nCreateSubscope() result:")
			// begin-create_subscope

			scopeIDProperty := &securityandcompliancecenterv3.ScopePropertyScopeID{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("1f689f08ec9b47b885c2659c17029581"),
			}

			scopeTypeProperty := &securityandcompliancecenterv3.ScopePropertyScopeType{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account.resource_group"),
			}
			scopePrototypeModel := &securityandcompliancecenterv3.ScopePrototype{
				Name:        core.StringPtr("ibm subscope"),
				Description: core.StringPtr("The subscope that is defined for IBM resources."),
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterv3.ScopePropertyIntf{scopeIDProperty, scopeTypeProperty},
			}

			createSubscopeOptions := securityAndComplianceCenterService.NewCreateSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)
			createSubscopeOptions.SetSubscopes([]securityandcompliancecenterv3.ScopePrototype{*scopePrototypeModel})

			subScopeResponse, response, err := securityAndComplianceCenterService.CreateSubscope(createSubscopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subScopeResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subScopeResponse).ToNot(BeNil())

			subScopeIDLink = *subScopeResponse.Subscopes[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved subScopeIDLink value: %v\n", subScopeIDLink)
		})
		It(`GetLatestReports request example`, func() {
			fmt.Println("\nGetLatestReports() result:")
			// begin-get_latest_reports

			getLatestReportsOptions := securityAndComplianceCenterService.NewGetLatestReportsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			getLatestReportsOptions.SetSort("profile_name")

			reportLatest, response, err := securityAndComplianceCenterService.GetLatestReports(getLatestReportsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportLatest, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_reports

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportLatest).ToNot(BeNil())

			accountIDForReportLink = *reportLatest.Reports[0].Account.ID
			fmt.Fprintf(GinkgoWriter, "Saved accountIDForReportLink value: %v\n", accountIDForReportLink)
			attachmentIDForReportLink = *reportLatest.Reports[0].Attachment.ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIDForReportLink value: %v\n", attachmentIDForReportLink)
			groupIDForReportLink = *reportLatest.Reports[0].GroupID
			fmt.Fprintf(GinkgoWriter, "Saved groupIDForReportLink value: %v\n", groupIDForReportLink)
			profileIDForReportLink = *reportLatest.Reports[0].Profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIDForReportLink value: %v\n", profileIDForReportLink)
			typeForReportLink = *reportLatest.Reports[0].Type
			fmt.Fprintf(GinkgoWriter, "Saved typeForReportLink value: %v\n", typeForReportLink)
		})
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

			additionalTargetAttributeModel := &securityandcompliancecenterv3.AdditionalTargetAttribute{
				Name:     core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-east"),
			}

			ruleTargetPrototypeModel := &securityandcompliancecenterv3.RuleTargetPrototype{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			conditionItemModel := &securityandcompliancecenterv3.ConditionItemConditionBase{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &securityandcompliancecenterv3.RequiredConfigConditionListConditionListConditionAnd{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []securityandcompliancecenterv3.ConditionItemIntf{conditionItemModel},
			}

			ruleParameterModel := &securityandcompliancecenterv3.RuleParameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &securityandcompliancecenterv3.Import{
				Parameters: []securityandcompliancecenterv3.RuleParameter{*ruleParameterModel},
			}

			createRuleOptions := securityAndComplianceCenterService.NewCreateRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"Example rule",
				ruleTargetPrototypeModel,
				requiredConfigModel,
			)
			createRuleOptions.SetVersion("1.0.0")
			createRuleOptions.SetImport(importModel)
			createRuleOptions.SetLabels([]string{})

			rule, response, err := securityAndComplianceCenterService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			ruleIDLink = *rule.ID
			fmt.Fprintf(GinkgoWriter, "Saved ruleIDLink value: %v\n", ruleIDLink)
		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := securityAndComplianceCenterService.NewGetRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				ruleIDLink,
			)

			rule, response, err := securityAndComplianceCenterService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
		It(`ReplaceRule request example`, func() {
			fmt.Println("\nReplaceRule() result:")
			// begin-replace_rule

			additionalTargetAttributeModel := &securityandcompliancecenterv3.AdditionalTargetAttribute{
				Name:     core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-south"),
			}

			ruleTargetPrototypeModel := &securityandcompliancecenterv3.RuleTargetPrototype{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			conditionItemModel := &securityandcompliancecenterv3.ConditionItemConditionBase{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &securityandcompliancecenterv3.RequiredConfigConditionListConditionListConditionAnd{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []securityandcompliancecenterv3.ConditionItemIntf{conditionItemModel},
			}

			ruleParameterModel := &securityandcompliancecenterv3.RuleParameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &securityandcompliancecenterv3.Import{
				Parameters: []securityandcompliancecenterv3.RuleParameter{*ruleParameterModel},
			}

			replaceRuleOptions := securityAndComplianceCenterService.NewReplaceRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				ruleIDLink,
				eTagLink,
				"Example rule from example SDK run",
				ruleTargetPrototypeModel,
				requiredConfigModel,
			)
			replaceRuleOptions.SetVersion("1.0.1")
			replaceRuleOptions.SetImport(importModel)
			replaceRuleOptions.SetLabels([]string{})

			rule, response, err := securityAndComplianceCenterService.ReplaceRule(replaceRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-replace_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
		It(`UpdateSettings request example`, func() {
			fmt.Println("\nUpdateSettings() result:")
			// begin-update_settings

			objectStoragePrototypeModel := &securityandcompliancecenterv3.ObjectStoragePrototype{
				Bucket:      core.StringPtr("px-scan-results"),
				InstanceCRN: core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::"),
			}

			eventNotificationsPrototypeModel := &securityandcompliancecenterv3.EventNotificationsPrototype{
				InstanceCRN:       core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"),
				SourceDescription: core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center."),
				SourceName:        core.StringPtr("scc-sdk-integration"),
			}

			updateSettingsOptions := securityAndComplianceCenterService.NewUpdateSettingsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			updateSettingsOptions.SetObjectStorage(objectStoragePrototypeModel)
			updateSettingsOptions.SetEventNotifications(eventNotificationsPrototypeModel)

			settings, response, err := securityAndComplianceCenterService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(200), Equal(204)))

			switch statusCode := response.StatusCode; statusCode {
			case 200:
				Expect(settings).ToNot(BeNil())
			case 204:
				Expect(settings).To(BeNil())
			}
		})
		It(`PostTestEvent request example`, func() {
			fmt.Println("\nPostTestEvent() result:")
			// begin-post_test_event

			postTestEventOptions := securityAndComplianceCenterService.NewPostTestEventOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			testEvent, response, err := securityAndComplianceCenterService.PostTestEvent(postTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})
		It(`ListInstanceAttachments request example`, func() {
			fmt.Println("\nListInstanceAttachments() result:")
			// begin-list_instance_attachments
			listInstanceAttachmentsOptions := &securityandcompliancecenterv3.ListInstanceAttachmentsOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:         &accountIDForReportLink,
				VersionGroupLabel: core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d"),
				Limit:             core.Int64Ptr(int64(10)),
				Sort:              core.StringPtr("created_on"),
				Direction:         core.StringPtr("desc"),
			}

			pager, err := securityAndComplianceCenterService.NewInstanceAttachmentsPager(listInstanceAttachmentsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.ProfileAttachment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_instance_attachments
		})
		It(`ListProfileAttachments request example`, func() {
			fmt.Println("\nListProfileAttachments() result:")
			// begin-list_profile_attachments

			listProfileAttachmentsOptions := securityAndComplianceCenterService.NewListProfileAttachmentsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
			)

			profileAttachmentCollection, response, err := securityAndComplianceCenterService.ListProfileAttachments(listProfileAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_profile_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachmentCollection).ToNot(BeNil())
		})
		It(`ReplaceProfileAttachment request example`, func() {
			fmt.Println("\nReplaceProfileAttachment() result:")
			// begin-replace_profile_attachment

			parameterModel1 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:        core.StringPtr("tls_version"),
				ParameterDisplayName: core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:        core.StringPtr("string_list"),
				ParameterValue:       core.StringPtr("['1.2', '1.3']"),
			}
			parameterModel2 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"),
				ParameterName:        core.StringPtr("ssh_port"),
				ParameterDisplayName: core.StringPtr("Network ACL rule for allowed IPs to SSH port"),
				ParameterType:        core.StringPtr("numeric"),
				ParameterValue:       core.StringPtr("23"),
			}
			parameterModel3 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
				ParameterName:        core.StringPtr("rdp_port"),
				ParameterValue:       core.StringPtr("3390"),
				ParameterDisplayName: core.StringPtr("Security group rule RDP allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel4 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
				ParameterName:        core.StringPtr("ssh_port"),
				ParameterValue:       core.StringPtr("23"),
				ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel5 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
				ParameterName:        core.StringPtr("rdp_port"),
				ParameterValue:       core.StringPtr("3340"),
				ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel6 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-96527f89-1867-4581-b923-1400e04661e0"),
				ParameterName:        core.StringPtr("exclude_default_security_groups"),
				ParameterValue:       core.StringPtr("['Update the parameter']"),
				ParameterDisplayName: core.StringPtr("Exclude the default security groups"),
				ParameterType:        core.StringPtr("string_list"),
			}

			attachmentParameters := []securityandcompliancecenterv3.Parameter{
				*parameterModel1,
				*parameterModel2,
				*parameterModel3,
				*parameterModel4,
				*parameterModel5,
				*parameterModel6,
			}

			AttachmentNotifications := &securityandcompliancecenterv3.AttachmentNotifications{
				Enabled: core.BoolPtr(false),
				Controls: &securityandcompliancecenterv3.AttachmentNotificationsControls{
					ThresholdLimit:   core.Int64Ptr(int64(15)),
					FailedControlIds: []string{},
				},
			}

			multiCloudScopePayloadModel := []securityandcompliancecenterv3.MultiCloudScopePayload{
				{
					ID: core.StringPtr("a1c2a74e-508d-48b3-8e46-5ab4f4795a7b"),
				},
			}

			replaceProfileAttachmentOptions := securityAndComplianceCenterService.NewReplaceProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
				attachmentIDLink,
				"This profile attachment can be deleted",
				"Profile Attachment SDK replacement",
				"daily",
				"disabled",
				AttachmentNotifications,
				multiCloudScopePayloadModel,
				attachmentParameters,
			)
			profileAttachment, response, err := securityAndComplianceCenterService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`GetProfileAttachment request example`, func() {
			fmt.Println("\nGetProfileAttachment() result:")
			// begin-get_profile_attachment

			getProfileAttachmentOptions := securityAndComplianceCenterService.NewGetProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
				attachmentIDLink,
			)

			profileAttachment, response, err := securityAndComplianceCenterService.GetProfileAttachment(getProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`UpgradeAttachment request example`, func() {
			fmt.Println("\nUpgradeAttachment() result:")
			// begin-upgrade_attachment

			parameterModel1 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:        core.StringPtr("tls_version"),
				ParameterDisplayName: core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:        core.StringPtr("string_list"),
				ParameterValue:       core.StringPtr("['1.2', '1.3']"),
			}
			parameterModel2 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"),
				ParameterName:        core.StringPtr("ssh_port"),
				ParameterDisplayName: core.StringPtr("Network ACL rule for allowed IPs to SSH port"),
				ParameterType:        core.StringPtr("numeric"),
				ParameterValue:       core.StringPtr("23"),
			}
			parameterModel3 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
				ParameterName:        core.StringPtr("rdp_port"),
				ParameterValue:       core.StringPtr("3390"),
				ParameterDisplayName: core.StringPtr("Security group rule RDP allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel4 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
				ParameterName:        core.StringPtr("ssh_port"),
				ParameterValue:       core.StringPtr("23"),
				ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel5 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
				ParameterName:        core.StringPtr("rdp_port"),
				ParameterValue:       core.StringPtr("3340"),
				ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
				ParameterType:        core.StringPtr("numeric"),
			}
			parameterModel6 := &securityandcompliancecenterv3.Parameter{
				AssessmentType:       core.StringPtr("automated"),
				AssessmentID:         core.StringPtr("rule-96527f89-1867-4581-b923-1400e04661e0"),
				ParameterName:        core.StringPtr("exclude_default_security_groups"),
				ParameterValue:       core.StringPtr("['Viewer']"),
				ParameterDisplayName: core.StringPtr("Exclude the default security groups"),
				ParameterType:        core.StringPtr("string_list"),
			}
			attachmentParameters := []securityandcompliancecenterv3.Parameter{
				*parameterModel1,
				*parameterModel2,
				*parameterModel3,
				*parameterModel4,
				*parameterModel5,
				*parameterModel6,
			}

			upgradeAttachmentOptions := securityAndComplianceCenterService.NewUpgradeAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
				attachmentIDLink,
			)
			upgradeAttachmentOptions.SetAttachmentParameters(attachmentParameters)

			profileAttachment, response, err := securityAndComplianceCenterService.UpgradeAttachment(upgradeAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-upgrade_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`ListControlLibraries request example`, func() {
			fmt.Println("\nListControlLibraries() result:")
			// begin-list_control_libraries

			listControlLibrariesOptions := securityAndComplianceCenterService.NewListControlLibrariesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			controlLibraryCollection, response, err := securityAndComplianceCenterService.ListControlLibraries(listControlLibrariesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_control_libraries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryCollection).ToNot(BeNil())
		})
		It(`ReplaceCustomControlLibrary request example`, func() {
			fmt.Println("\nReplaceCustomControlLibrary() result:")
			// begin-replace_custom_control_library

			assessmentPrototypeModel := &securityandcompliancecenterv3.Assessment{
				// Manual Update
				AssessmentID:          core.StringPtr("rule-5bdfc82b-5eed-4405-b116-fa76292ec003"),
				AssessmentDescription: core.StringPtr("Ensure that IAM IDentity has cbr enabled"),
			}

			controlSpecificationPrototypeModel := &securityandcompliancecenterv3.ControlSpecification{
				ComponentID: core.StringPtr("iam-identity"),
				Environment: core.StringPtr("ibm-cloud"),
				Description: core.StringPtr("CBR security policies"),
				Assessments: []securityandcompliancecenterv3.Assessment{*assessmentPrototypeModel},
			}

			controlDocModel := &securityandcompliancecenterv3.ControlDoc{}

			controlPrototypeModel := &securityandcompliancecenterv3.Control{
				ControlName:           core.StringPtr("security"),
				ControlDescription:    core.StringPtr("Check whether IAM Identity has context-based restrictions enabled"),
				ControlCategory:       core.StringPtr("cbr"),
				ControlRequirement:    core.BoolPtr(true),
				ControlParent:         core.StringPtr(""),
				ControlSpecifications: []securityandcompliancecenterv3.ControlSpecification{*controlSpecificationPrototypeModel},
				ControlDocs:           controlDocModel,
				Status:                core.StringPtr("enabled"),
			}

			replaceCustomControlLibraryOptions := securityAndComplianceCenterService.NewReplaceCustomControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				controlLibraryIDLink,
				"testString",
				"testString",
				"custom",
				"0.0.1",
				[]securityandcompliancecenterv3.Control{*controlPrototypeModel},
			)

			controlLibrary, response, err := securityAndComplianceCenterService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-replace_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`GetControlLibrary request example`, func() {
			fmt.Println("\nGetControlLibrary() result:")
			// begin-get_control_library

			getControlLibraryOptions := securityAndComplianceCenterService.NewGetControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				controlLibraryIDLink,
			)

			controlLibrary, response, err := securityAndComplianceCenterService.GetControlLibrary(getControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-get_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := securityAndComplianceCenterService.NewListProfilesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			profileCollection, response, err := securityAndComplianceCenterService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileCollection).ToNot(BeNil())
		})
		It(`ReplaceProfile request example`, func() {
			fmt.Println("\nReplaceProfile() result:")
			// begin-replace_profile

			replaceProfileOptions := securityAndComplianceCenterService.NewReplaceProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
				"SDK example test run",
				"This profile can be deleted",
			)

			profile, response, err := securityAndComplianceCenterService.ReplaceProfile(replaceProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := securityAndComplianceCenterService.NewGetProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profile, response, err := securityAndComplianceCenterService.GetProfile(getProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-get_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`ReplaceProfileParameters request example`, func() {
			fmt.Println("\nReplaceProfileParameters() result:")
			// begin-replace_profile_parameters

			replaceProfileParametersOptions := securityAndComplianceCenterService.NewReplaceProfileParametersOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterService.ReplaceProfileParameters(replaceProfileParametersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileDefaultParametersResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_parameters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
		It(`ListProfileParameters request example`, func() {
			fmt.Println("\nListProfileParameters() result:")
			// begin-list_profile_parameters

			listProfileParametersOptions := securityAndComplianceCenterService.NewListProfileParametersOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterService.ListProfileParameters(listProfileParametersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileDefaultParametersResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_profile_parameters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
		It(`CompareProfiles request example`, func() {
			fmt.Println("\nCompareProfiles() result:")
			// begin-compare_profiles

			compareProfilesOptions := securityAndComplianceCenterService.NewCompareProfilesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
			)

			comparePredefinedProfilesResponse, response, err := securityAndComplianceCenterService.CompareProfiles(compareProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(comparePredefinedProfilesResponse, "", "  ")
			fmt.Println(string(b))

			// end-compare_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(comparePredefinedProfilesResponse).ToNot(BeNil())
		})
		It(`ListScopes request example`, func() {
			fmt.Println("\nListScopes() result:")
			// begin-list_scopes
			listScopesOptions := &securityandcompliancecenterv3.ListScopesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:      core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterService.NewScopesPager(listScopesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.Scope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_scopes
		})
		It(`UpdateScope request example`, func() {
			fmt.Println("\nUpdateScope() result:")
			// begin-update_scope

			updateScopeOptions := securityAndComplianceCenterService.NewUpdateScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)
			updateScopeOptions.SetName("updated name of scope")
			updateScopeOptions.SetDescription("updated scope description")

			scope, response, err := securityAndComplianceCenterService.UpdateScope(updateScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-update_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
		It(`GetScope request example`, func() {
			fmt.Println("\nGetScope() result:")
			// begin-get_scope

			getScopeOptions := securityAndComplianceCenterService.NewGetScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)

			scope, response, err := securityAndComplianceCenterService.GetScope(getScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-get_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
		It(`ListSubscopes request example`, func() {
			fmt.Println("\nListSubscopes() result:")
			// begin-list_subscopes
			listSubscopesOptions := &securityandcompliancecenterv3.ListSubscopesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    &scopeIDLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterService.NewSubscopesPager(listSubscopesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.SubScope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_subscopes
		})
		It(`GetSubscope request example`, func() {
			fmt.Println("\nGetSubscope() result:")
			// begin-get_subscope

			getSubscopeOptions := securityAndComplianceCenterService.NewGetSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				subScopeIDLink,
			)

			subScope, response, err := securityAndComplianceCenterService.GetSubscope(getSubscopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subScope, "", "  ")
			fmt.Println(string(b))

			// end-get_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
		It(`UpdateSubscope request example`, func() {
			fmt.Println("\nUpdateSubscope() result:")
			// begin-update_subscope

			updateSubscopeOptions := securityAndComplianceCenterService.NewUpdateSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				subScopeIDLink,
			)
			updateSubscopeOptions.SetName("SDK updated name of scope")
			updateSubscopeOptions.SetDescription("updated scope description")

			subScope, response, err := securityAndComplianceCenterService.UpdateSubscope(updateSubscopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subScope, "", "  ")
			fmt.Println(string(b))

			// end-update_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
		It(`CreateTarget request example`, func() {
			fmt.Println("\nCreateTarget() result:")
			// begin-create_target

			createTargetOptions := securityAndComplianceCenterService.NewCreateTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"62ecf99b240144dea9125666249edfcb",
				"Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54",
				"Sample Target from SDK run",
			)

			target, response, err := securityAndComplianceCenterService.CreateTarget(createTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-create_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
		})
		It(`ListTargets request example`, func() {
			fmt.Println("\nListTargets() result:")
			// begin-list_targets

			listTargetsOptions := securityAndComplianceCenterService.NewListTargetsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			targetCollection, response, err := securityAndComplianceCenterService.ListTargets(listTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})
		It(`GetTarget request example`, func() {
			fmt.Println("\nGetTarget() result:")
			// begin-get_target

			getTargetOptions := securityAndComplianceCenterService.NewGetTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				targetIDLink,
			)

			target, response, err := securityAndComplianceCenterService.GetTarget(getTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-get_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
		It(`ReplaceTarget request example`, func() {
			fmt.Println("\nReplaceTarget() result:")
			// begin-replace_target

			replaceTargetOptions := securityAndComplianceCenterService.NewReplaceTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				targetIDLink,
				"62ecf99b240144dea9125666249edfcb",
				"Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54",
				"Sample Target from SDK run",
			)

			target, response, err := securityAndComplianceCenterService.ReplaceTarget(replaceTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-replace_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
		It(`ListProviderTypes request example`, func() {
			fmt.Println("\nListProviderTypes() result:")
			// begin-list_provider_types

			listProviderTypesOptions := securityAndComplianceCenterService.NewListProviderTypesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			providerTypeCollection, response, err := securityAndComplianceCenterService.ListProviderTypes(listProviderTypesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_provider_types

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeCollection).ToNot(BeNil())

			for _, providerType := range providerTypeCollection.ProviderTypes {
				if *providerType.Name == "Caveonix" {
					providerTypeIDLink = *providerType.ID
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeIDLink value: %v\n", providerTypeIDLink)
		})
		It(`CreateProviderTypeInstance request example`, func() {
			fmt.Println("\nCreateProviderTypeInstance() result:")
			// begin-create_provider_type_instance

			createProviderTypeInstanceOptions := securityAndComplianceCenterService.NewCreateProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				providerTypeIDLink,
				"providerTypeFromSDK run",
			)

			providerTypeInstance, response, err := securityAndComplianceCenterService.CreateProviderTypeInstance(createProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstance, "", "  ")
			fmt.Println(string(b))

			// end-create_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(providerTypeInstance).ToNot(BeNil())

			providerTypeInstanceIDLink = *providerTypeInstance.ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeInstanceIDLink value: %v\n", providerTypeInstanceIDLink)
		})
		It(`ListProviderTypeInstances request example`, func() {
			fmt.Println("\nListProviderTypeInstances() result:")
			// begin-list_provider_type_instances

			listProviderTypeInstancesOptions := securityAndComplianceCenterService.NewListProviderTypeInstancesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				providerTypeIDLink,
			)

			providerTypeInstanceCollection, response, err := securityAndComplianceCenterService.ListProviderTypeInstances(listProviderTypeInstancesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstanceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_provider_type_instances

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceCollection).ToNot(BeNil())
		})
		It(`GetProviderTypeInstance request example`, func() {
			fmt.Println("\nGetProviderTypeInstance() result:")
			// begin-get_provider_type_instance

			getProviderTypeInstanceOptions := securityAndComplianceCenterService.NewGetProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				providerTypeIDLink,
				providerTypeInstanceIDLink,
			)

			providerTypeInstance, response, err := securityAndComplianceCenterService.GetProviderTypeInstance(getProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstance, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
		It(`UpdateProviderTypeInstance request example`, func() {
			fmt.Println("\nUpdateProviderTypeInstance() result:")
			// begin-update_provider_type_instance

			updateProviderTypeInstanceOptions := securityAndComplianceCenterService.NewUpdateProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				providerTypeIDLink,
				providerTypeInstanceIDLink,
				"Provider Type Instance from SDK run",
			)

			providerTypeInstance, response, err := securityAndComplianceCenterService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstance, "", "  ")
			fmt.Println(string(b))

			// end-update_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
		It(`GetProviderTypeByID request example`, func() {
			fmt.Println("\nGetProviderTypeByID() result:")
			// begin-get_provider_type_by_id

			getProviderTypeByIDOptions := securityAndComplianceCenterService.NewGetProviderTypeByIDOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				providerTypeIDLink,
			)

			providerType, response, err := securityAndComplianceCenterService.GetProviderTypeByID(getProviderTypeByIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerType, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_type_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerType).ToNot(BeNil())
		})
		It(`GetScanReport request example`, func() {
			fmt.Println("\nGetScanReport() result:")
			// begin-get_scan_report

			getScanReportOptions := securityAndComplianceCenterService.NewGetScanReportOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				scanIDforReportLink,
			)

			scanReport, response, err := securityAndComplianceCenterService.GetScanReport(getScanReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scanReport, "", "  ")
			fmt.Println(string(b))

			// end-get_scan_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReport).ToNot(BeNil())
		})
		It(`GetScanReportDownloadFile request example`, func() {
			fmt.Println("\nGetScanReportDownloadFile() result:")
			// begin-get_scan_report_download_file

			getScanReportDownloadFileOptions := securityAndComplianceCenterService.NewGetScanReportDownloadFileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				scanIDforReportLink,
			)

			result, response, err := securityAndComplianceCenterService.GetScanReportDownloadFile(getScanReportDownloadFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-get_scan_report_download_file

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`ListReports request example`, func() {
			fmt.Println("\nListReports() result:")
			// begin-list_reports
			listReportsOptions := &securityandcompliancecenterv3.ListReportsOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AttachmentID: &attachmentIDForReportLink,
				Limit:        core.Int64Ptr(int64(10)),
				Sort:         core.StringPtr("profile_name"),
			}

			pager, err := securityAndComplianceCenterService.NewReportsPager(listReportsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.Report
			for i := 0; pager.HasNext() && i < 5; i++ {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_reports
		})
		It(`GetReport request example`, func() {
			fmt.Println("\nGetReport() result:")
			// begin-get_report

			getReportOptions := securityAndComplianceCenterService.NewGetReportOptions(
				reportIDForReportLink,
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			report, response, err := securityAndComplianceCenterService.GetReport(getReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-get_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
		It(`GetReportSummary request example`, func() {
			fmt.Println("\nGetReportSummary() result:")
			// begin-get_report_summary

			getReportSummaryOptions := securityAndComplianceCenterService.NewGetReportSummaryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			reportSummary, response, err := securityAndComplianceCenterService.GetReportSummary(getReportSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_report_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
		It(`GetReportDownloadFile request example`, func() {
			fmt.Println("\nGetReportDownloadFile() result:")
			// begin-get_report_download_file

			getReportDownloadFileOptions := securityAndComplianceCenterService.NewGetReportDownloadFileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			result, response, err := securityAndComplianceCenterService.GetReportDownloadFile(getReportDownloadFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-get_report_download_file

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`GetReportControls request example`, func() {
			fmt.Println("\nGetReportControls() result:")
			// begin-get_report_controls

			getReportControlsOptions := securityAndComplianceCenterService.NewGetReportControlsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)
			getReportControlsOptions.SetStatus("compliant")

			reportControls, response, err := securityAndComplianceCenterService.GetReportControls(getReportControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportControls, "", "  ")
			fmt.Println(string(b))

			// end-get_report_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportControls).ToNot(BeNil())
		})
		It(`GetReportRule request example`, func() {
			fmt.Println("\nGetReportRule() result:")
			// begin-get_report_rule

			getReportRuleOptions := securityAndComplianceCenterService.NewGetReportRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				"rule-238a6025-2522-4d36-831b-a32f81f97304",
			)

			ruleInfo, response, err := securityAndComplianceCenterService.GetReportRule(getReportRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleInfo, "", "  ")
			fmt.Println(string(b))

			// end-get_report_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleInfo).ToNot(BeNil())
		})
		It(`ListReportEvaluations request example`, func() {
			fmt.Println("\nListReportEvaluations() result:")
			// begin-list_report_evaluations
			listReportEvaluationsOptions := &securityandcompliancecenterv3.ListReportEvaluationsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
				Status:     core.StringPtr("failure"),
				Limit:      core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.Evaluation
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_report_evaluations
		})
		It(`ListReportResources request example`, func() {
			fmt.Println("\nListReportResources() result:")
			// begin-list_report_resources
			listReportResourcesOptions := &securityandcompliancecenterv3.ListReportResourcesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
			}

			pager, err := securityAndComplianceCenterService.NewReportResourcesPager(listReportResourcesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.Resource
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_report_resources
		})
		It(`GetReportTags request example`, func() {
			fmt.Println("\nGetReportTags() result:")
			// begin-get_report_tags

			getReportTagsOptions := securityAndComplianceCenterService.NewGetReportTagsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			reportTags, response, err := securityAndComplianceCenterService.GetReportTags(getReportTagsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportTags, "", "  ")
			fmt.Println(string(b))

			// end-get_report_tags

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportTags).ToNot(BeNil())
		})
		It(`GetReportViolationsDrift request example`, func() {
			fmt.Println("\nGetReportViolationsDrift() result:")
			// begin-get_report_violations_drift

			getReportViolationsDriftOptions := securityAndComplianceCenterService.NewGetReportViolationsDriftOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			reportViolationsDrift, response, err := securityAndComplianceCenterService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportViolationsDrift, "", "  ")
			fmt.Println(string(b))

			// end-get_report_violations_drift

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportViolationsDrift).ToNot(BeNil())
		})
		It(`ListScanReports request example`, func() {
			fmt.Println("\nListScanReports() result:")
			// begin-list_scan_reports

			listScanReportsOptions := securityAndComplianceCenterService.NewListScanReportsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			scanReportCollection, response, err := securityAndComplianceCenterService.ListScanReports(listScanReportsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scanReportCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_scan_reports

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReportCollection).ToNot(BeNil())
		})
		It(`CreateScanReport request example`, func() {
			fmt.Println("\nCreateScanReport() result:")
			// begin-create_scan_report

			createScanReportOptions := securityAndComplianceCenterService.NewCreateScanReportOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				"csv",
			)

			createScanReport, response, err := securityAndComplianceCenterService.CreateScanReport(createScanReportOptions)
			if err != nil && response.StatusCode != 409 {
				panic(err)
			}
			b, _ := json.MarshalIndent(createScanReport, "", "  ")
			fmt.Println(string(b))

			// end-create_scan_report

			Expect(response.StatusCode).To(Or(Equal(202), Equal(409)))
			if response.StatusCode == 202 {
				Expect(err).To(BeNil())
				Expect(createScanReport).ToNot(BeNil())
			} else {
				Expect(createScanReport).To(BeNil())
				Expect(err).ToNot(BeNil())
			}
		})
		It(`CreateScan request example`, func() {
			fmt.Println("\nCreateScan() result:")
			// begin-create_scan

			createScanOptions := securityAndComplianceCenterService.NewCreateScanOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			createScanOptions.SetAttachmentID(attachmentIDLink)

			createScanResponse, response, err := securityAndComplianceCenterService.CreateScan(createScanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createScanResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_scan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createScanResponse).ToNot(BeNil())
		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules
			listRulesOptions := &securityandcompliancecenterv3.ListRulesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Type:        core.StringPtr("system_defined"),
				Search:      core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				Sort:        core.StringPtr("updated_on"),
			}

			pager, err := securityAndComplianceCenterService.NewRulesPager(listRulesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterv3.Rule
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_rules
		})
		It(`ListServices request example`, func() {
			fmt.Println("\nListServices() result:")
			// begin-list_services

			listServicesOptions := securityAndComplianceCenterService.NewListServicesOptions()

			serviceCollection, response, err := securityAndComplianceCenterService.ListServices(listServicesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_services

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceCollection).ToNot(BeNil())
		})
		It(`GetService request example`, func() {
			fmt.Println("\nGetService() result:")
			// begin-get_service

			getServiceOptions := securityAndComplianceCenterService.NewGetServiceOptions(
				"cloud-object-storage",
			)

			service, response, err := securityAndComplianceCenterService.GetService(getServiceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(service, "", "  ")
			fmt.Println(string(b))

			// end-get_service

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(service).ToNot(BeNil())
		})
		It(`DeleteProfileAttachment request example`, func() {
			fmt.Println("\nDeleteProfileAttachment() result:")
			// begin-delete_profile_attachment

			deleteProfileAttachmentOptions := securityAndComplianceCenterService.NewDeleteProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				oldProfileIDForReportLink,
				attachmentIDLink,
			)

			profileAttachment, response, err := securityAndComplianceCenterService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-delete_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`DeleteCustomControlLibrary request example`, func() {
			fmt.Println("\nDeleteCustomControlLibrary() result:")
			// begin-delete_custom_control_library

			deleteCustomControlLibraryOptions := securityAndComplianceCenterService.NewDeleteCustomControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				controlLibraryIDLink,
			)

			controlLibrary, response, err := securityAndComplianceCenterService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`DeleteCustomProfile request example`, func() {
			fmt.Println("\nDeleteCustomProfile() result:")
			// begin-delete_custom_profile

			deleteCustomProfileOptions := securityAndComplianceCenterService.NewDeleteCustomProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profile, response, err := securityAndComplianceCenterService.DeleteCustomProfile(deleteCustomProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`DeleteSubscope request example`, func() {
			// begin-delete_subscope

			deleteSubscopeOptions := securityAndComplianceCenterService.NewDeleteSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				subScopeIDLink,
			)

			response, err := securityAndComplianceCenterService.DeleteSubscope(deleteSubscopeOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSubscope(): %d\n", response.StatusCode)
			}

			// end-delete_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteScope request example`, func() {
			// begin-delete_scope

			deleteScopeOptions := securityAndComplianceCenterService.NewDeleteScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)

			response, err := securityAndComplianceCenterService.DeleteScope(deleteScopeOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteScope(): %d\n", response.StatusCode)
			}

			// end-delete_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTarget request example`, func() {
			// begin-delete_target

			deleteTargetOptions := securityAndComplianceCenterService.NewDeleteTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				targetIDLink,
			)

			response, err := securityAndComplianceCenterService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTarget(): %d\n", response.StatusCode)
			}

			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteProviderTypeInstance request example`, func() {
			// begin-delete_provider_type_instance

			deleteProviderTypeInstanceOptions := securityAndComplianceCenterService.NewDeleteProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				providerTypeIDLink,
				providerTypeInstanceIDLink,
			)

			response, err := securityAndComplianceCenterService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteProviderTypeInstance(): %d\n", response.StatusCode)
			}

			// end-delete_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteRule request example`, func() {
			// begin-delete_rule

			deleteRuleOptions := securityAndComplianceCenterService.NewDeleteRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				ruleIDLink,
			)

			response, err := securityAndComplianceCenterService.DeleteRule(deleteRuleOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRule(): %d\n", response.StatusCode)
			}

			// end-delete_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
