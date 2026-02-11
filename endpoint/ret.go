package endpoint

import "fmt"

// MutateAccountBudgetProposal https://googleads.googleapis.com/v23/customers/{customer_id=*}/accountBudgetProposals:mutate
func MutateAccountBudgetProposal(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/accountBudgetProposals:mutate", customer_id)
}

// CreateAccountLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/accountLinks:create
func CreateAccountLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/accountLinks:create", customer_id)
}

// MutateAccountLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/accountLinks:mutate
func MutateAccountLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/accountLinks:mutate", customer_id)
}

// MutateAdGroupAdLabels https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupAdLabels:mutate
func MutateAdGroupAdLabels(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupAdLabels:mutate", customer_id)
}

// MutateAdGroupAds https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupAds:mutate
func MutateAdGroupAds(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupAds:mutate", customer_id)
}

// RemoveAutomaticallyCreatedAssets https://googleads.googleapis.com/v23/{ad_group_ad=customers/*/adGroupAds/*}:removeAutomaticallyCreatedAssets
func RemoveAutomaticallyCreatedAssets(ad_group_ad string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:removeAutomaticallyCreatedAssets", ad_group_ad)
}

// MutateAdGroupAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupAssets:mutate
func MutateAdGroupAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupAssets:mutate", customer_id)
}

// MutateAdGroupAssetSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupAssetSets:mutate
func MutateAdGroupAssetSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupAssetSets:mutate", customer_id)
}

// MutateAdGroupBidModifiers https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupBidModifiers:mutate
func MutateAdGroupBidModifiers(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupBidModifiers:mutate", customer_id)
}

// MutateAdGroupCriterionCustomizers https://googleads.googleapis.com/v23/customers/{customer_id=*}/AdGroupCriterionCustomizers:mutate
func MutateAdGroupCriterionCustomizers(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/AdGroupCriterionCustomizers:mutate", customer_id)
}

// MutateAdGroupCriterionLabels https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupCriterionLabels:mutate
func MutateAdGroupCriterionLabels(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupCriterionLabels:mutate", customer_id)
}

// MutateAdGroupCriteria https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupCriteria:mutate
func MutateAdGroupCriteria(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupCriteria:mutate", customer_id)
}

// MutateAdGroupCustomizers https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupCustomizers:mutate
func MutateAdGroupCustomizers(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupCustomizers:mutate", customer_id)
}

// MutateAdGroupLabels https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroupLabels:mutate
func MutateAdGroupLabels(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroupLabels:mutate", customer_id)
}

// MutateAdGroups https://googleads.googleapis.com/v23/customers/{customer_id=*}/adGroups:mutate
func MutateAdGroups(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adGroups:mutate", customer_id)
}

// MutateAdParameters https://googleads.googleapis.com/v23/customers/{customer_id=*}/adParameters:mutate
func MutateAdParameters(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/adParameters:mutate", customer_id)
}

// MutateAds https://googleads.googleapis.com/v23/customers/{customer_id=*}/ads:mutate
func MutateAds(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/ads:mutate", customer_id)
}

// GenerateText https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetGenerations:generateText
func GenerateText(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetGenerations:generateText", customer_id)
}

// GenerateImages https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetGenerations:generateImages
func GenerateImages(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetGenerations:generateImages", customer_id)
}

// MutateAssetGroupAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetGroupAssets:mutate
func MutateAssetGroupAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetGroupAssets:mutate", customer_id)
}

// MutateAssetGroupListingGroupFilters https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetGroupListingGroupFilters:mutate
func MutateAssetGroupListingGroupFilters(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetGroupListingGroupFilters:mutate", customer_id)
}

// MutateAssetGroups https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetGroups:mutate
func MutateAssetGroups(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetGroups:mutate", customer_id)
}

// MutateAssetGroupSignals https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetGroupSignals:mutate
func MutateAssetGroupSignals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetGroupSignals:mutate", customer_id)
}

// MutateAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}/assets:mutate
func MutateAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assets:mutate", customer_id)
}

// MutateAssetSetAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetSetAssets:mutate
func MutateAssetSetAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetSetAssets:mutate", customer_id)
}

// MutateAssetSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/assetSets:mutate
func MutateAssetSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/assetSets:mutate", customer_id)
}

// GenerateInsightsFinderReport https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateInsightsFinderReport
func GenerateInsightsFinderReport(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateInsightsFinderReport", customer_id)
}

// ListAudienceInsightsAttributes https://googleads.googleapis.com/v23/customers/{customer_id=*}:searchAudienceInsightsAttributes
func ListAudienceInsightsAttributes(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:searchAudienceInsightsAttributes", customer_id)
}

// ListInsightsEligibleDates https://googleads.googleapis.com/v23/audienceInsights:listInsightsEligibleDates
func ListInsightsEligibleDates() string {
	return "https://googleads.googleapis.com/v23/audienceInsights:listInsightsEligibleDates"
}

// GenerateAudienceCompositionInsights https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateAudienceCompositionInsights
func GenerateAudienceCompositionInsights(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateAudienceCompositionInsights", customer_id)
}

// GenerateAudienceDefinition https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateAudienceDefinition
func GenerateAudienceDefinition(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateAudienceDefinition", customer_id)
}

// GenerateSuggestedTargetingInsights https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateSuggestedTargetingInsights
func GenerateSuggestedTargetingInsights(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateSuggestedTargetingInsights", customer_id)
}

// GenerateAudienceOverlapInsights https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateAudienceOverlapInsights
func GenerateAudienceOverlapInsights(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateAudienceOverlapInsights", customer_id)
}

// GenerateTargetingSuggestionMetrics https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateTargetingSuggestionMetrics
func GenerateTargetingSuggestionMetrics(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateTargetingSuggestionMetrics", customer_id)
}

// MutateAudiences https://googleads.googleapis.com/v23/customers/{customer_id=*}/audiences:mutate
func MutateAudiences(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/audiences:mutate", customer_id)
}

// RemoveCampaignAutomaticallyCreatedAsset https://googleads.googleapis.com/v23/customers/{customer_id=*}:removeCampaignAutomaticallyCreatedAsset
func RemoveCampaignAutomaticallyCreatedAsset(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:removeCampaignAutomaticallyCreatedAsset", customer_id)
}

// MutateBatchJob https://googleads.googleapis.com/v23/customers/{customer_id=*}/batchJobs:mutate
func MutateBatchJob(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/batchJobs:mutate", customer_id)
}

// ListBatchJobResults https://googleads.googleapis.com/v23/{resource_name=customers/*/batchJobs/*}:listResults
func ListBatchJobResults(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:listResults", resource_name)
}

// RunBatchJob https://googleads.googleapis.com/v23/{resource_name=customers/*/batchJobs/*}:run
func RunBatchJob(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:run", resource_name)
}

// AddBatchJobOperations https://googleads.googleapis.com/v23/{resource_name=customers/*/batchJobs/*}:addOperations
func AddBatchJobOperations(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:addOperations", resource_name)
}

// ListBenchmarksAvailableDates https://googleads.googleapis.com/v23:listBenchmarksAvailableDates
func ListBenchmarksAvailableDates() string {
	return "https://googleads.googleapis.com/v23:listBenchmarksAvailableDates"
}

// ListBenchmarksLocations https://googleads.googleapis.com/v23:listBenchmarksLocations
func ListBenchmarksLocations() string {
	return "https://googleads.googleapis.com/v23:listBenchmarksLocations"
}

// ListBenchmarksProducts https://googleads.googleapis.com/v23:listBenchmarksProducts
func ListBenchmarksProducts() string {
	return "https://googleads.googleapis.com/v23:listBenchmarksProducts"
}

// ListBenchmarksSources https://googleads.googleapis.com/v23:listBenchmarksSources
func ListBenchmarksSources() string {
	return "https://googleads.googleapis.com/v23:listBenchmarksSources"
}

// GenerateBenchmarksMetrics https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateBenchmarksMetrics
func GenerateBenchmarksMetrics(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateBenchmarksMetrics", customer_id)
}

// MutateBiddingDataExclusions https://googleads.googleapis.com/v23/customers/{customer_id=*}/biddingDataExclusions:mutate
func MutateBiddingDataExclusions(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/biddingDataExclusions:mutate", customer_id)
}

// MutateBiddingSeasonalityAdjustments https://googleads.googleapis.com/v23/customers/{customer_id=*}/biddingSeasonalityAdjustments:mutate
func MutateBiddingSeasonalityAdjustments(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/biddingSeasonalityAdjustments:mutate", customer_id)
}

// MutateBiddingStrategies https://googleads.googleapis.com/v23/customers/{customer_id=*}/biddingStrategies:mutate
func MutateBiddingStrategies(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/biddingStrategies:mutate", customer_id)
}

// MutateBillingSetup https://googleads.googleapis.com/v23/customers/{customer_id=*}/billingSetups:mutate
func MutateBillingSetup(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/billingSetups:mutate", customer_id)
}

// SuggestBrands https://googleads.googleapis.com/v23/customers/{customer_id=*}:suggestBrands
func SuggestBrands(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:suggestBrands", customer_id)
}

// MutateCampaignAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignAssets:mutate
func MutateCampaignAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignAssets:mutate", customer_id)
}

// MutateCampaignAssetSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignAssetSets:mutate
func MutateCampaignAssetSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignAssetSets:mutate", customer_id)
}

// MutateCampaignBidModifiers https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignBidModifiers:mutate
func MutateCampaignBidModifiers(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignBidModifiers:mutate", customer_id)
}

// MutateCampaignBudgets https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignBudgets:mutate
func MutateCampaignBudgets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignBudgets:mutate", customer_id)
}

// MutateCampaignConversionGoals https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignConversionGoals:mutate
func MutateCampaignConversionGoals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignConversionGoals:mutate", customer_id)
}

// MutateCampaignCriteria https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignCriteria:mutate
func MutateCampaignCriteria(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignCriteria:mutate", customer_id)
}

// MutateCampaignCustomizers https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignCustomizers:mutate
func MutateCampaignCustomizers(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignCustomizers:mutate", customer_id)
}

// MutateCampaignDrafts https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignDrafts:mutate
func MutateCampaignDrafts(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignDrafts:mutate", customer_id)
}

// PromoteCampaignDraft https://googleads.googleapis.com/v23/{campaign_draft=customers/*/campaignDrafts/*}:promote
func PromoteCampaignDraft(campaign_draft string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:promote", campaign_draft)
}

// ListCampaignDraftAsyncErrors https://googleads.googleapis.com/v23/{resource_name=customers/*/campaignDrafts/*}:listAsyncErrors
func ListCampaignDraftAsyncErrors(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:listAsyncErrors", resource_name)
}

// MutateCampaignGoalConfigs https://googleads.googleapis.com/v23/customers/{customer_id=*}/CampaignGoalConfigs:mutate
func MutateCampaignGoalConfigs(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/CampaignGoalConfigs:mutate", customer_id)
}

// MutateCampaignGroups https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignGroups:mutate
func MutateCampaignGroups(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignGroups:mutate", customer_id)
}

// MutateCampaignLabels https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignLabels:mutate
func MutateCampaignLabels(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignLabels:mutate", customer_id)
}

// ConfigureCampaignLifecycleGoals https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignLifecycleGoal:configureCampaignLifecycleGoals
func ConfigureCampaignLifecycleGoals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignLifecycleGoal:configureCampaignLifecycleGoals", customer_id)
}

// MutateCampaigns https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaigns:mutate
func MutateCampaigns(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaigns:mutate", customer_id)
}

// EnablePMaxBrandGuidelines https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaigns:enablePMaxBrandGuidelines
func EnablePMaxBrandGuidelines(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaigns:enablePMaxBrandGuidelines", customer_id)
}

// MutateCampaignSharedSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/campaignSharedSets:mutate
func MutateCampaignSharedSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/campaignSharedSets:mutate", customer_id)
}

// GenerateCreatorInsights https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateCreatorInsights
func GenerateCreatorInsights(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateCreatorInsights", customer_id)
}

// GenerateTrendingInsights https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateTrendingInsights
func GenerateTrendingInsights(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateTrendingInsights", customer_id)
}

// MutateConversionActions https://googleads.googleapis.com/v23/customers/{customer_id=*}/conversionActions:mutate
func MutateConversionActions(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/conversionActions:mutate", customer_id)
}

// UploadConversionAdjustments https://googleads.googleapis.com/v23/customers/{customer_id=*}:uploadConversionAdjustments
func UploadConversionAdjustments(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:uploadConversionAdjustments", customer_id)
}

// MutateConversionCustomVariables https://googleads.googleapis.com/v23/customers/{customer_id=*}/conversionCustomVariables:mutate
func MutateConversionCustomVariables(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/conversionCustomVariables:mutate", customer_id)
}

// MutateConversionGoalCampaignConfigs https://googleads.googleapis.com/v23/customers/{customer_id=*}/conversionGoalCampaignConfigs:mutate
func MutateConversionGoalCampaignConfigs(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/conversionGoalCampaignConfigs:mutate", customer_id)
}

// UploadClickConversions https://googleads.googleapis.com/v23/customers/{customer_id=*}:uploadClickConversions
func UploadClickConversions(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:uploadClickConversions", customer_id)
}

// UploadCallConversions https://googleads.googleapis.com/v23/customers/{customer_id=*}:uploadCallConversions
func UploadCallConversions(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:uploadCallConversions", customer_id)
}

// MutateConversionValueRules https://googleads.googleapis.com/v23/customers/{customer_id=*}/conversionValueRules:mutate
func MutateConversionValueRules(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/conversionValueRules:mutate", customer_id)
}

// MutateConversionValueRuleSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/conversionValueRuleSets:mutate
func MutateConversionValueRuleSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/conversionValueRuleSets:mutate", customer_id)
}

// MutateCustomAudiences https://googleads.googleapis.com/v23/customers/{customer_id=*}/customAudiences:mutate
func MutateCustomAudiences(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customAudiences:mutate", customer_id)
}

// MutateCustomConversionGoals https://googleads.googleapis.com/v23/customers/{customer_id=*}/customConversionGoals:mutate
func MutateCustomConversionGoals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customConversionGoals:mutate", customer_id)
}

// MutateCustomInterests https://googleads.googleapis.com/v23/customers/{customer_id=*}/customInterests:mutate
func MutateCustomInterests(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customInterests:mutate", customer_id)
}

// MutateCustomerAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerAssets:mutate
func MutateCustomerAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerAssets:mutate", customer_id)
}

// MutateCustomerAssetSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerAssetSets:mutate
func MutateCustomerAssetSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerAssetSets:mutate", customer_id)
}

// MutateCustomerClientLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerClientLinks:mutate
func MutateCustomerClientLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerClientLinks:mutate", customer_id)
}

// MutateCustomerConversionGoals https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerConversionGoals:mutate
func MutateCustomerConversionGoals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerConversionGoals:mutate", customer_id)
}

// MutateCustomerCustomizers https://googleads.googleapis.com/v23/customers/{customer_id=*}/CustomerCustomizers:mutate
func MutateCustomerCustomizers(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/CustomerCustomizers:mutate", customer_id)
}

// MutateCustomerLabels https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerLabels:mutate
func MutateCustomerLabels(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerLabels:mutate", customer_id)
}

// ConfigureCustomerLifecycleGoals https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerLifecycleGoal:configureCustomerLifecycleGoals
func ConfigureCustomerLifecycleGoals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerLifecycleGoal:configureCustomerLifecycleGoals", customer_id)
}

// MutateCustomerManagerLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerManagerLinks:mutate
func MutateCustomerManagerLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerManagerLinks:mutate", customer_id)
}

// MoveManagerLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerManagerLinks:moveManagerLink
func MoveManagerLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerManagerLinks:moveManagerLink", customer_id)
}

// MutateCustomerNegativeCriteria https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerNegativeCriteria:mutate
func MutateCustomerNegativeCriteria(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerNegativeCriteria:mutate", customer_id)
}

// MutateCustomer https://googleads.googleapis.com/v23/customers/{customer_id=*}:mutate
func MutateCustomer(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:mutate", customer_id)
}

// ListAccessibleCustomers https://googleads.googleapis.com/v23/customers:listAccessibleCustomers
func ListAccessibleCustomers() string {
	return "https://googleads.googleapis.com/v23/customers:listAccessibleCustomers"
}

// CreateCustomerClient https://googleads.googleapis.com/v23/customers/{customer_id=*}:createCustomerClient
func CreateCustomerClient(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:createCustomerClient", customer_id)
}

// MutateCustomerSkAdNetworkConversionValueSchema https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerSkAdNetworkConversionValueSchemas:mutate
func MutateCustomerSkAdNetworkConversionValueSchema(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerSkAdNetworkConversionValueSchemas:mutate", customer_id)
}

// MutateCustomerUserAccessInvitation https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerUserAccessInvitations:mutate
func MutateCustomerUserAccessInvitation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerUserAccessInvitations:mutate", customer_id)
}

// MutateCustomerUserAccess https://googleads.googleapis.com/v23/customers/{customer_id=*}/customerUserAccesses:mutate
func MutateCustomerUserAccess(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customerUserAccesses:mutate", customer_id)
}

// MutateCustomizerAttributes https://googleads.googleapis.com/v23/customers/{customer_id=*}/customizerAttributes:mutate
func MutateCustomizerAttributes(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/customizerAttributes:mutate", customer_id)
}

// CreateDataLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/dataLinks:create
func CreateDataLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/dataLinks:create", customer_id)
}

// RemoveDataLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/dataLinks:remove
func RemoveDataLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/dataLinks:remove", customer_id)
}

// UpdateDataLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/dataLinks:update
func UpdateDataLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/dataLinks:update", customer_id)
}

// MutateExperimentArms https://googleads.googleapis.com/v23/customers/{customer_id=*}/experimentArms:mutate
func MutateExperimentArms(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/experimentArms:mutate", customer_id)
}

// MutateExperiments https://googleads.googleapis.com/v23/customers/{customer_id=*}/experiments:mutate
func MutateExperiments(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/experiments:mutate", customer_id)
}

// EndExperiment https://googleads.googleapis.com/v23/{experiment=customers/*/experiments/*}:endExperiment
func EndExperiment(experiment string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:endExperiment", experiment)
}

// ListExperimentAsyncErrors https://googleads.googleapis.com/v23/{resource_name=customers/*/experiments/*}:listExperimentAsyncErrors
func ListExperimentAsyncErrors(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:listExperimentAsyncErrors", resource_name)
}

// GraduateExperiment https://googleads.googleapis.com/v23/{experiment=customers/*/experiments/*}:graduateExperiment
func GraduateExperiment(experiment string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:graduateExperiment", experiment)
}

// ScheduleExperiment https://googleads.googleapis.com/v23/{resource_name=customers/*/experiments/*}:scheduleExperiment
func ScheduleExperiment(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:scheduleExperiment", resource_name)
}

// PromoteExperiment https://googleads.googleapis.com/v23/{resource_name=customers/*/experiments/*}:promoteExperiment
func PromoteExperiment(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:promoteExperiment", resource_name)
}

// SuggestGeoTargetConstants https://googleads.googleapis.com/v23/geoTargetConstants:suggest
func SuggestGeoTargetConstants() string {
	return "https://googleads.googleapis.com/v23/geoTargetConstants:suggest"
}

// MutateGoals https://googleads.googleapis.com/v23/customers/{customer_id=*}/Goals:mutate
func MutateGoals(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/Goals:mutate", customer_id)
}

// GetGoogleAdsField https://googleads.googleapis.com/v23/{resource_name=googleAdsFields/*}
func GetGoogleAdsField(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s", resource_name)
}

// SearchGoogleAdsFields https://googleads.googleapis.com/v23/googleAdsFields:search
func SearchGoogleAdsFields() string {
	return "https://googleads.googleapis.com/v23/googleAdsFields:search"
}

// Search https://googleads.googleapis.com/v23/customers/{customer_id=*}/googleAds:search
func Search(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/googleAds:search", customer_id)
}

// SearchStream https://googleads.googleapis.com/v23/customers/{customer_id=*}/googleAds:searchStream
func SearchStream(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/googleAds:searchStream", customer_id)
}

// Mutate https://googleads.googleapis.com/v23/customers/{customer_id=*}/googleAds:mutate
func Mutate(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/googleAds:mutate", customer_id)
}

// StartIdentityVerification https://googleads.googleapis.com/v23/customers/{customer_id=*}:startIdentityVerification
func StartIdentityVerification(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:startIdentityVerification", customer_id)
}

// GetIdentityVerification https://googleads.googleapis.com/v23/customers/{customer_id=*}/getIdentityVerification
func GetIdentityVerification(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/getIdentityVerification", customer_id)
}

// FetchIncentive https://googleads.googleapis.com/v23/incentives:fetchIncentive
func FetchIncentive() string {
	return "https://googleads.googleapis.com/v23/incentives:fetchIncentive"
}

// ApplyIncentive https://googleads.googleapis.com/v23/customers/{customer_id=*}/incentives/{selected_incentive_id=*}:applyIncentive
func ApplyIncentive(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/incentives/{selected_incentive_id=*}:applyIncentive", customer_id)
}

// ListInvoices https://googleads.googleapis.com/v23/customers/{customer_id=*}/invoices
func ListInvoices(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/invoices", customer_id)
}

// MutateKeywordPlanAdGroupKeywords https://googleads.googleapis.com/v23/customers/{customer_id=*}/keywordPlanAdGroupKeywords:mutate
func MutateKeywordPlanAdGroupKeywords(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/keywordPlanAdGroupKeywords:mutate", customer_id)
}

// MutateKeywordPlanAdGroups https://googleads.googleapis.com/v23/customers/{customer_id=*}/keywordPlanAdGroups:mutate
func MutateKeywordPlanAdGroups(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/keywordPlanAdGroups:mutate", customer_id)
}

// MutateKeywordPlanCampaignKeywords https://googleads.googleapis.com/v23/customers/{customer_id=*}/keywordPlanCampaignKeywords:mutate
func MutateKeywordPlanCampaignKeywords(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/keywordPlanCampaignKeywords:mutate", customer_id)
}

// MutateKeywordPlanCampaigns https://googleads.googleapis.com/v23/customers/{customer_id=*}/keywordPlanCampaigns:mutate
func MutateKeywordPlanCampaigns(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/keywordPlanCampaigns:mutate", customer_id)
}

// GenerateKeywordIdeas https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateKeywordIdeas
func GenerateKeywordIdeas(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateKeywordIdeas", customer_id)
}

// GenerateKeywordHistoricalMetrics https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateKeywordHistoricalMetrics
func GenerateKeywordHistoricalMetrics(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateKeywordHistoricalMetrics", customer_id)
}

// GenerateAdGroupThemes https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateAdGroupThemes
func GenerateAdGroupThemes(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateAdGroupThemes", customer_id)
}

// GenerateKeywordForecastMetrics https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateKeywordForecastMetrics
func GenerateKeywordForecastMetrics(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateKeywordForecastMetrics", customer_id)
}

// MutateKeywordPlans https://googleads.googleapis.com/v23/customers/{customer_id=*}/keywordPlans:mutate
func MutateKeywordPlans(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/keywordPlans:mutate", customer_id)
}

// SuggestKeywordThemeConstants https://googleads.googleapis.com/v23/keywordThemeConstants:suggest
func SuggestKeywordThemeConstants() string {
	return "https://googleads.googleapis.com/v23/keywordThemeConstants:suggest"
}

// MutateLabels https://googleads.googleapis.com/v23/customers/{customer_id=*}/labels:mutate
func MutateLabels(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/labels:mutate", customer_id)
}

// AppendLeadConversation https://googleads.googleapis.com/v23/customers/{customer_id=*}/localServices:appendLeadConversation
func AppendLeadConversation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/localServices:appendLeadConversation", customer_id)
}

// ProvideLeadFeedback https://googleads.googleapis.com/v23/{resource_name=customers/*/localServicesLeads/*}:provideLeadFeedback
func ProvideLeadFeedback(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:provideLeadFeedback", resource_name)
}

// CreateOfflineUserDataJob https://googleads.googleapis.com/v23/customers/{customer_id=*}/offlineUserDataJobs:create
func CreateOfflineUserDataJob(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/offlineUserDataJobs:create", customer_id)
}

// AddOfflineUserDataJobOperations https://googleads.googleapis.com/v23/{resource_name=customers/*/offlineUserDataJobs/*}:addOperations
func AddOfflineUserDataJobOperations(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:addOperations", resource_name)
}

// RunOfflineUserDataJob https://googleads.googleapis.com/v23/{resource_name=customers/*/offlineUserDataJobs/*}:run
func RunOfflineUserDataJob(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:run", resource_name)
}

// ListPaymentsAccounts https://googleads.googleapis.com/v23/customers/{customer_id=*}/paymentsAccounts
func ListPaymentsAccounts(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/paymentsAccounts", customer_id)
}

// CreateProductLinkInvitation https://googleads.googleapis.com/v23/customers/{customer_id=*}/productLinkInvitations:create
func CreateProductLinkInvitation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/productLinkInvitations:create", customer_id)
}

// UpdateProductLinkInvitation https://googleads.googleapis.com/v23/customers/{customer_id=*}/productLinkInvitations:update
func UpdateProductLinkInvitation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/productLinkInvitations:update", customer_id)
}

// RemoveProductLinkInvitation https://googleads.googleapis.com/v23/customers/{customer_id=*}/productLinkInvitations:remove
func RemoveProductLinkInvitation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/productLinkInvitations:remove", customer_id)
}

// CreateProductLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/productLinks:create
func CreateProductLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/productLinks:create", customer_id)
}

// RemoveProductLink https://googleads.googleapis.com/v23/customers/{customer_id=*}/productLinks:remove
func RemoveProductLink(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/productLinks:remove", customer_id)
}

// GenerateConversionRates https://googleads.googleapis.com/v23:generateConversionRates
func GenerateConversionRates() string {
	return "https://googleads.googleapis.com/v23:generateConversionRates"
}

// ListPlannableLocations https://googleads.googleapis.com/v23:listPlannableLocations
func ListPlannableLocations() string {
	return "https://googleads.googleapis.com/v23:listPlannableLocations"
}

// ListPlannableProducts https://googleads.googleapis.com/v23:listPlannableProducts
func ListPlannableProducts() string {
	return "https://googleads.googleapis.com/v23:listPlannableProducts"
}

// GenerateReachForecast https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateReachForecast
func GenerateReachForecast(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateReachForecast", customer_id)
}

// ListPlannableUserLists https://googleads.googleapis.com/v23:listPlannableUserLists
func ListPlannableUserLists() string {
	return "https://googleads.googleapis.com/v23:listPlannableUserLists"
}

// ListPlannableUserInterests https://googleads.googleapis.com/v23:listPlannableUserInterests
func ListPlannableUserInterests() string {
	return "https://googleads.googleapis.com/v23:listPlannableUserInterests"
}

// ApplyRecommendation https://googleads.googleapis.com/v23/customers/{customer_id=*}/recommendations:apply
func ApplyRecommendation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/recommendations:apply", customer_id)
}

// DismissRecommendation https://googleads.googleapis.com/v23/customers/{customer_id=*}/recommendations:dismiss
func DismissRecommendation(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/recommendations:dismiss", customer_id)
}

// GenerateRecommendations https://googleads.googleapis.com/v23/customers/{customer_id=*}/recommendations:generate
func GenerateRecommendations(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/recommendations:generate", customer_id)
}

// MutateRecommendationSubscription https://googleads.googleapis.com/v23/customers/{customer_id=*}/recommendationSubscriptions:mutateRecommendationSubscription
func MutateRecommendationSubscription(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/recommendationSubscriptions:mutateRecommendationSubscription", customer_id)
}

// MutateRemarketingActions https://googleads.googleapis.com/v23/customers/{customer_id=*}/remarketingActions:mutate
func MutateRemarketingActions(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/remarketingActions:mutate", customer_id)
}

// GenerateShareablePreviews https://googleads.googleapis.com/v23/customers/{customer_id=*}:generateShareablePreviews
func GenerateShareablePreviews(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:generateShareablePreviews", customer_id)
}

// MutateSharedCriteria https://googleads.googleapis.com/v23/customers/{customer_id=*}/sharedCriteria:mutate
func MutateSharedCriteria(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/sharedCriteria:mutate", customer_id)
}

// MutateSharedSets https://googleads.googleapis.com/v23/customers/{customer_id=*}/sharedSets:mutate
func MutateSharedSets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/sharedSets:mutate", customer_id)
}

// GetSmartCampaignStatus https://googleads.googleapis.com/v23/{resource_name=customers/*/smartCampaignSettings/*}:getSmartCampaignStatus
func GetSmartCampaignStatus(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:getSmartCampaignStatus", resource_name)
}

// MutateSmartCampaignSettings https://googleads.googleapis.com/v23/customers/{customer_id=*}/smartCampaignSettings:mutate
func MutateSmartCampaignSettings(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/smartCampaignSettings:mutate", customer_id)
}

// SuggestSmartCampaignBudgetOptions https://googleads.googleapis.com/v23/customers/{customer_id=*}:suggestSmartCampaignBudgetOptions
func SuggestSmartCampaignBudgetOptions(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:suggestSmartCampaignBudgetOptions", customer_id)
}

// SuggestSmartCampaignAd https://googleads.googleapis.com/v23/customers/{customer_id=*}:suggestSmartCampaignAd
func SuggestSmartCampaignAd(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:suggestSmartCampaignAd", customer_id)
}

// SuggestKeywordThemes https://googleads.googleapis.com/v23/customers/{customer_id=*}:suggestKeywordThemes
func SuggestKeywordThemes(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:suggestKeywordThemes", customer_id)
}

// RegenerateShareableLinkId https://googleads.googleapis.com/v23/{resource_name=customers/*/thirdPartyAppAnalyticsLinks/*}:regenerateShareableLinkId
func RegenerateShareableLinkId(resource_name string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/%s:regenerateShareableLinkId", resource_name)
}

// SuggestTravelAssets https://googleads.googleapis.com/v23/customers/{customer_id=*}:suggestTravelAssets
func SuggestTravelAssets(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:suggestTravelAssets", customer_id)
}

// UploadUserData https://googleads.googleapis.com/v23/customers/{customer_id=*}:uploadUserData
func UploadUserData(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s:uploadUserData", customer_id)
}

// MutateUserListCustomerTypes https://googleads.googleapis.com/v23/customers/{customer_id=*}/userListCustomerTypes:mutate
func MutateUserListCustomerTypes(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/userListCustomerTypes:mutate", customer_id)
}

// MutateUserLists https://googleads.googleapis.com/v23/customers/{customer_id=*}/userLists:mutate
func MutateUserLists(customer_id string) string {
	return fmt.Sprintf("https://googleads.googleapis.com/v23/customers/%s/userLists:mutate", customer_id)
}
