//go:build integration
// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	securityandcompliancecenterv3 "github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the securityandcompliancecenterv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SecurityAndComplianceCenterV3 Integration Tests`, func() {
	const externalConfigFile = "../security_and_compliance_center_v3.env"

	var (
		err                                error
		securityAndComplianceCenterService *securityandcompliancecenterv3.SecurityAndComplianceCenterApiV3
		serviceURL                         string
		config                             map[string]string

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
		scopeIDLink                                string
		scopeIDforReportLink                       string
		scanIDforReportLink                        string
		subScopeIDLink                             string
		targetIDLink                               string
		typeForReportLink                          string
		workloadProtectionCRNLink                  string
	)

	shouldSkipTest := func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(securityandcompliancecenterv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
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
			securityAndComplianceCenterServiceOptions := &securityandcompliancecenterv3.SecurityAndComplianceCenterApiV3Options{}

			securityAndComplianceCenterService, err = securityandcompliancecenterv3.NewSecurityAndComplianceCenterV3(securityAndComplianceCenterServiceOptions)
			Expect(err).To(BeNil())
			Expect(securityAndComplianceCenterService).ToNot(BeNil())
			Expect(securityAndComplianceCenterService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			securityAndComplianceCenterService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSettings - List settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &securityandcompliancecenterv3.GetSettingsOptions{
				InstanceID: &instanceIDForLink,
			}

			settings, response, err := securityAndComplianceCenterService.GetSettings(getSettingsOptions)
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
	})

	Describe(`CreateScope - Create a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScope(createScopeOptions *CreateScopeOptions)`, func() {
			scopeIDPropertyModel := &securityandcompliancecenterv3.ScopePropertyScopeID{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("ff88f007f9ff4622aac4fbc0eda36255"),
			}

			scopeTypePropertyModel := &securityandcompliancecenterv3.ScopePropertyScopeType{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			createScopeOptions := &securityandcompliancecenterv3.CreateScopeOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Name:        core.StringPtr("SDK test Sample Scope"),
				Description: core.StringPtr("The scope was created by a SDK run. This scope can be deleted"),
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterv3.ScopePropertyIntf{scopeIDPropertyModel, scopeTypePropertyModel},
			}

			scope, response, err := securityAndComplianceCenterService.CreateScope(createScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())

			scopeIDLink = *scope.ID
			fmt.Fprintf(GinkgoWriter, "Saved scopeIDLink value: %v\n", scopeIDLink)
		})
	})

	Describe(`CreateSubscope - Create a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSubscope(createSubscopeOptions *CreateSubscopeOptions)`, func() {
			scopePropertyModel := &securityandcompliancecenterv3.ScopePropertyScopeID{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("1f689f08ec9b47b885c2659c17029581"),
			}
			scopeTypePropertyModel := &securityandcompliancecenterv3.ScopePropertyScopeType{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account.resource_group"),
			}

			scopePrototypeModel := &securityandcompliancecenterv3.ScopePrototype{
				Name:        core.StringPtr("ibm subscope"),
				Description: core.StringPtr("The subscope that is defined for IBM resources."),
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterv3.ScopePropertyIntf{scopePropertyModel, scopeTypePropertyModel},
			}

			createSubscopeOptions := &securityandcompliancecenterv3.CreateSubscopeOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
				Subscopes:  []securityandcompliancecenterv3.ScopePrototype{*scopePrototypeModel},
			}

			subScopeResponse, response, err := securityAndComplianceCenterService.CreateSubscope(createSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subScopeResponse).ToNot(BeNil())

			subScopeIDLink = *subScopeResponse.Subscopes[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved subScopeIDLink value: %v\n", subScopeIDLink)
		})
	})

	Describe(`CreateCustomControlLibrary - Create a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions)`, func() {
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
				ControlDescription:    core.StringPtr("This control library can be deleted"),
				ControlCategory:       core.StringPtr("test-control"),
				ControlRequirement:    core.BoolPtr(true),
				ControlParent:         core.StringPtr(""),
				ControlSpecifications: []securityandcompliancecenterv3.ControlSpecificationPrototype{*controlSpecificationPrototypeModel},
				ControlDocs:           controlDocModel,
				Status:                core.StringPtr("disabled"),
			}

			createCustomControlLibraryOptions := &securityandcompliancecenterv3.CreateCustomControlLibraryOptions{
				InstanceID:                core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryName:        core.StringPtr("custom control library from SDK"),
				ControlLibraryDescription: core.StringPtr("This is a custom control library made from the SDK test framework"),
				ControlLibraryType:        core.StringPtr("custom"),
				ControlLibraryVersion:     core.StringPtr("0.0.1"),
				Controls:                  []securityandcompliancecenterv3.ControlPrototype{*controlPrototypeModel},
				AccountID:                 &accountIDForReportLink,
			}

			controlLibrary, response, err := securityAndComplianceCenterService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(controlLibrary).ToNot(BeNil())

			controlLibraryIDLink = *controlLibrary.ID
			fmt.Fprintf(GinkgoWriter, "Saved controlLibraryIDLink value: %v\n", controlLibraryIDLink)
		})
	})

	Describe(`CreateProfile - Create a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
			profileAssessmentPrototypeModel := &securityandcompliancecenterv3.ProfileAssessmentPrototype{
				AssessmentID: core.StringPtr("testString"),
			}

			profileControlSpecificationPrototypeModel := &securityandcompliancecenterv3.ProfileControlSpecificationPrototype{
				ControlSpecificationID: core.StringPtr("testString"),
				Assessments:            []securityandcompliancecenterv3.ProfileAssessmentPrototype{*profileAssessmentPrototypeModel},
			}

			profileControlsPrototypeModel := &securityandcompliancecenterv3.ProfileControlsPrototype{
				ControlLibraryID:      core.StringPtr("51ca566e-c559-412b-8d64-f05b57044c32"),
				ControlID:             core.StringPtr("2ce21ba3-0548-49a3-88e2-1122632218f4"),
				ControlSpecifications: []securityandcompliancecenterv3.ProfileControlSpecificationPrototype{*profileControlSpecificationPrototypeModel},
			}

			defaultParametersPrototypeModel := &securityandcompliancecenterv3.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr(`["1.2","1.3"]`),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			createProfileOptions := &securityandcompliancecenterv3.CreateProfileOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileName:        core.StringPtr("Profile Example"),
				ProfileDescription: core.StringPtr("This is a profile made from a SDK run"),
				ProfileVersion:     core.StringPtr("0.0.1"),
				Latest:             core.BoolPtr(true),
				Controls:           []securityandcompliancecenterv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []securityandcompliancecenterv3.DefaultParametersPrototype{*defaultParametersPrototypeModel},
				AccountID:          &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterService.CreateProfile(createProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())

			if profileIDLink == "" {
				profileIDLink = *profile.ID
			}
			fmt.Fprintf(GinkgoWriter, "Saved profileIDLink value: %v\n", profileIDLink)
		})
	})

	Describe(`CreateProfileAttachment - Create a profile attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfileAttachment(createProfileAttachmentOptions *CreateProfileAttachmentOptions)`, func() {
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

			attachmentNotificationsControlsModel := &securityandcompliancecenterv3.AttachmentNotificationsControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentNotificationsModel := &securityandcompliancecenterv3.AttachmentNotifications{
				Enabled:  core.BoolPtr(true),
				Controls: attachmentNotificationsControlsModel,
			}

			multiCloudScopePayloadModel := &securityandcompliancecenterv3.MultiCloudScopePayload{
				ID: core.StringPtr("a1c2a74e-508d-48b3-8e46-5ab4f4795a7b"),
			}

			profileAttachmentBaseModel := &securityandcompliancecenterv3.ProfileAttachmentBase{
				AttachmentParameters: []securityandcompliancecenterv3.Parameter{*parameterModel1, *parameterModel2, *parameterModel3, *parameterModel4, *parameterModel5, *parameterModel6},
				Description:          core.StringPtr("This is a profile attachment targeting IBM CIS Foundation using a SDK"),
				Name:                 core.StringPtr("SDK test Profile Attachment for IBM CIS Foundation"),
				Notifications:        attachmentNotificationsModel,
				Schedule:             core.StringPtr("daily"),
				Scope:                []securityandcompliancecenterv3.MultiCloudScopePayload{*multiCloudScopePayloadModel},
				Status:               core.StringPtr("disabled"),
			}

			createProfileAttachmentOptions := &securityandcompliancecenterv3.CreateProfileAttachmentOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:   &oldProfileIDForReportLink,
				Attachments: []securityandcompliancecenterv3.ProfileAttachmentBase{*profileAttachmentBaseModel},
				AccountID:   &accountIDForReportLink,
			}

			profileAttachmentResponse, response, err := securityAndComplianceCenterService.CreateProfileAttachment(createProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profileAttachmentResponse).ToNot(BeNil())

			attachmentIDLink = *profileAttachmentResponse.Attachments[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIDLink value: %v\n", attachmentIDLink)
		})
	})

	Describe(`GetLatestReports - List latest reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions)`, func() {
			getLatestReportsOptions := &securityandcompliancecenterv3.GetLatestReportsOptions{
				InstanceID: &instanceIDForLink,
				Sort:       core.StringPtr("profile_name"),
			}

			reportLatest, response, err := securityAndComplianceCenterService.GetLatestReports(getLatestReportsOptions)
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
	})

	Describe(`CreateRule - Create a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
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
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("hard_quota"),
				Operator:    core.StringPtr("num_equals"),
				Value:       core.StringPtr("${hard_quota}"),
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

			createRuleOptions := &securityandcompliancecenterv3.CreateRuleOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Description:    core.StringPtr("Example rule"),
				Target:         ruleTargetPrototypeModel,
				RequiredConfig: requiredConfigModel,
				Version:        core.StringPtr("1.0.0"),
				Import:         importModel,
				Labels:         []string{},
			}

			rule, response, err := securityAndComplianceCenterService.CreateRule(createRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			ruleIDLink = *rule.ID
			fmt.Fprintf(GinkgoWriter, "Saved ruleIDLink value: %v\n", ruleIDLink)
		})
	})

	Describe(`GetRule - Get a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &securityandcompliancecenterv3.GetRuleOptions{
				InstanceID: &instanceIDForLink,
				RuleID:     &ruleIDLink,
			}

			rule, response, err := securityAndComplianceCenterService.GetRule(getRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("etag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
	})

	Describe(`ReplaceRule - Update a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
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
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("hard_quota"),
				Operator:    core.StringPtr("num_equals"),
				Value:       core.StringPtr("${hard_quota}"),
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

			replaceRuleOptions := &securityandcompliancecenterv3.ReplaceRuleOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				RuleID:         &ruleIDLink,
				IfMatch:        &eTagLink,
				Description:    core.StringPtr("Example SDK rule. This rule can be deleted."),
				Target:         ruleTargetPrototypeModel,
				RequiredConfig: requiredConfigModel,
				Version:        core.StringPtr("1.0.1"),
				Import:         importModel,
				Labels:         []string{},
			}

			rule, response, err := securityAndComplianceCenterService.ReplaceRule(replaceRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
	})

	Describe(`UpdateSettings - Update settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			objectStoragePrototypeModel := &securityandcompliancecenterv3.ObjectStoragePrototype{
				Bucket:      core.StringPtr("px-scan-results"),
				InstanceCRN: core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::"),
			}

			eventNotificationsPrototypeModel := &securityandcompliancecenterv3.EventNotificationsPrototype{
				InstanceCRN:       core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"),
				SourceDescription: core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center."),
				SourceName:        core.StringPtr("scc-sdk-integration"),
			}

			updateSettingsOptions := &securityandcompliancecenterv3.UpdateSettingsOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ObjectStorage:      objectStoragePrototypeModel,
				EventNotifications: eventNotificationsPrototypeModel,
			}

			settings, response, err := securityAndComplianceCenterService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(200), Equal(204)))
			if response.StatusCode == 200 {
				Expect(settings).ToNot(BeNil())
			} else {
				Expect(settings).To(BeNil())
			}
		})
	})

	Describe(`PostTestEvent - Create a test event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostTestEvent(postTestEventOptions *PostTestEventOptions)`, func() {
			postTestEventOptions := &securityandcompliancecenterv3.PostTestEventOptions{
				InstanceID: &instanceIDForLink,
			}

			testEvent, response, err := securityAndComplianceCenterService.PostTestEvent(postTestEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceAttachments - Get all instance attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceAttachments(listInstanceAttachmentsOptions *ListInstanceAttachmentsOptions) with pagination`, func() {
			listInstanceAttachmentsOptions := &securityandcompliancecenterv3.ListInstanceAttachmentsOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:         &accountIDForReportLink,
				VersionGroupLabel: core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d"),
				Limit:             core.Int64Ptr(int64(10)),
				Sort:              core.StringPtr("created_on"),
				Direction:         core.StringPtr("desc"),
				Start:             core.StringPtr("testString"),
			}

			listInstanceAttachmentsOptions.Start = nil
			listInstanceAttachmentsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterv3.ProfileAttachment
			for {
				profileAttachmentCollection, response, err := securityAndComplianceCenterService.ListInstanceAttachments(listInstanceAttachmentsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(profileAttachmentCollection).ToNot(BeNil())
				allResults = append(allResults, profileAttachmentCollection.Attachments...)

				listInstanceAttachmentsOptions.Start, err = profileAttachmentCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listInstanceAttachmentsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListInstanceAttachments(listInstanceAttachmentsOptions *ListInstanceAttachmentsOptions) using InstanceAttachmentsPager`, func() {
			listInstanceAttachmentsOptions := &securityandcompliancecenterv3.ListInstanceAttachmentsOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:         &accountIDForReportLink,
				VersionGroupLabel: core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d"),
				Limit:             core.Int64Ptr(int64(10)),
				Sort:              core.StringPtr("created_on"),
				Direction:         core.StringPtr("desc"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewInstanceAttachmentsPager(listInstanceAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.ProfileAttachment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewInstanceAttachmentsPager(listInstanceAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListInstanceAttachments() returned a total of %d item(s) using InstanceAttachmentsPager.\n", len(allResults))
		})
	})

	Describe(`ListProfileAttachments - Get all attachments tied to a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileAttachments(listProfileAttachmentsOptions *ListProfileAttachmentsOptions)`, func() {
			listProfileAttachmentsOptions := &securityandcompliancecenterv3.ListProfileAttachmentsOptions{
				InstanceID: &instanceIDForLink,
				ProfileID:  &oldProfileIDForReportLink,
				AccountID:  &accountIDForReportLink,
			}

			profileAttachmentCollection, response, err := securityAndComplianceCenterService.ListProfileAttachments(listProfileAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachmentCollection).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions)`, func() {
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

			attachmentNotificationsControlsModel := &securityandcompliancecenterv3.AttachmentNotificationsControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentNotificationsModel := &securityandcompliancecenterv3.AttachmentNotifications{
				Enabled:  core.BoolPtr(false),
				Controls: attachmentNotificationsControlsModel,
			}

			multiCloudScopePayloadModel := &securityandcompliancecenterv3.MultiCloudScopePayload{
				ID: core.StringPtr("a1c2a74e-508d-48b3-8e46-5ab4f4795a7b"),
			}

			replaceProfileAttachmentOptions := &securityandcompliancecenterv3.ReplaceProfileAttachmentOptions{
				InstanceID:           core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:            &oldProfileIDForReportLink,
				AttachmentID:         &attachmentIDLink,
				AttachmentParameters: []securityandcompliancecenterv3.Parameter{*parameterModel1, *parameterModel2, *parameterModel3, *parameterModel4, *parameterModel5, *parameterModel6},
				Description:          core.StringPtr("testString"),
				Name:                 core.StringPtr("testString"),
				Notifications:        attachmentNotificationsModel,
				Schedule:             core.StringPtr("daily"),
				Scope:                []securityandcompliancecenterv3.MultiCloudScopePayload{*multiCloudScopePayloadModel},
				Status:               core.StringPtr("disabled"),
				AccountID:            &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`GetProfileAttachment - Get an attachment for a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions)`, func() {
			getProfileAttachmentOptions := &securityandcompliancecenterv3.GetProfileAttachmentOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:    &profileIDLink,
				AttachmentID: &attachmentIDLink,
				AccountID:    &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterService.GetProfileAttachment(getProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`UpgradeAttachment - Upgrade an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpgradeAttachment(upgradeAttachmentOptions *UpgradeAttachmentOptions)`, func() {
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

			upgradeAttachmentOptions := &securityandcompliancecenterv3.UpgradeAttachmentOptions{
				InstanceID:           core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:            &oldProfileIDForReportLink,
				AttachmentID:         &attachmentIDLink,
				AttachmentParameters: []securityandcompliancecenterv3.Parameter{*parameterModel1, *parameterModel2, *parameterModel3, *parameterModel4, *parameterModel5, *parameterModel6},
				AccountID:            &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterService.UpgradeAttachment(upgradeAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`ListControlLibraries - Get all control libraries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions)`, func() {
			listControlLibrariesOptions := &securityandcompliancecenterv3.ListControlLibrariesOptions{
				InstanceID: &instanceIDForLink,
				AccountID:  &accountIDForReportLink,
			}

			controlLibraryCollection, response, err := securityAndComplianceCenterService.ListControlLibraries(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryCollection).ToNot(BeNil())
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) using ControlLibrariesPager`, func() {
			listControlLibrariesOptions := &securityandcompliancecenterv3.ListControlLibrariesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewControlLibrariesPager(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.ControlLibrary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewControlLibrariesPager(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListControlLibraries() returned a total of %d item(s) using ControlLibrariesPager.\n", len(allResults))
		})
	})

	Describe(`ReplaceCustomControlLibrary - Update a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions)`, func() {
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

			replaceCustomControlLibraryOptions := &securityandcompliancecenterv3.ReplaceCustomControlLibraryOptions{
				InstanceID:                core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryID:          &controlLibraryIDLink,
				ControlLibraryName:        core.StringPtr("Custom Library from SDK gen"),
				ControlLibraryDescription: core.StringPtr("This control library can be deleted"),
				ControlLibraryType:        core.StringPtr("custom"),
				ControlLibraryVersion:     core.StringPtr("0.0.2"),
				Controls:                  []securityandcompliancecenterv3.Control{*controlPrototypeModel},
			}

			controlLibrary, response, err := securityAndComplianceCenterService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`GetControlLibrary - Get a control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions)`, func() {
			getControlLibraryOptions := &securityandcompliancecenterv3.GetControlLibraryOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryID: &controlLibraryIDLink,
				AccountID:        &accountIDForReportLink,
			}

			controlLibrary, response, err := securityAndComplianceCenterService.GetControlLibrary(getControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`ListProfiles - Get all profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {
			listProfilesOptions := &securityandcompliancecenterv3.ListProfilesOptions{
				InstanceID: &instanceIDForLink,
				AccountID:  &accountIDForReportLink,
			}

			profileCollection, response, err := securityAndComplianceCenterService.ListProfiles(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileCollection).ToNot(BeNil())
		})
		It(`ListProfiles(listProfilesOptions *ListProfileOptions) using ProfilesPager`, func() {
			listProfilesOptions := &securityandcompliancecenterv3.ListProfilesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewProfilesPager(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.Profile
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewProfilesPager(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProfiles() returned a total of %d item(s) using ProfilesPager.\n", len(allResults))
		})
	})

	Describe(`ReplaceProfile - Update a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfile(replaceProfileOptions *ReplaceProfileOptions)`, func() {
			profileAssessmentPrototypeModel := &securityandcompliancecenterv3.ProfileAssessmentPrototype{
				AssessmentID: core.StringPtr("testString"),
			}

			profileControlSpecificationPrototypeModel := &securityandcompliancecenterv3.ProfileControlSpecificationPrototype{
				ControlSpecificationID: core.StringPtr("testString"),
				Assessments:            []securityandcompliancecenterv3.ProfileAssessmentPrototype{*profileAssessmentPrototypeModel},
			}

			profileControlsPrototypeModel := &securityandcompliancecenterv3.ProfileControlsPrototype{
				ControlLibraryID:      core.StringPtr("51ca566e-c559-412b-8d64-f05b57044c32"),
				ControlID:             core.StringPtr("2ce21ba3-0548-49a3-88e2-1122632218f4"),
				ControlSpecifications: []securityandcompliancecenterv3.ProfileControlSpecificationPrototype{*profileControlSpecificationPrototypeModel},
			}

			defaultParametersPrototypeModel := &securityandcompliancecenterv3.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr(`["1.2","1.3"]`),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			replaceProfileOptions := &securityandcompliancecenterv3.ReplaceProfileOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:          &profileIDLink,
				ProfileName:        core.StringPtr("SDK test run"),
				ProfileType:        core.StringPtr("custom"),
				ProfileDescription: core.StringPtr("This profile can be deleted"),
				ProfileVersion:     core.StringPtr("0.0.2"),
				Latest:             core.BoolPtr(true),
				Controls:           []securityandcompliancecenterv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []securityandcompliancecenterv3.DefaultParametersPrototype{*defaultParametersPrototypeModel},
				AccountID:          &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterService.ReplaceProfile(replaceProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`GetProfile - Get a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {
			getProfileOptions := &securityandcompliancecenterv3.GetProfileOptions{
				InstanceID: &instanceIDForLink,
				ProfileID:  &profileIDLink,
				AccountID:  &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterService.GetProfile(getProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileParameters - Update custom profile parameters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileParameters(replaceProfileParametersOptions *ReplaceProfileParametersOptions)`, func() {
			defaultParametersModel := &securityandcompliancecenterv3.DefaultParameters{
				AssessmentType:        core.StringPtr("testString"),
				AssessmentID:          core.StringPtr("testString"),
				ParameterName:         core.StringPtr("testString"),
				ParameterDefaultValue: core.StringPtr("testString"),
				ParameterDisplayName:  core.StringPtr("testString"),
				ParameterType:         core.StringPtr("testString"),
			}

			replaceProfileParametersOptions := &securityandcompliancecenterv3.ReplaceProfileParametersOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:         &profileIDLink,
				ID:                core.StringPtr("testString"),
				DefaultParameters: []securityandcompliancecenterv3.DefaultParameters{*defaultParametersModel},
				AccountID:         &accountIDForReportLink,
			}

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterService.ReplaceProfileParameters(replaceProfileParametersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
	})

	Describe(`ListProfileParameters - List profile parameters for a given profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileParameters(listProfileParametersOptions *ListProfileParametersOptions)`, func() {
			listProfileParametersOptions := &securityandcompliancecenterv3.ListProfileParametersOptions{
				InstanceID: &instanceIDForLink,
				ProfileID:  &profileIDLink,
			}

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterService.ListProfileParameters(listProfileParametersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
	})

	Describe(`CompareProfiles - Compare profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompareProfiles(compareProfilesOptions *CompareProfilesOptions)`, func() {
			compareProfilesOptions := &securityandcompliancecenterv3.CompareProfilesOptions{
				InstanceID: &instanceIDForLink,
				ProfileID:  &oldProfileIDForReportLink,
				AccountID:  &accountIDForReportLink,
			}

			comparePredefinedProfilesResponse, response, err := securityAndComplianceCenterService.CompareProfiles(compareProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(comparePredefinedProfilesResponse).ToNot(BeNil())
		})
	})

	Describe(`ListScopes - Get all scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions) with pagination`, func() {
			listScopesOptions := &securityandcompliancecenterv3.ListScopesOptions{
				InstanceID: &instanceIDForLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			listScopesOptions.Start = nil
			listScopesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterv3.Scope
			for {
				scopeCollection, response, err := securityAndComplianceCenterService.ListScopes(listScopesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(scopeCollection).ToNot(BeNil())
				allResults = append(allResults, scopeCollection.Scopes...)

				listScopesOptions.Start, err = scopeCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listScopesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions) using ScopesPager`, func() {
			listScopesOptions := &securityandcompliancecenterv3.ListScopesOptions{
				InstanceID: &instanceIDForLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewScopesPager(listScopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.Scope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewScopesPager(listScopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListScopes() returned a total of %d item(s) using ScopesPager.\n", len(allResults))
		})
	})

	Describe(`UpdateScope - Update a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateScope(updateScopeOptions *UpdateScopeOptions)`, func() {
			updateScopeOptions := &securityandcompliancecenterv3.UpdateScopeOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				Name:        core.StringPtr("updated name of scope SDK"),
				Description: core.StringPtr("updated scope description. This scope can be deleted."),
			}

			scope, response, err := securityAndComplianceCenterService.UpdateScope(updateScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
	})

	Describe(`GetScope - Get a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScope(getScopeOptions *GetScopeOptions)`, func() {
			getScopeOptions := &securityandcompliancecenterv3.GetScopeOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
			}

			scope, response, err := securityAndComplianceCenterService.GetScope(getScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
	})

	Describe(`ListSubscopes - Get all subscopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSubscopes(listSubscopesOptions *ListSubscopesOptions) with pagination`, func() {
			listSubscopesOptions := &securityandcompliancecenterv3.ListSubscopesOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
				Limit:      core.Int64Ptr(int64(10)),
				Start:      core.StringPtr("testString"),
			}

			listSubscopesOptions.Start = nil
			listSubscopesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterv3.SubScope
			for {
				subScopeCollection, response, err := securityAndComplianceCenterService.ListSubscopes(listSubscopesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(subScopeCollection).ToNot(BeNil())
				allResults = append(allResults, subScopeCollection.Subscopes...)

				listSubscopesOptions.Start, err = subScopeCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listSubscopesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSubscopes(listSubscopesOptions *ListSubscopesOptions) using SubscopesPager`, func() {
			listSubscopesOptions := &securityandcompliancecenterv3.ListSubscopesOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewSubscopesPager(listSubscopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.SubScope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewSubscopesPager(listSubscopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSubscopes() returned a total of %d item(s) using SubscopesPager.\n", len(allResults))
		})
	})

	Describe(`GetSubscope - Get a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubscope(getSubscopeOptions *GetSubscopeOptions)`, func() {
			getSubscopeOptions := &securityandcompliancecenterv3.GetSubscopeOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
				SubscopeID: &subScopeIDLink,
			}

			subScope, response, err := securityAndComplianceCenterService.GetSubscope(getSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
	})

	Describe(`UpdateSubscope - Update a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSubscope(updateSubscopeOptions *UpdateSubscopeOptions)`, func() {
			updateSubscopeOptions := &securityandcompliancecenterv3.UpdateSubscopeOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				SubscopeID:  &subScopeIDLink,
				Name:        core.StringPtr("SDK updated name of subscope"),
				Description: core.StringPtr("updated scope description. This subscope can be deleted"),
			}

			subScope, response, err := securityAndComplianceCenterService.UpdateSubscope(updateSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
	})

	Describe(`CreateTarget - Create a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
			// accountModel := &securityandcompliancecenterv3.Account{
			// 	ID: core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
			// 	Name: core.StringPtr("NIST"),
			// 	Type: core.StringPtr("account_type"),
			// }

			// resourceModel := &securityandcompliancecenterv3.Resource{
			// 	ReportID: core.StringPtr("30b434b3-cb08-4845-af10-7a8fc682b6a8"),
			// 	HomeAccountID: core.StringPtr("2411ffdc16844b07b42521c3443f456d"),
			// 	ID: core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::"),
			// 	ResourceName: core.StringPtr("jeff's key"),
			// 	Account: accountModel,
			// 	ComponentID: core.StringPtr("cloud-object_storage"),
			// 	ComponentName: core.StringPtr("cloud-object_storage"),
			// 	Environment: core.StringPtr("ibm cloud"),
			// 	Status: core.StringPtr("compliant"),
			// 	TotalCount: core.Int64Ptr(int64(140)),
			// 	PassCount: core.Int64Ptr(int64(123)),
			// 	FailureCount: core.Int64Ptr(int64(12)),
			// 	ErrorCount: core.Int64Ptr(int64(5)),
			// 	SkippedCount: core.Int64Ptr(int64(7)),
			// 	CompletedCount: core.Int64Ptr(int64(135)),
			// 	ServiceName: core.StringPtr("pm-20"),
			// 	InstanceCRN: core.StringPtr("testString"),
			// }

			// credentialModel := &securityandcompliancecenterv3.Credential{
			// 	SecretCRN: core.StringPtr("dummy"),
			// 	Resources: []securityandcompliancecenterv3.Resource{*resourceModel},
			// }

			createTargetOptions := &securityandcompliancecenterv3.CreateTargetOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:        core.StringPtr("62ecf99b240144dea9125666249edfcb"),
				TrustedProfileID: core.StringPtr("Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54"),
				Name:             core.StringPtr("Sample Target from SDK run"),
				// Credentials: []securityandcompliancecenterv3.Credential{*credentialModel},
			}

			target, response, err := securityAndComplianceCenterService.CreateTarget(createTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())
			targetIDLink = *target.ID
		})
	})

	Describe(`ListTargets - Get a list of targets with pagination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {
			listTargetsOptions := &securityandcompliancecenterv3.ListTargetsOptions{
				InstanceID: &instanceIDForLink,
			}

			targetCollection, response, err := securityAndComplianceCenterService.ListTargets(listTargetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})
	})

	Describe(`GetTarget - Get a target by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
			getTargetOptions := &securityandcompliancecenterv3.GetTargetOptions{
				InstanceID: &instanceIDForLink,
				TargetID:   &targetIDLink,
			}

			target, response, err := securityAndComplianceCenterService.GetTarget(getTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`ReplaceTarget - replace a target by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {
			replaceTargetOptions := &securityandcompliancecenterv3.ReplaceTargetOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				TargetID:         &targetIDLink,
				AccountID:        core.StringPtr("62ecf99b240144dea9125666249edfcb"),
				TrustedProfileID: core.StringPtr("Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54"),
				Name:             core.StringPtr("Target creation from SDK run. This can be deleted"),
			}

			target, response, err := securityAndComplianceCenterService.ReplaceTarget(replaceTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`ListProviderTypes - List provider types`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviderTypes(listProviderTypesOptions *ListProviderTypesOptions)`, func() {
			listProviderTypesOptions := &securityandcompliancecenterv3.ListProviderTypesOptions{
				InstanceID: &instanceIDForLink,
			}

			providerTypeCollection, response, err := securityAndComplianceCenterService.ListProviderTypes(listProviderTypesOptions)
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
	})

	Describe(`CreateProviderTypeInstance - Create a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProviderTypeInstance(createProviderTypeInstanceOptions *CreateProviderTypeInstanceOptions)`, func() {
			createProviderTypeInstanceOptions := &securityandcompliancecenterv3.CreateProviderTypeInstanceOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID: &providerTypeIDLink,
				Name:           core.StringPtr("caveonix-instance-1"),
			}

			providerTypeInstance, response, err := securityAndComplianceCenterService.CreateProviderTypeInstance(createProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(providerTypeInstance).ToNot(BeNil())

			providerTypeInstanceIDLink = *providerTypeInstance.ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeInstanceIDLink value: %v\n", providerTypeInstanceIDLink)
		})
	})

	Describe(`ListProviderTypeInstances - List instances of a specific provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviderTypeInstances(listProviderTypeInstancesOptions *ListProviderTypeInstancesOptions)`, func() {
			listProviderTypeInstancesOptions := &securityandcompliancecenterv3.ListProviderTypeInstancesOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID: &providerTypeIDLink,
			}

			providerTypeInstanceCollection, response, err := securityAndComplianceCenterService.ListProviderTypeInstances(listProviderTypeInstancesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceCollection).ToNot(BeNil())
		})
	})

	Describe(`GetProviderTypeInstance - Get a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypeInstance(getProviderTypeInstanceOptions *GetProviderTypeInstanceOptions)`, func() {
			getProviderTypeInstanceOptions := &securityandcompliancecenterv3.GetProviderTypeInstanceOptions{
				InstanceID:             core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID:         &providerTypeIDLink,
				ProviderTypeInstanceID: &providerTypeInstanceIDLink,
			}

			providerTypeInstance, response, err := securityAndComplianceCenterService.GetProviderTypeInstance(getProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
	})

	Describe(`UpdateProviderTypeInstance - Update a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProviderTypeInstance(updateProviderTypeInstanceOptions *UpdateProviderTypeInstanceOptions)`, func() {
			updateProviderTypeInstanceOptions := &securityandcompliancecenterv3.UpdateProviderTypeInstanceOptions{
				InstanceID:             core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID:         &providerTypeIDLink,
				ProviderTypeInstanceID: &providerTypeInstanceIDLink,
				Name:                   core.StringPtr("caveonix-update-instance-1"),
			}

			providerTypeInstance, response, err := securityAndComplianceCenterService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
	})

	Describe(`GetProviderTypeByID - Get a provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypeByID(getProviderTypeByIDOptions *GetProviderTypeByIDOptions)`, func() {
			getProviderTypeByIDOptions := &securityandcompliancecenterv3.GetProviderTypeByIDOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID: &providerTypeIDLink,
			}

			providerType, response, err := securityAndComplianceCenterService.GetProviderTypeByID(getProviderTypeByIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerType).ToNot(BeNil())
		})
	})

	Describe(`GetScanReport - Get a scan report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScanReport(getScanReportOptions *GetScanReportOptions)`, func() {
			getScanReportOptions := &securityandcompliancecenterv3.GetScanReportOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				JobID:      &scanIDforReportLink,
			}

			scanReport, response, err := securityAndComplianceCenterService.GetScanReport(getScanReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReport).ToNot(BeNil())
		})
	})

	Describe(`GetScanReportDownloadFile - Get a scan report details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScanReportDownloadFile(getScanReportDownloadFileOptions *GetScanReportDownloadFileOptions)`, func() {
			getScanReportDownloadFileOptions := &securityandcompliancecenterv3.GetScanReportDownloadFileOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				JobID:      &scanIDforReportLink,
				Accept:     core.StringPtr("application/csv"),
			}

			result, response, err := securityAndComplianceCenterService.GetScanReportDownloadFile(getScanReportDownloadFileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ListReports - List reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) with pagination`, func() {
			listReportsOptions := &securityandcompliancecenterv3.ListReportsOptions{
				InstanceID: &instanceIDForLink,
				Limit:      core.Int64Ptr(int64(10)),
				Sort:       core.StringPtr("profile_name"),
			}

			listReportsOptions.Start = nil
			listReportsOptions.Limit = core.Int64Ptr(10)

			var allResults []securityandcompliancecenterv3.Report
			for i := 0; i < 1; i++ {
				reportCollection, response, err := securityAndComplianceCenterService.ListReports(listReportsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reportCollection).ToNot(BeNil())
				allResults = append(allResults, reportCollection.Reports...)
				fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) within a page.\n", len(allResults))

				listReportsOptions.Start, err = reportCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listReportsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) using ReportsPager`, func() {
			listReportsOptions := &securityandcompliancecenterv3.ListReportsOptions{
				InstanceID: &instanceIDForLink,
				Limit:      core.Int64Ptr(int64(10)),
				Sort:       core.StringPtr("profile_name"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewReportsPager(listReportsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.Report
			for i := 0; pager.HasNext() && i < 1; i++ {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
				fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) within a page using ReportsPager.\n", len(allResults))
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewReportsPager(listReportsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())
			// Manual Change: Took out GetAll because it would take too long
		})
	})

	Describe(`GetReport - Get a report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReport(getReportOptions *GetReportOptions)`, func() {
			getReportOptions := &securityandcompliancecenterv3.GetReportOptions{
				ReportID:   &reportIDForReportLink,
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDforReportLink,
			}

			report, response, err := securityAndComplianceCenterService.GetReport(getReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
	})

	Describe(`GetReportSummary - Get a report summary`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions)`, func() {
			getReportSummaryOptions := &securityandcompliancecenterv3.GetReportSummaryOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
			}

			reportSummary, response, err := securityAndComplianceCenterService.GetReportSummary(getReportSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
	})

	Describe(`GetReportDownloadFile - Get report evaluation details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportDownloadFile(getReportDownloadFileOptions *GetReportDownloadFileOptions)`, func() {
			getReportDownloadFileOptions := &securityandcompliancecenterv3.GetReportDownloadFileOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:       &reportIDForReportLink,
				Accept:         core.StringPtr("application/csv"),
				ExcludeSummary: core.BoolPtr(true),
			}

			result, response, err := securityAndComplianceCenterService.GetReportDownloadFile(getReportDownloadFileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetReportControls - Get report controls`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportControls(getReportControlsOptions *GetReportControlsOptions)`, func() {
			getReportControlsOptions := &securityandcompliancecenterv3.GetReportControlsOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				Status:     core.StringPtr("compliant"),
				Sort:       core.StringPtr("control_name"),
				ScopeID:    &scopeIDforReportLink,
			}

			reportControls, response, err := securityAndComplianceCenterService.GetReportControls(getReportControlsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportControls).ToNot(BeNil())
		})
	})

	Describe(`GetReportRule - Get a report rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportRule(getReportRuleOptions *GetReportRuleOptions)`, func() {
			getReportRuleOptions := &securityandcompliancecenterv3.GetReportRuleOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				RuleID:     core.StringPtr("rule-238a6025-2522-4d36-831b-a32f81f97304"),
			}

			ruleInfo, response, err := securityAndComplianceCenterService.GetReportRule(getReportRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleInfo).ToNot(BeNil())
		})
	})

	Describe(`ListReportEvaluations - List report evaluations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) with pagination`, func() {
			listReportEvaluationsOptions := &securityandcompliancecenterv3.ListReportEvaluationsOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				Status:     core.StringPtr("failure"),
				Start:      core.StringPtr("testString"),
				Limit:      core.Int64Ptr(int64(10)),
			}

			listReportEvaluationsOptions.Start = nil
			listReportEvaluationsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterv3.Evaluation
			for {
				evaluationPage, response, err := securityAndComplianceCenterService.ListReportEvaluations(listReportEvaluationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(evaluationPage).ToNot(BeNil())
				allResults = append(allResults, evaluationPage.Evaluations...)

				listReportEvaluationsOptions.Start, err = evaluationPage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportEvaluationsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) using ReportEvaluationsPager`, func() {
			listReportEvaluationsOptions := &securityandcompliancecenterv3.ListReportEvaluationsOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				Status:     core.StringPtr("failure"),
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.Evaluation
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportEvaluations() returned a total of %d item(s) using ReportEvaluationsPager.\n", len(allResults))
		})
	})

	Describe(`ListReportResources - List report resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) with pagination`, func() {
			listReportResourcesOptions := &securityandcompliancecenterv3.ListReportResourcesOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				AccountID:  &accountIDForReportLink,
				Status:     core.StringPtr("compliant"),
				Limit:      core.Int64Ptr(int64(10)),
				ScopeID:    &scopeIDforReportLink,
			}

			listReportResourcesOptions.Start = nil
			listReportResourcesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterv3.Resource
			for {
				resourcePage, response, err := securityAndComplianceCenterService.ListReportResources(listReportResourcesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(resourcePage).ToNot(BeNil())
				allResults = append(allResults, resourcePage.Resources...)

				listReportResourcesOptions.Start, err = resourcePage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportResourcesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) using ReportResourcesPager`, func() {
			listReportResourcesOptions := &securityandcompliancecenterv3.ListReportResourcesOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.Resource
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportResources() returned a total of %d item(s) using ReportResourcesPager.\n", len(allResults))
		})
	})

	Describe(`GetReportTags - List report tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportTags(getReportTagsOptions *GetReportTagsOptions)`, func() {
			getReportTagsOptions := &securityandcompliancecenterv3.GetReportTagsOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
			}

			reportTags, response, err := securityAndComplianceCenterService.GetReportTags(getReportTagsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportTags).ToNot(BeNil())
		})
	})

	Describe(`GetReportViolationsDrift - Get report violations drift`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions)`, func() {
			getReportViolationsDriftOptions := &securityandcompliancecenterv3.GetReportViolationsDriftOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				ScopeID:    &scopeIDforReportLink,
			}

			reportViolationsDrift, response, err := securityAndComplianceCenterService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportViolationsDrift).ToNot(BeNil())
		})
	})

	Describe(`ListScanReports - List scan reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScanReports(listScanReportsOptions *ListScanReportsOptions)`, func() {
			listScanReportsOptions := &securityandcompliancecenterv3.ListScanReportsOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				ScopeID:    &scopeIDforReportLink,
			}

			scanReportCollection, response, err := securityAndComplianceCenterService.ListScanReports(listScanReportsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReportCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateScanReport - Create a scan report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScanReport(createScanReportOptions *CreateScanReportOptions)`, func() {
			createScanReportOptions := &securityandcompliancecenterv3.CreateScanReportOptions{
				InstanceID: &instanceIDForLink,
				ReportID:   &reportIDForReportLink,
				Format:     core.StringPtr("csv"),
				ScopeID:    &scopeIDLink,
			}

			createScanReport, response, err := securityAndComplianceCenterService.CreateScanReport(createScanReportOptions)
			Expect(response.StatusCode).To(Or(Equal(202), Equal(404), Equal(409)))
			if response.StatusCode == 202 {
				Expect(err).To(BeNil())
				Expect(createScanReport).ToNot(BeNil())
			} else {
				Expect(createScanReport).To(BeNil())
				Expect(err).ToNot(BeNil())
			}
		})
	})

	Describe(`CreateScan - Create a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScan(createScanOptions *CreateScanOptions)`, func() {
			createScanOptions := &securityandcompliancecenterv3.CreateScanOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AttachmentID: &attachmentIDLink,
				AccountID:    &accountIDForReportLink,
			}

			createScanResponse, response, err := securityAndComplianceCenterService.CreateScan(createScanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createScanResponse).ToNot(BeNil())
		})
	})

	Describe(`ListRules - Get all rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) with pagination`, func() {
			listRulesOptions := &securityandcompliancecenterv3.ListRulesOptions{
				InstanceID: &instanceIDForLink,
				Limit:      core.Int64Ptr(int64(10)),
				Start:      core.StringPtr("testString"),
				Type:       core.StringPtr("system_defined"),
				Sort:       core.StringPtr("updated_on"),
			}

			listRulesOptions.Start = nil
			listRulesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterv3.Rule
			for i := 0; i < 3; i++ {
				ruleCollection, response, err := securityAndComplianceCenterService.ListRules(listRulesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(ruleCollection).ToNot(BeNil())
				allResults = append(allResults, ruleCollection.Rules...)

				listRulesOptions.Start, err = ruleCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listRulesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) using RulesPager`, func() {
			listRulesOptions := &securityandcompliancecenterv3.ListRulesOptions{
				InstanceID: &instanceIDForLink,
				Limit:      core.Int64Ptr(int64(10)),
				Type:       core.StringPtr("system_defined"),
				Sort:       core.StringPtr("updated_on"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterService.NewRulesPager(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterv3.Rule
			for i := 0; i < 3 && pager.HasNext(); i++ {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			fmt.Fprintf(GinkgoWriter, "ListRules() returned a total of %d item(s) using RulesPager.\n", len(allResults))
		})
	})

	Describe(`ListServices - List services`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListServices(listServicesOptions *ListServicesOptions)`, func() {
			listServicesOptions := &securityandcompliancecenterv3.ListServicesOptions{}

			serviceCollection, response, err := securityAndComplianceCenterService.ListServices(listServicesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceCollection).ToNot(BeNil())
		})
	})

	Describe(`GetService - Get a service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetService(getServiceOptions *GetServiceOptions)`, func() {
			getServiceOptions := &securityandcompliancecenterv3.GetServiceOptions{
				ServicesName: core.StringPtr("cloud-object-storage"),
			}

			service, response, err := securityAndComplianceCenterService.GetService(getServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(service).ToNot(BeNil())
		})
	})

	Describe(`DeleteProfileAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions)`, func() {
			deleteProfileAttachmentOptions := &securityandcompliancecenterv3.DeleteProfileAttachmentOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:    &profileIDLink,
				AttachmentID: &attachmentIDLink,
				AccountID:    &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomControlLibrary - Delete a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomControlLibrary(deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions)`, func() {
			deleteCustomControlLibraryOptions := &securityandcompliancecenterv3.DeleteCustomControlLibraryOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryID: &controlLibraryIDLink,
				AccountID:        &accountIDForReportLink,
			}

			controlLibrary, response, err := securityAndComplianceCenterService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomProfile - Delete a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions)`, func() {
			deleteCustomProfileOptions := &securityandcompliancecenterv3.DeleteCustomProfileOptions{
				InstanceID: &instanceIDForLink,
				ProfileID:  &profileIDLink,
				AccountID:  &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterService.DeleteCustomProfile(deleteCustomProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`DeleteSubscope - Delete a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSubscope(deleteSubscopeOptions *DeleteSubscopeOptions)`, func() {
			deleteSubscopeOptions := &securityandcompliancecenterv3.DeleteSubscopeOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
				SubscopeID: &subScopeIDLink,
			}

			response, err := securityAndComplianceCenterService.DeleteSubscope(deleteSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteScope - Delete a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteScope(deleteScopeOptions *DeleteScopeOptions)`, func() {
			deleteScopeOptions := &securityandcompliancecenterv3.DeleteScopeOptions{
				InstanceID: &instanceIDForLink,
				ScopeID:    &scopeIDLink,
			}

			response, err := securityAndComplianceCenterService.DeleteScope(deleteScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTarget - Delete a target by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
			deleteTargetOptions := &securityandcompliancecenterv3.DeleteTargetOptions{
				InstanceID: &instanceIDForLink,
				TargetID:   &targetIDLink,
			}

			response, err := securityAndComplianceCenterService.DeleteTarget(deleteTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteProviderTypeInstance - Delete a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions *DeleteProviderTypeInstanceOptions)`, func() {
			deleteProviderTypeInstanceOptions := &securityandcompliancecenterv3.DeleteProviderTypeInstanceOptions{
				InstanceID:             core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID:         &providerTypeIDLink,
				ProviderTypeInstanceID: &providerTypeInstanceIDLink,
			}

			response, err := securityAndComplianceCenterService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteRule - Delete a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
			deleteRuleOptions := &securityandcompliancecenterv3.DeleteRuleOptions{
				InstanceID: &instanceIDForLink,
				RuleID:     &ruleIDLink,
			}

			response, err := securityAndComplianceCenterService.DeleteRule(deleteRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
