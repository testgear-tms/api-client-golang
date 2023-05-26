# Api client for TestGear TMS

## Getting Started

### Compatibility

| TestGear | API Client |
|---------|------------|
| 4.2     | 1.0.0      |

### Installation

```sh
go get github.com/testgear-tms/api-client-golang
```

## Examples

Please follow the [installation](#installation) instruction and execute the following Java code:

```go
import (
	"context"
	"fmt"

	tmsclient "github.com/testgear-tms/api-client-golang"
)

const (
	host   = "Your TMS address"
	key    = "Your private token"
	scheme = "https or http"
)

func main() {
	ctx := context.WithValue(context.Background(), tmsclient.ContextAPIKeys, map[string]tmsclient.APIKey{
		"Bearer or PrivateToken": {
			Key:    key,
			Prefix: "PrivateToken",
		},
	})

	configuration := tmsclient.NewConfiguration()
	configuration.Host = host
	configuration.Scheme = scheme
	apiClient := tmsclient.NewAPIClient(configuration)

	id := "123"

	_, err := apiClient.AttachmentsApi.ApiV2AttachmentsIdDelete(ctx, id).Execute()
	if err != nil {
		fmt.Println(err)
	}
}
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AttachmentsApi* | [**ApiV2AttachmentsIdDelete**](docs/AttachmentsApi.md#apiv2attachmentsiddelete) | **Delete** /api/v2/attachments/{id} | Delete attachment file
*AttachmentsApi* | [**ApiV2AttachmentsIdGet**](docs/AttachmentsApi.md#apiv2attachmentsidget) | **Get** /api/v2/attachments/{id} | Download attachment file
*AttachmentsApi* | [**ApiV2AttachmentsOccupiedFileStorageSizeGet**](docs/AttachmentsApi.md#apiv2attachmentsoccupiedfilestoragesizeget) | **Get** /api/v2/attachments/occupiedFileStorageSize | Get size of attachments storage in bytes
*AttachmentsApi* | [**ApiV2AttachmentsPost**](docs/AttachmentsApi.md#apiv2attachmentspost) | **Post** /api/v2/attachments | Upload new attachment file
*AutoTestsApi* | [**ApiV2AutoTestsFlakyBulkPost**](docs/AutoTestsApi.md#apiv2autotestsflakybulkpost) | **Post** /api/v2/autoTests/flaky/bulk | Set \&quot;Flaky\&quot; status for multiple autotests
*AutoTestsApi* | [**ApiV2AutoTestsIdPatch**](docs/AutoTestsApi.md#apiv2autotestsidpatch) | **Patch** /api/v2/autoTests/{id} | Patch auto test
*AutoTestsApi* | [**ApiV2AutoTestsIdTestResultsSearchPost**](docs/AutoTestsApi.md#apiv2autotestsidtestresultssearchpost) | **Post** /api/v2/autoTests/{id}/testResults/search | Get test results history for autotest
*AutoTestsApi* | [**ApiV2AutoTestsIdWorkItemsChangedIdGet**](docs/AutoTestsApi.md#apiv2autotestsidworkitemschangedidget) | **Get** /api/v2/autoTests/{id}/workItems/changed/id | Get identifiers of changed linked work items
*AutoTestsApi* | [**ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost**](docs/AutoTestsApi.md#apiv2autotestsidworkitemschangedworkitemidapprovepost) | **Post** /api/v2/autoTests/{id}/workItems/changed/{workItemId}/approve | Approve changes to work items linked to autotest
*AutoTestsApi* | [**ApiV2AutoTestsSearchPost**](docs/AutoTestsApi.md#apiv2autotestssearchpost) | **Post** /api/v2/autoTests/search | Search for autotests
*AutoTestsApi* | [**CreateAutoTest**](docs/AutoTestsApi.md#createautotest) | **Post** /api/v2/autoTests | Create autotest
*AutoTestsApi* | [**CreateMultiple**](docs/AutoTestsApi.md#createmultiple) | **Post** /api/v2/autoTests/bulk | Create multiple autotests
*AutoTestsApi* | [**DeleteAutoTest**](docs/AutoTestsApi.md#deleteautotest) | **Delete** /api/v2/autoTests/{id} | Delete autotest
*AutoTestsApi* | [**DeleteAutoTestLinkFromWorkItem**](docs/AutoTestsApi.md#deleteautotestlinkfromworkitem) | **Delete** /api/v2/autoTests/{id}/workItems | Unlink autotest from work item
*AutoTestsApi* | [**GetAllAutoTests**](docs/AutoTestsApi.md#getallautotests) | **Get** /api/v2/autoTests | 
*AutoTestsApi* | [**GetAutoTestAverageDuration**](docs/AutoTestsApi.md#getautotestaverageduration) | **Get** /api/v2/autoTests/{id}/averageDuration | Get average autotest duration
*AutoTestsApi* | [**GetAutoTestById**](docs/AutoTestsApi.md#getautotestbyid) | **Get** /api/v2/autoTests/{id} | Get autotest by internal or global ID
*AutoTestsApi* | [**GetAutoTestChronology**](docs/AutoTestsApi.md#getautotestchronology) | **Get** /api/v2/autoTests/{id}/chronology | Get autotest chronology
*AutoTestsApi* | [**GetTestRuns**](docs/AutoTestsApi.md#gettestruns) | **Get** /api/v2/autoTests/{id}/testRuns | Get completed tests runs for autotests
*AutoTestsApi* | [**GetWorkItemResults**](docs/AutoTestsApi.md#getworkitemresults) | **Get** /api/v2/autoTests/{id}/testResultHistory | 
*AutoTestsApi* | [**GetWorkItemsLinkedToAutoTest**](docs/AutoTestsApi.md#getworkitemslinkedtoautotest) | **Get** /api/v2/autoTests/{id}/workItems | Get work items linked to autotest
*AutoTestsApi* | [**LinkAutoTestToWorkItem**](docs/AutoTestsApi.md#linkautotesttoworkitem) | **Post** /api/v2/autoTests/{id}/workItems | Link autotest with work items
*AutoTestsApi* | [**UpdateAutoTest**](docs/AutoTestsApi.md#updateautotest) | **Put** /api/v2/autoTests | Update autotest
*AutoTestsApi* | [**UpdateMultiple**](docs/AutoTestsApi.md#updatemultiple) | **Put** /api/v2/autoTests/bulk | Update multiple autotests
*BackgroundJobsApi* | [**ApiV2BackgroundJobsGet**](docs/BackgroundJobsApi.md#apiv2backgroundjobsget) | **Get** /api/v2/backgroundJobs | Get current user background jobs
*ConfigurationsApi* | [**ApiV2ConfigurationsCreateByParametersPost**](docs/ConfigurationsApi.md#apiv2configurationscreatebyparameterspost) | **Post** /api/v2/configurations/createByParameters | Create Configurations by parameters
*ConfigurationsApi* | [**ApiV2ConfigurationsIdPatch**](docs/ConfigurationsApi.md#apiv2configurationsidpatch) | **Patch** /api/v2/configurations/{id} | Patch configuration
*ConfigurationsApi* | [**ApiV2ConfigurationsSearchPost**](docs/ConfigurationsApi.md#apiv2configurationssearchpost) | **Post** /api/v2/configurations/search | Search for configurations
*ConfigurationsApi* | [**CreateConfiguration**](docs/ConfigurationsApi.md#createconfiguration) | **Post** /api/v2/configurations | Create Configuration
*ConfigurationsApi* | [**GetConfigurationById**](docs/ConfigurationsApi.md#getconfigurationbyid) | **Get** /api/v2/configurations/{id} | Get configuration by internal or global ID
*ConfigurationsApi* | [**UpdateConfiguration**](docs/ConfigurationsApi.md#updateconfiguration) | **Put** /api/v2/configurations | Update Configuration
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatesidcustomattributesexcludepost) | **Post** /api/v2/customAttributes/templates/{id}/customAttributes/exclude | Exclude CustomAttributes from CustomAttributeTemplate
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatesidcustomattributesincludepost) | **Post** /api/v2/customAttributes/templates/{id}/customAttributes/include | Include CustomAttributes to CustomAttributeTemplate
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesIdDelete**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatesiddelete) | **Delete** /api/v2/customAttributes/templates/{id} | Delete CustomAttributeTemplate
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesIdGet**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatesidget) | **Get** /api/v2/customAttributes/templates/{id} | Get CustomAttributeTemplate by ID
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesNameGet**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatesnameget) | **Get** /api/v2/customAttributes/templates/{name} | Get CustomAttributeTemplate by name
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesPost**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatespost) | **Post** /api/v2/customAttributes/templates | Create CustomAttributeTemplate
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesPut**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatesput) | **Put** /api/v2/customAttributes/templates | Update custom attributes template
*CustomAttributeTemplatesApi* | [**ApiV2CustomAttributesTemplatesSearchPost**](docs/CustomAttributeTemplatesApi.md#apiv2customattributestemplatessearchpost) | **Post** /api/v2/customAttributes/templates/search | Search CustomAttributeTemplates
*CustomAttributesApi* | [**ApiV2CustomAttributesGlobalIdDelete**](docs/CustomAttributesApi.md#apiv2customattributesglobaliddelete) | **Delete** /api/v2/customAttributes/global/{id} | Delete global attribute
*CustomAttributesApi* | [**ApiV2CustomAttributesGlobalIdPut**](docs/CustomAttributesApi.md#apiv2customattributesglobalidput) | **Put** /api/v2/customAttributes/global/{id} | Edit global attribute
*CustomAttributesApi* | [**ApiV2CustomAttributesGlobalPost**](docs/CustomAttributesApi.md#apiv2customattributesglobalpost) | **Post** /api/v2/customAttributes/global | Create global attribute
*CustomAttributesApi* | [**ApiV2CustomAttributesIdGet**](docs/CustomAttributesApi.md#apiv2customattributesidget) | **Get** /api/v2/customAttributes/{id} | Get attribute
*CustomAttributesApi* | [**ApiV2CustomAttributesSearchPost**](docs/CustomAttributesApi.md#apiv2customattributessearchpost) | **Post** /api/v2/customAttributes/search | Search for attributes
*NotificationsApi* | [**ApiV2NotificationsCountGet**](docs/NotificationsApi.md#apiv2notificationscountget) | **Get** /api/v2/notifications/count | Get unread Notifications total in last 7 days
*NotificationsApi* | [**ApiV2NotificationsGet**](docs/NotificationsApi.md#apiv2notificationsget) | **Get** /api/v2/notifications | Get all Notifications for current User
*NotificationsApi* | [**ApiV2NotificationsIdReadPost**](docs/NotificationsApi.md#apiv2notificationsidreadpost) | **Post** /api/v2/notifications/{id}/read | Set Notification as read
*NotificationsApi* | [**ApiV2NotificationsReadPost**](docs/NotificationsApi.md#apiv2notificationsreadpost) | **Post** /api/v2/notifications/read | Set all Notifications as read
*NotificationsApi* | [**ApiV2NotificationsSearchPost**](docs/NotificationsApi.md#apiv2notificationssearchpost) | **Post** /api/v2/notifications/search | Search Notifications for current User
*ParametersApi* | [**ApiV2ParametersBulkPost**](docs/ParametersApi.md#apiv2parametersbulkpost) | **Post** /api/v2/parameters/bulk | Create multiple parameters
*ParametersApi* | [**ApiV2ParametersBulkPut**](docs/ParametersApi.md#apiv2parametersbulkput) | **Put** /api/v2/parameters/bulk | Update multiple parameters
*ParametersApi* | [**ApiV2ParametersGroupsGet**](docs/ParametersApi.md#apiv2parametersgroupsget) | **Get** /api/v2/parameters/groups | Get parameters as group
*ParametersApi* | [**ApiV2ParametersKeyNameNameExistsGet**](docs/ParametersApi.md#apiv2parameterskeynamenameexistsget) | **Get** /api/v2/parameters/key/name/{name}/exists | Check existence parameter key in system
*ParametersApi* | [**ApiV2ParametersKeyValuesGet**](docs/ParametersApi.md#apiv2parameterskeyvaluesget) | **Get** /api/v2/parameters/{key}/values | Get all parameter key values
*ParametersApi* | [**ApiV2ParametersKeysGet**](docs/ParametersApi.md#apiv2parameterskeysget) | **Get** /api/v2/parameters/keys | Get all parameter keys
*ParametersApi* | [**ApiV2ParametersSearchPost**](docs/ParametersApi.md#apiv2parameterssearchpost) | **Post** /api/v2/parameters/search | Search for parameters
*ParametersApi* | [**CreateParameter**](docs/ParametersApi.md#createparameter) | **Post** /api/v2/parameters | Create parameter
*ParametersApi* | [**DeleteByName**](docs/ParametersApi.md#deletebyname) | **Delete** /api/v2/parameters/name/{name} | Delete parameter by name
*ParametersApi* | [**DeleteByParameterKeyId**](docs/ParametersApi.md#deletebyparameterkeyid) | **Delete** /api/v2/parameters/keyId/{keyId} | Delete parameters by parameter key identifier
*ParametersApi* | [**DeleteParameter**](docs/ParametersApi.md#deleteparameter) | **Delete** /api/v2/parameters/{id} | Delete parameter
*ParametersApi* | [**GetAllParameters**](docs/ParametersApi.md#getallparameters) | **Get** /api/v2/parameters | Get all parameters
*ParametersApi* | [**GetParameterById**](docs/ParametersApi.md#getparameterbyid) | **Get** /api/v2/parameters/{id} | Get parameter by ID
*ParametersApi* | [**ObsoleteDeleteByName**](docs/ParametersApi.md#obsoletedeletebyname) | **Post** /api/v2/parameters/deleteByName | 
*ParametersApi* | [**UpdateParameter**](docs/ParametersApi.md#updateparameter) | **Put** /api/v2/parameters | Update parameter
*ProjectsApi* | [**AddGlobaAttributesToProject**](docs/ProjectsApi.md#addglobaattributestoproject) | **Post** /api/v2/projects/{id}/globalAttributes | Add global attributes to project
*ProjectsApi* | [**ApiV2ProjectsIdAttributesTemplatesSearchPost**](docs/ProjectsApi.md#apiv2projectsidattributestemplatessearchpost) | **Post** /api/v2/projects/{id}/attributes/templates/search | Search for custom attributes templates
*ProjectsApi* | [**ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete**](docs/ProjectsApi.md#apiv2projectsidattributestemplatestemplateiddelete) | **Delete** /api/v2/projects/{id}/attributes/templates/{templateId} | Delete CustomAttributeTemplate from Project
*ProjectsApi* | [**ApiV2ProjectsIdAttributesTemplatesTemplateIdPost**](docs/ProjectsApi.md#apiv2projectsidattributestemplatestemplateidpost) | **Post** /api/v2/projects/{id}/attributes/templates/{templateId} | Add CustomAttributeTemplate to Project
*ProjectsApi* | [**ApiV2ProjectsIdFailureClassesGet**](docs/ProjectsApi.md#apiv2projectsidfailureclassesget) | **Get** /api/v2/projects/{id}/failureClasses | Get Project FailureClasses
*ProjectsApi* | [**ApiV2ProjectsIdFavoritePut**](docs/ProjectsApi.md#apiv2projectsidfavoriteput) | **Put** /api/v2/projects/{id}/favorite | Mark Project as favorite
*ProjectsApi* | [**ApiV2ProjectsIdFiltersGet**](docs/ProjectsApi.md#apiv2projectsidfiltersget) | **Get** /api/v2/projects/{id}/filters | Get Project filters
*ProjectsApi* | [**ApiV2ProjectsIdPatch**](docs/ProjectsApi.md#apiv2projectsidpatch) | **Patch** /api/v2/projects/{id} | Patch project
*ProjectsApi* | [**ApiV2ProjectsIdTestPlansAnalyticsGet**](docs/ProjectsApi.md#apiv2projectsidtestplansanalyticsget) | **Get** /api/v2/projects/{id}/testPlans/analytics | Get TestPlans analytics
*ProjectsApi* | [**ApiV2ProjectsIdTestPlansDeleteBulkPost**](docs/ProjectsApi.md#apiv2projectsidtestplansdeletebulkpost) | **Post** /api/v2/projects/{id}/testPlans/delete/bulk | Delete multiple test plans
*ProjectsApi* | [**ApiV2ProjectsIdTestPlansNameExistsGet**](docs/ProjectsApi.md#apiv2projectsidtestplansnameexistsget) | **Get** /api/v2/projects/{id}/testPlans/{name}/exists | Checks if TestPlan exists with the specified name exists for the project
*ProjectsApi* | [**ApiV2ProjectsIdTestPlansRestoreBulkPost**](docs/ProjectsApi.md#apiv2projectsidtestplansrestorebulkpost) | **Post** /api/v2/projects/{id}/testPlans/restore/bulk | Restore multiple test plans
*ProjectsApi* | [**ApiV2ProjectsIdTestPlansSearchPost**](docs/ProjectsApi.md#apiv2projectsidtestplanssearchpost) | **Post** /api/v2/projects/{id}/testPlans/search | Get Project TestPlans with analytics
*ProjectsApi* | [**ApiV2ProjectsIdTestRunsActiveGet**](docs/ProjectsApi.md#apiv2projectsidtestrunsactiveget) | **Get** /api/v2/projects/{id}/testRuns/active | Get active Project TestRuns
*ProjectsApi* | [**ApiV2ProjectsIdTestRunsFullGet**](docs/ProjectsApi.md#apiv2projectsidtestrunsfullget) | **Get** /api/v2/projects/{id}/testRuns/full | Get Project TestRuns full models
*ProjectsApi* | [**ApiV2ProjectsIdWorkItemsSearchIdPost**](docs/ProjectsApi.md#apiv2projectsidworkitemssearchidpost) | **Post** /api/v2/projects/{id}/workItems/search/id | Search for work items and extract IDs only
*ProjectsApi* | [**ApiV2ProjectsIdWorkItemsSearchPost**](docs/ProjectsApi.md#apiv2projectsidworkitemssearchpost) | **Post** /api/v2/projects/{id}/workItems/search | Search for work items
*ProjectsApi* | [**ApiV2ProjectsIdWorkItemsTagsGet**](docs/ProjectsApi.md#apiv2projectsidworkitemstagsget) | **Get** /api/v2/projects/{id}/workItems/tags | Get WorkItems Tags
*ProjectsApi* | [**ApiV2ProjectsNameNameExistsGet**](docs/ProjectsApi.md#apiv2projectsnamenameexistsget) | **Get** /api/v2/projects/name/{name}/exists | 
*ProjectsApi* | [**ApiV2ProjectsSearchPost**](docs/ProjectsApi.md#apiv2projectssearchpost) | **Post** /api/v2/projects/search | Search for projects
*ProjectsApi* | [**BackgroundImportProject**](docs/ProjectsApi.md#backgroundimportproject) | **Post** /api/v2/projects/import/json | Import project from JSON file in background job
*ProjectsApi* | [**BackgroundImportToExistingProject**](docs/ProjectsApi.md#backgroundimporttoexistingproject) | **Post** /api/v2/projects/{id}/import/json | Import project from JSON file into existing project in background job
*ProjectsApi* | [**BackgroundImportZipProject**](docs/ProjectsApi.md#backgroundimportzipproject) | **Post** /api/v2/projects/import/zip | Import project from Zip file in background job
*ProjectsApi* | [**BackgroundImportZipToExistingProject**](docs/ProjectsApi.md#backgroundimportziptoexistingproject) | **Post** /api/v2/projects/{id}/import/zip | Import project from Zip file into existing project in background job
*ProjectsApi* | [**CallImport**](docs/ProjectsApi.md#callimport) | **Post** /api/v2/projects/import | Import project from JSON file
*ProjectsApi* | [**CreateCustomAttributeTestPlanProjectRelations**](docs/ProjectsApi.md#createcustomattributetestplanprojectrelations) | **Post** /api/v2/projects/{id}/testPlans/attributes | Add attributes to project&#39;s test plans
*ProjectsApi* | [**CreateProject**](docs/ProjectsApi.md#createproject) | **Post** /api/v2/projects | Create project
*ProjectsApi* | [**CreateProjectsAttribute**](docs/ProjectsApi.md#createprojectsattribute) | **Post** /api/v2/projects/{id}/attributes | Create project attribute
*ProjectsApi* | [**DeleteCustomAttributeTestPlanProjectRelations**](docs/ProjectsApi.md#deletecustomattributetestplanprojectrelations) | **Delete** /api/v2/projects/{id}/testPlans/attribute/{attributeId} | Delete attribute from project&#39;s test plans
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /api/v2/projects/{id} | Delete project
*ProjectsApi* | [**DeleteProjectAutoTests**](docs/ProjectsApi.md#deleteprojectautotests) | **Delete** /api/v2/projects/{id}/autoTests | Delete project
*ProjectsApi* | [**DeleteProjectsAttribute**](docs/ProjectsApi.md#deleteprojectsattribute) | **Delete** /api/v2/projects/{id}/attributes/{attributeId} | Delete project attribute
*ProjectsApi* | [**Export**](docs/ProjectsApi.md#export) | **Post** /api/v2/projects/{id}/export | Export project as JSON file
*ProjectsApi* | [**ExportProjectJson**](docs/ProjectsApi.md#exportprojectjson) | **Post** /api/v2/projects/{id}/export/json | Export project as JSON file in background job
*ProjectsApi* | [**ExportProjectWithTestPlansJson**](docs/ProjectsApi.md#exportprojectwithtestplansjson) | **Post** /api/v2/projects/{id}/export/testPlans/json | Export project as JSON file with test plans in background job
*ProjectsApi* | [**ExportProjectWithTestPlansZip**](docs/ProjectsApi.md#exportprojectwithtestplanszip) | **Post** /api/v2/projects/{id}/export/testPlans/zip | Export project as Zip file with test plans in background job
*ProjectsApi* | [**ExportProjectZip**](docs/ProjectsApi.md#exportprojectzip) | **Post** /api/v2/projects/{id}/export/zip | Export project as Zip file in background job
*ProjectsApi* | [**ExportWithTestPlansAndConfigurations**](docs/ProjectsApi.md#exportwithtestplansandconfigurations) | **Post** /api/v2/projects/{id}/export-by-testPlans | Export project with test plans, test suites and test points as JSON file
*ProjectsApi* | [**GetAllProjects**](docs/ProjectsApi.md#getallprojects) | **Get** /api/v2/projects | Get all projects
*ProjectsApi* | [**GetAttributeByProjectId**](docs/ProjectsApi.md#getattributebyprojectid) | **Get** /api/v2/projects/{id}/attributes/{attributeId} | Get project attribute
*ProjectsApi* | [**GetAttributesByProjectId**](docs/ProjectsApi.md#getattributesbyprojectid) | **Get** /api/v2/projects/{id}/attributes | Get project attributes
*ProjectsApi* | [**GetAutoTestsNamespaces**](docs/ProjectsApi.md#getautotestsnamespaces) | **Get** /api/v2/projects/{id}/autoTestsNamespaces | Get namespaces of autotests in project
*ProjectsApi* | [**GetConfigurationsByProjectId**](docs/ProjectsApi.md#getconfigurationsbyprojectid) | **Get** /api/v2/projects/{id}/configurations | Get project configurations
*ProjectsApi* | [**GetCustomAttributeTestPlanProjectRelations**](docs/ProjectsApi.md#getcustomattributetestplanprojectrelations) | **Get** /api/v2/projects/{id}/testPlans/attributes | Get project&#39;s test plan attributes
*ProjectsApi* | [**GetProjectById**](docs/ProjectsApi.md#getprojectbyid) | **Get** /api/v2/projects/{id} | Get project by ID
*ProjectsApi* | [**GetSectionsByProjectId**](docs/ProjectsApi.md#getsectionsbyprojectid) | **Get** /api/v2/projects/{id}/sections | Get project sections
*ProjectsApi* | [**GetTestPlansByProjectId**](docs/ProjectsApi.md#gettestplansbyprojectid) | **Get** /api/v2/projects/{id}/testPlans | Get project test plans
*ProjectsApi* | [**GetTestRunsByProjectId**](docs/ProjectsApi.md#gettestrunsbyprojectid) | **Get** /api/v2/projects/{id}/testRuns | Get project test runs
*ProjectsApi* | [**GetWorkItemsByProjectId**](docs/ProjectsApi.md#getworkitemsbyprojectid) | **Get** /api/v2/projects/{id}/workItems | Get project work items
*ProjectsApi* | [**ImportToExistingProject**](docs/ProjectsApi.md#importtoexistingproject) | **Post** /api/v2/projects/{id}/import | Import project from JSON file into existing project
*ProjectsApi* | [**RestoreProject**](docs/ProjectsApi.md#restoreproject) | **Post** /api/v2/projects/{id}/restore | Restore project
*ProjectsApi* | [**SearchAttributesInProject**](docs/ProjectsApi.md#searchattributesinproject) | **Post** /api/v2/projects/{id}/attributes/search | Search for attributes used in the project
*ProjectsApi* | [**SearchTestPlanAttributesInProject**](docs/ProjectsApi.md#searchtestplanattributesinproject) | **Post** /api/v2/projects/{id}/testPlans/attributes/search | Search for attributes used in the project test plans
*ProjectsApi* | [**UpdateCustomAttributeTestPlanProjectRelations**](docs/ProjectsApi.md#updatecustomattributetestplanprojectrelations) | **Put** /api/v2/projects/{id}/testPlans/attribute | Update attribute of project&#39;s test plans
*ProjectsApi* | [**UpdateProject**](docs/ProjectsApi.md#updateproject) | **Put** /api/v2/projects | Update project
*ProjectsApi* | [**UpdateProjectsAttribute**](docs/ProjectsApi.md#updateprojectsattribute) | **Put** /api/v2/projects/{id}/attributes | Edit attribute of the project
*SectionsApi* | [**ApiV2SectionsIdPatch**](docs/SectionsApi.md#apiv2sectionsidpatch) | **Patch** /api/v2/sections/{id} | Patch section
*SectionsApi* | [**CreateSection**](docs/SectionsApi.md#createsection) | **Post** /api/v2/sections | Create section
*SectionsApi* | [**DeleteSection**](docs/SectionsApi.md#deletesection) | **Delete** /api/v2/sections/{id} | Delete section
*SectionsApi* | [**GetSectionById**](docs/SectionsApi.md#getsectionbyid) | **Get** /api/v2/sections/{id} | Get section
*SectionsApi* | [**GetWorkItemsBySectionId**](docs/SectionsApi.md#getworkitemsbysectionid) | **Get** /api/v2/sections/{id}/workItems | Get section work items
*SectionsApi* | [**Move**](docs/SectionsApi.md#move) | **Post** /api/v2/sections/move | Move section with all work items into another section
*SectionsApi* | [**Rename**](docs/SectionsApi.md#rename) | **Post** /api/v2/sections/rename | Rename section
*SectionsApi* | [**UpdateSection**](docs/SectionsApi.md#updatesection) | **Put** /api/v2/sections | Update section
*TagsApi* | [**ApiV2TagsGet**](docs/TagsApi.md#apiv2tagsget) | **Get** /api/v2/tags | Get all Tags
*TagsApi* | [**ApiV2TagsTestPlansTagsGet**](docs/TagsApi.md#apiv2tagstestplanstagsget) | **Get** /api/v2/tags/testPlansTags | Get all Tags that are used in TestPlans
*TestPlansApi* | [**AddTestPointsWithSections**](docs/TestPlansApi.md#addtestpointswithsections) | **Post** /api/v2/testPlans/{id}/test-points/withSections | Add test-points to TestPlan with sections
*TestPlansApi* | [**AddWorkItemsWithSections**](docs/TestPlansApi.md#addworkitemswithsections) | **Post** /api/v2/testPlans/{id}/workItems/withSections | Add WorkItems to TestPlan with Sections as TestSuites
*TestPlansApi* | [**ApiV2TestPlansIdAnalyticsGet**](docs/TestPlansApi.md#apiv2testplansidanalyticsget) | **Get** /api/v2/testPlans/{id}/analytics | Get analytics by TestPlan
*TestPlansApi* | [**ApiV2TestPlansIdAutobalancePost**](docs/TestPlansApi.md#apiv2testplansidautobalancepost) | **Post** /api/v2/testPlans/{id}/autobalance | Distribute test points between the users
*TestPlansApi* | [**ApiV2TestPlansIdConfigurationsGet**](docs/TestPlansApi.md#apiv2testplansidconfigurationsget) | **Get** /api/v2/testPlans/{id}/configurations | Get TestPlan configurations
*TestPlansApi* | [**ApiV2TestPlansIdExportTestPointsXlsxPost**](docs/TestPlansApi.md#apiv2testplansidexporttestpointsxlsxpost) | **Post** /api/v2/testPlans/{id}/export/testPoints/xlsx | Export TestPoints from TestPlan in xls format
*TestPlansApi* | [**ApiV2TestPlansIdExportTestResultHistoryXlsxPost**](docs/TestPlansApi.md#apiv2testplansidexporttestresulthistoryxlsxpost) | **Post** /api/v2/testPlans/{id}/export/testResultHistory/xlsx | Export TestResults history from TestPlan in xls format
*TestPlansApi* | [**ApiV2TestPlansIdHistoryGet**](docs/TestPlansApi.md#apiv2testplansidhistoryget) | **Get** /api/v2/testPlans/{id}/history | Get TestPlan history
*TestPlansApi* | [**ApiV2TestPlansIdLinksGet**](docs/TestPlansApi.md#apiv2testplansidlinksget) | **Get** /api/v2/testPlans/{id}/links | Get Links of TestPlan
*TestPlansApi* | [**ApiV2TestPlansIdPatch**](docs/TestPlansApi.md#apiv2testplansidpatch) | **Patch** /api/v2/testPlans/{id} | Patch test plan
*TestPlansApi* | [**ApiV2TestPlansIdTestPointsLastResultsGet**](docs/TestPlansApi.md#apiv2testplansidtestpointslastresultsget) | **Get** /api/v2/testPlans/{id}/testPoints/lastResults | Get TestPoints with last result from TestPlan
*TestPlansApi* | [**ApiV2TestPlansIdTestPointsResetPost**](docs/TestPlansApi.md#apiv2testplansidtestpointsresetpost) | **Post** /api/v2/testPlans/{id}/testPoints/reset | Reset TestPoints status of TestPlan
*TestPlansApi* | [**ApiV2TestPlansIdTestPointsTesterDelete**](docs/TestPlansApi.md#apiv2testplansidtestpointstesterdelete) | **Delete** /api/v2/testPlans/{id}/testPoints/tester | Unassign users from multiple test points
*TestPlansApi* | [**ApiV2TestPlansIdTestPointsTesterUserIdPost**](docs/TestPlansApi.md#apiv2testplansidtestpointstesteruseridpost) | **Post** /api/v2/testPlans/{id}/testPoints/tester/{userId} | Assign user as a tester to multiple test points
*TestPlansApi* | [**ApiV2TestPlansIdTestRunsGet**](docs/TestPlansApi.md#apiv2testplansidtestrunsget) | **Get** /api/v2/testPlans/{id}/testRuns | Get TestRuns of TestPlan
*TestPlansApi* | [**ApiV2TestPlansIdTestRunsSearchPost**](docs/TestPlansApi.md#apiv2testplansidtestrunssearchpost) | **Post** /api/v2/testPlans/{id}/testRuns/search | Search TestRuns of TestPlan
*TestPlansApi* | [**ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet**](docs/TestPlansApi.md#apiv2testplansidtestrunstestresultslastmodifiedmodifieddateget) | **Get** /api/v2/testPlans/{id}/testRuns/testResults/lastModified/modifiedDate | Get last modification date of test plan&#39;s test results
*TestPlansApi* | [**ApiV2TestPlansIdUnlockRequestPost**](docs/TestPlansApi.md#apiv2testplansidunlockrequestpost) | **Post** /api/v2/testPlans/{id}/unlock/request | Send unlock TestPlan notification
*TestPlansApi* | [**ApiV2TestPlansShortsPost**](docs/TestPlansApi.md#apiv2testplansshortspost) | **Post** /api/v2/testPlans/shorts | Get TestPlans short models by Project identifiers
*TestPlansApi* | [**Clone**](docs/TestPlansApi.md#clone) | **Post** /api/v2/testPlans/{id}/clone | Clone TestPlan
*TestPlansApi* | [**Complete**](docs/TestPlansApi.md#complete) | **Post** /api/v2/testPlans/{id}/complete | Complete TestPlan
*TestPlansApi* | [**CreateTestPlan**](docs/TestPlansApi.md#createtestplan) | **Post** /api/v2/testPlans | Create TestPlan
*TestPlansApi* | [**DeleteTestPlan**](docs/TestPlansApi.md#deletetestplan) | **Delete** /api/v2/testPlans/{id} | Delete TestPlan
*TestPlansApi* | [**GetTestPlanById**](docs/TestPlansApi.md#gettestplanbyid) | **Get** /api/v2/testPlans/{id} | Get TestPlan by Id
*TestPlansApi* | [**GetTestSuitesById**](docs/TestPlansApi.md#gettestsuitesbyid) | **Get** /api/v2/testPlans/{id}/testSuites | Get TestSuites Tree By Id
*TestPlansApi* | [**Pause**](docs/TestPlansApi.md#pause) | **Post** /api/v2/testPlans/{id}/pause | Pause TestPlan
*TestPlansApi* | [**RestoreTestPlan**](docs/TestPlansApi.md#restoretestplan) | **Post** /api/v2/testPlans/{id}/restore | Restore TestPlan
*TestPlansApi* | [**Start**](docs/TestPlansApi.md#start) | **Post** /api/v2/testPlans/{id}/start | Start TestPlan
*TestPlansApi* | [**UpdateTestPlan**](docs/TestPlansApi.md#updatetestplan) | **Put** /api/v2/testPlans | Update TestPlan
*TestPointsApi* | [**ApiV2TestPointsIdTestRunsGet**](docs/TestPointsApi.md#apiv2testpointsidtestrunsget) | **Get** /api/v2/testPoints/{id}/testRuns | Get all test runs which use test point
*TestPointsApi* | [**ApiV2TestPointsIdWorkItemGet**](docs/TestPointsApi.md#apiv2testpointsidworkitemget) | **Get** /api/v2/testPoints/{id}/workItem | Get work item represented by test point
*TestPointsApi* | [**ApiV2TestPointsSearchIdPost**](docs/TestPointsApi.md#apiv2testpointssearchidpost) | **Post** /api/v2/testPoints/search/id | Search for test points and extract IDs only
*TestPointsApi* | [**ApiV2TestPointsSearchPost**](docs/TestPointsApi.md#apiv2testpointssearchpost) | **Post** /api/v2/testPoints/search | Search for test points
*TestResultsApi* | [**ApiV2TestResultsIdAggregatedGet**](docs/TestResultsApi.md#apiv2testresultsidaggregatedget) | **Get** /api/v2/testResults/{id}/aggregated | Get test result by ID aggregated with previous results
*TestResultsApi* | [**ApiV2TestResultsIdAttachmentsAttachmentIdPut**](docs/TestResultsApi.md#apiv2testresultsidattachmentsattachmentidput) | **Put** /api/v2/testResults/{id}/attachments/{attachmentId} | Attach file to the test result
*TestResultsApi* | [**ApiV2TestResultsIdAttachmentsInfoGet**](docs/TestResultsApi.md#apiv2testresultsidattachmentsinfoget) | **Get** /api/v2/testResults/{id}/attachments/info | Get test result attachments meta-information
*TestResultsApi* | [**ApiV2TestResultsIdGet**](docs/TestResultsApi.md#apiv2testresultsidget) | **Get** /api/v2/testResults/{id} | Get test result by ID
*TestResultsApi* | [**ApiV2TestResultsIdPut**](docs/TestResultsApi.md#apiv2testresultsidput) | **Put** /api/v2/testResults/{id} | Edit test result by ID
*TestResultsApi* | [**ApiV2TestResultsSearchPost**](docs/TestResultsApi.md#apiv2testresultssearchpost) | **Post** /api/v2/testResults/search | Search for test results
*TestResultsApi* | [**ApiV2TestResultsStatisticsFilterPost**](docs/TestResultsApi.md#apiv2testresultsstatisticsfilterpost) | **Post** /api/v2/testResults/statistics/filter | Search for test results and extract statistics
*TestResultsApi* | [**CreateAttachment**](docs/TestResultsApi.md#createattachment) | **Post** /api/v2/testResults/{id}/attachments | Upload and link attachment to TestResult
*TestResultsApi* | [**DeleteAttachment**](docs/TestResultsApi.md#deleteattachment) | **Delete** /api/v2/testResults/{id}/attachments/{attachmentId} | Remove attachment and unlink from TestResult
*TestResultsApi* | [**DownloadAttachment**](docs/TestResultsApi.md#downloadattachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId} | Get attachment of TestResult
*TestResultsApi* | [**GetAttachment**](docs/TestResultsApi.md#getattachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId}/info | Get Metadata of TestResult&#39;s attachment
*TestResultsApi* | [**GetAttachments**](docs/TestResultsApi.md#getattachments) | **Get** /api/v2/testResults/{id}/attachments | Get all attachments of TestResult
*TestRunsApi* | [**ApiV2TestRunsIdStatisticsFilterPost**](docs/TestRunsApi.md#apiv2testrunsidstatisticsfilterpost) | **Post** /api/v2/testRuns/{id}/statistics/filter | Search for the test run test results and build statistics
*TestRunsApi* | [**ApiV2TestRunsIdTestPointsResultsGet**](docs/TestRunsApi.md#apiv2testrunsidtestpointsresultsget) | **Get** /api/v2/testRuns/{id}/testPoints/results | Get test results from the test run grouped by test points
*TestRunsApi* | [**ApiV2TestRunsIdTestResultsBulkPut**](docs/TestRunsApi.md#apiv2testrunsidtestresultsbulkput) | **Put** /api/v2/testRuns/{id}/testResults/bulk | Partial edit of multiple test results in the test run
*TestRunsApi* | [**ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet**](docs/TestRunsApi.md#apiv2testrunsidtestresultslastmodifiedmodificationdateget) | **Get** /api/v2/testRuns/{id}/testResults/lastModified/modificationDate | Get modification date of last test result of the test run
*TestRunsApi* | [**ApiV2TestRunsSearchPost**](docs/TestRunsApi.md#apiv2testrunssearchpost) | **Post** /api/v2/testRuns/search | Search for test runs
*TestRunsApi* | [**CompleteTestRun**](docs/TestRunsApi.md#completetestrun) | **Post** /api/v2/testRuns/{id}/complete | Complete TestRun
*TestRunsApi* | [**CreateAndFillByAutoTests**](docs/TestRunsApi.md#createandfillbyautotests) | **Post** /api/v2/testRuns/byAutoTests | Create test runs based on autotests and configurations
*TestRunsApi* | [**CreateAndFillByConfigurations**](docs/TestRunsApi.md#createandfillbyconfigurations) | **Post** /api/v2/testRuns/byConfigurations | Create test runs picking the needed test points
*TestRunsApi* | [**CreateAndFillByWorkItems**](docs/TestRunsApi.md#createandfillbyworkitems) | **Post** /api/v2/testRuns/byWorkItems | Create test run based on configurations and work items
*TestRunsApi* | [**CreateEmpty**](docs/TestRunsApi.md#createempty) | **Post** /api/v2/testRuns | Create empty TestRun
*TestRunsApi* | [**GetTestRunById**](docs/TestRunsApi.md#gettestrunbyid) | **Get** /api/v2/testRuns/{id} | Get TestRun by Id
*TestRunsApi* | [**SetAutoTestResultsForTestRun**](docs/TestRunsApi.md#setautotestresultsfortestrun) | **Post** /api/v2/testRuns/{id}/testResults | Send test results to the test runs in the system
*TestRunsApi* | [**StartTestRun**](docs/TestRunsApi.md#starttestrun) | **Post** /api/v2/testRuns/{id}/start | Start TestRun
*TestRunsApi* | [**StopTestRun**](docs/TestRunsApi.md#stoptestrun) | **Post** /api/v2/testRuns/{id}/stop | Stop TestRun
*TestRunsApi* | [**UpdateEmpty**](docs/TestRunsApi.md#updateempty) | **Put** /api/v2/testRuns | Update empty TestRun
*TestSuitesApi* | [**AddTestPointsToTestSuite**](docs/TestSuitesApi.md#addtestpointstotestsuite) | **Post** /api/v2/testSuites/{id}/test-points | Add test-points to test suite
*TestSuitesApi* | [**ApiV2TestSuitesIdPatch**](docs/TestSuitesApi.md#apiv2testsuitesidpatch) | **Patch** /api/v2/testSuites/{id} | Patch test suite
*TestSuitesApi* | [**ApiV2TestSuitesIdRefreshPost**](docs/TestSuitesApi.md#apiv2testsuitesidrefreshpost) | **Post** /api/v2/testSuites/{id}/refresh | Refresh test suite. Only dynamic test suites are supported by this method
*TestSuitesApi* | [**CreateTestSuite**](docs/TestSuitesApi.md#createtestsuite) | **Post** /api/v2/testSuites | Create TestSuite
*TestSuitesApi* | [**DeleteTestSuite**](docs/TestSuitesApi.md#deletetestsuite) | **Delete** /api/v2/testSuites/{id} | Delete TestSuite
*TestSuitesApi* | [**GetConfigurationsByTestSuiteId**](docs/TestSuitesApi.md#getconfigurationsbytestsuiteid) | **Get** /api/v2/testSuites/{id}/configurations | Get Configurations By Id
*TestSuitesApi* | [**GetTestPointsById**](docs/TestSuitesApi.md#gettestpointsbyid) | **Get** /api/v2/testSuites/{id}/testPoints | Get TestPoints By Id
*TestSuitesApi* | [**GetTestResultsById**](docs/TestSuitesApi.md#gettestresultsbyid) | **Get** /api/v2/testSuites/{id}/testResults | Get TestResults By Id
*TestSuitesApi* | [**GetTestSuiteById**](docs/TestSuitesApi.md#gettestsuitebyid) | **Get** /api/v2/testSuites/{id} | Get TestSuite by Id
*TestSuitesApi* | [**GetWorkItemsById**](docs/TestSuitesApi.md#getworkitemsbyid) | **Get** /api/v2/testSuites/{id}/workItems | 
*TestSuitesApi* | [**SearchWorkItems**](docs/TestSuitesApi.md#searchworkitems) | **Post** /api/v2/testSuites/{id}/workItems/search | Search WorkItems
*TestSuitesApi* | [**SetConfigurationsByTestSuiteId**](docs/TestSuitesApi.md#setconfigurationsbytestsuiteid) | **Post** /api/v2/testSuites/{id}/configurations | Set Configurations By TestSuite Id
*TestSuitesApi* | [**SetWorkItemsByTestSuiteId**](docs/TestSuitesApi.md#setworkitemsbytestsuiteid) | **Post** /api/v2/testSuites/{id}/workItems | Set WorkItems By TestSuite Id
*TestSuitesApi* | [**UpdateTestSuite**](docs/TestSuitesApi.md#updatetestsuite) | **Put** /api/v2/testSuites | Update TestSuite
*WebhooksApi* | [**ApiV2WebhooksGet**](docs/WebhooksApi.md#apiv2webhooksget) | **Get** /api/v2/webhooks | Get all webhooks
*WebhooksApi* | [**ApiV2WebhooksIdDelete**](docs/WebhooksApi.md#apiv2webhooksiddelete) | **Delete** /api/v2/webhooks/{id} | Delete webhook by ID
*WebhooksApi* | [**ApiV2WebhooksIdGet**](docs/WebhooksApi.md#apiv2webhooksidget) | **Get** /api/v2/webhooks/{id} | Get webhook by ID
*WebhooksApi* | [**ApiV2WebhooksIdPut**](docs/WebhooksApi.md#apiv2webhooksidput) | **Put** /api/v2/webhooks/{id} | Edit webhook by ID
*WebhooksApi* | [**ApiV2WebhooksPost**](docs/WebhooksApi.md#apiv2webhookspost) | **Post** /api/v2/webhooks | Create webhook
*WebhooksApi* | [**ApiV2WebhooksSearchPost**](docs/WebhooksApi.md#apiv2webhookssearchpost) | **Post** /api/v2/webhooks/search | Search for webhooks
*WebhooksApi* | [**ApiV2WebhooksSpecialVariablesGet**](docs/WebhooksApi.md#apiv2webhooksspecialvariablesget) | **Get** /api/v2/webhooks/specialVariables | Get special variables for webhook event type
*WebhooksLogsApi* | [**ApiV2WebhooksLogsGet**](docs/WebhooksLogsApi.md#apiv2webhookslogsget) | **Get** /api/v2/webhooks/logs | Get all webhook logs
*WebhooksLogsApi* | [**ApiV2WebhooksLogsIdDelete**](docs/WebhooksLogsApi.md#apiv2webhookslogsiddelete) | **Delete** /api/v2/webhooks/logs/{id} | Delete webhook log by ID
*WebhooksLogsApi* | [**ApiV2WebhooksLogsIdGet**](docs/WebhooksLogsApi.md#apiv2webhookslogsidget) | **Get** /api/v2/webhooks/logs/{id} | Get webhook log by ID
*WorkItemsApi* | [**ApiV2WorkItemsIdAttachmentsPost**](docs/WorkItemsApi.md#apiv2workitemsidattachmentspost) | **Post** /api/v2/workItems/{id}/attachments | Upload and link attachment to WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdCheckListTransformToTestCasePost**](docs/WorkItemsApi.md#apiv2workitemsidchecklisttransformtotestcasepost) | **Post** /api/v2/workItems/{id}/checkList/transformTo/testCase | Transform CheckList to TestCase
*WorkItemsApi* | [**ApiV2WorkItemsIdHistoryGet**](docs/WorkItemsApi.md#apiv2workitemsidhistoryget) | **Get** /api/v2/workItems/{id}/history | Get change history of WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdLikeDelete**](docs/WorkItemsApi.md#apiv2workitemsidlikedelete) | **Delete** /api/v2/workItems/{id}/like | Delete like from WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdLikePost**](docs/WorkItemsApi.md#apiv2workitemsidlikepost) | **Post** /api/v2/workItems/{id}/like | Set like to WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdLikesCountGet**](docs/WorkItemsApi.md#apiv2workitemsidlikescountget) | **Get** /api/v2/workItems/{id}/likes/count | Get likes count of WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdLikesGet**](docs/WorkItemsApi.md#apiv2workitemsidlikesget) | **Get** /api/v2/workItems/{id}/likes | Get likes of WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdTestResultsHistoryGet**](docs/WorkItemsApi.md#apiv2workitemsidtestresultshistoryget) | **Get** /api/v2/workItems/{id}/testResults/history | Get test results history of WorkItem
*WorkItemsApi* | [**ApiV2WorkItemsIdVersionVersionIdActualPost**](docs/WorkItemsApi.md#apiv2workitemsidversionversionidactualpost) | **Post** /api/v2/workItems/{id}/version/{versionId}/actual | Set WorkItem as actual
*WorkItemsApi* | [**ApiV2WorkItemsMovePost**](docs/WorkItemsApi.md#apiv2workitemsmovepost) | **Post** /api/v2/workItems/move | Move WorkItem to another section
*WorkItemsApi* | [**ApiV2WorkItemsSearchPost**](docs/WorkItemsApi.md#apiv2workitemssearchpost) | **Post** /api/v2/workItems/search | Search for work items
*WorkItemsApi* | [**ApiV2WorkItemsSharedStepIdReferencesSectionsPost**](docs/WorkItemsApi.md#apiv2workitemssharedstepidreferencessectionspost) | **Post** /api/v2/workItems/{sharedStepId}/references/sections | Get SharedStep references in sections
*WorkItemsApi* | [**ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost**](docs/WorkItemsApi.md#apiv2workitemssharedstepidreferencesworkitemspost) | **Post** /api/v2/workItems/{sharedStepId}/references/workItems | Get SharedStep references in workitems
*WorkItemsApi* | [**ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet**](docs/WorkItemsApi.md#apiv2workitemssharedstepssharedstepidreferencesget) | **Get** /api/v2/workItems/sharedSteps/{sharedStepId}/references | Get SharedStep references
*WorkItemsApi* | [**CreateWorkItem**](docs/WorkItemsApi.md#createworkitem) | **Post** /api/v2/workItems | Create Test Case, Checklist or Shared Step
*WorkItemsApi* | [**DeleteAllWorkItemsFromAutoTest**](docs/WorkItemsApi.md#deleteallworkitemsfromautotest) | **Delete** /api/v2/workItems/{id}/autoTests | Delete all links AutoTests from WorkItem by Id or GlobalId
*WorkItemsApi* | [**DeleteWorkItem**](docs/WorkItemsApi.md#deleteworkitem) | **Delete** /api/v2/workItems/{id} | Delete Test Case, Checklist or Shared Step by Id or GlobalId
*WorkItemsApi* | [**GetAutoTestsForWorkItem**](docs/WorkItemsApi.md#getautotestsforworkitem) | **Get** /api/v2/workItems/{id}/autoTests | Get all AutoTests linked to WorkItem by Id or GlobalId
*WorkItemsApi* | [**GetIterations**](docs/WorkItemsApi.md#getiterations) | **Get** /api/v2/workItems/{id}/iterations | Get iterations by workitem Id or GlobalId
*WorkItemsApi* | [**GetWorkItemById**](docs/WorkItemsApi.md#getworkitembyid) | **Get** /api/v2/workItems/{id} | Get Test Case, Checklist or Shared Step by Id or GlobalId
*WorkItemsApi* | [**GetWorkItemChronology**](docs/WorkItemsApi.md#getworkitemchronology) | **Get** /api/v2/workItems/{id}/chronology | Get WorkItem chronology by Id or GlobalId
*WorkItemsApi* | [**GetWorkItemVersions**](docs/WorkItemsApi.md#getworkitemversions) | **Get** /api/v2/workItems/{id}/versions | Get WorkItem versions
*WorkItemsApi* | [**PurgeWorkItem**](docs/WorkItemsApi.md#purgeworkitem) | **Post** /api/v2/workItems/{id}/purge | Permanently delete test case, checklist or shared steps from archive
*WorkItemsApi* | [**RestoreWorkItem**](docs/WorkItemsApi.md#restoreworkitem) | **Post** /api/v2/workItems/{id}/restore | Restore test case, checklist or shared steps from archive
*WorkItemsApi* | [**UpdateWorkItem**](docs/WorkItemsApi.md#updateworkitem) | **Put** /api/v2/workItems | Update Test Case, Checklist or Shared Step
*WorkItemsCommentsApi* | [**ApiV2WorkItemsCommentsCommentIdDelete**](docs/WorkItemsCommentsApi.md#apiv2workitemscommentscommentiddelete) | **Delete** /api/v2/workItems/comments/{commentId} | Delete WorkItem comment
*WorkItemsCommentsApi* | [**ApiV2WorkItemsCommentsPost**](docs/WorkItemsCommentsApi.md#apiv2workitemscommentspost) | **Post** /api/v2/workItems/comments | Create WorkItem comment
*WorkItemsCommentsApi* | [**ApiV2WorkItemsCommentsPut**](docs/WorkItemsCommentsApi.md#apiv2workitemscommentsput) | **Put** /api/v2/workItems/comments | Update work item comment
*WorkItemsCommentsApi* | [**ApiV2WorkItemsIdCommentsGet**](docs/WorkItemsCommentsApi.md#apiv2workitemsidcommentsget) | **Get** /api/v2/workItems/{id}/comments | Get work item comments

## Documentation for Models

You can see the documentation [here](https://github.com/testgear-tms/api-client-golang/blob/main/docs/Readme.md)

# Contributing

You can help to develop the project. Any contributions are **greatly appreciated**.

* If you have suggestions for adding or removing projects, feel free to [open an issue](https://github.com/testgear-tms/api-client-golang/issues/new) to discuss it, or directly create a pull request after you edit the *README.md* file with necessary changes.
* Please make sure you check your spelling and grammar.
* Create individual PR for each suggestion.
* Please also read through the [Code Of Conduct](https://github.com/testgear-tms/api-client-golang/blob/main/CODE_OF_CONDUCT.md) before posting your first idea as well.

# License

Distributed under the Apache-2.0 License. See [LICENSE](https://github.com/testgear-tms/api-client-golang/blob/main/LICENSE.md) for more information.
