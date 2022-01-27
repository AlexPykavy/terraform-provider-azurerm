package profiles

import "strings"

type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "application"
	IdentityTypeKey             IdentityType = "key"
	IdentityTypeManagedIdentity IdentityType = "managedIdentity"
	IdentityTypeUser            IdentityType = "user"
)

func PossibleValuesForIdentityType() []string {
	return []string{
		string(IdentityTypeApplication),
		string(IdentityTypeKey),
		string(IdentityTypeManagedIdentity),
		string(IdentityTypeUser),
	}
}

func parseIdentityType(input string) (*IdentityType, error) {
	vals := map[string]IdentityType{
		"application":     IdentityTypeApplication,
		"key":             IdentityTypeKey,
		"managedidentity": IdentityTypeManagedIdentity,
		"user":            IdentityTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityType(input)
	return &out, nil
}

type OptimizationType string

const (
	OptimizationTypeDynamicSiteAcceleration     OptimizationType = "DynamicSiteAcceleration"
	OptimizationTypeGeneralMediaStreaming       OptimizationType = "GeneralMediaStreaming"
	OptimizationTypeGeneralWebDelivery          OptimizationType = "GeneralWebDelivery"
	OptimizationTypeLargeFileDownload           OptimizationType = "LargeFileDownload"
	OptimizationTypeVideoOnDemandMediaStreaming OptimizationType = "VideoOnDemandMediaStreaming"
)

func PossibleValuesForOptimizationType() []string {
	return []string{
		string(OptimizationTypeDynamicSiteAcceleration),
		string(OptimizationTypeGeneralMediaStreaming),
		string(OptimizationTypeGeneralWebDelivery),
		string(OptimizationTypeLargeFileDownload),
		string(OptimizationTypeVideoOnDemandMediaStreaming),
	}
}

func parseOptimizationType(input string) (*OptimizationType, error) {
	vals := map[string]OptimizationType{
		"dynamicsiteacceleration":     OptimizationTypeDynamicSiteAcceleration,
		"generalmediastreaming":       OptimizationTypeGeneralMediaStreaming,
		"generalwebdelivery":          OptimizationTypeGeneralWebDelivery,
		"largefiledownload":           OptimizationTypeLargeFileDownload,
		"videoondemandmediastreaming": OptimizationTypeVideoOnDemandMediaStreaming,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OptimizationType(input)
	return &out, nil
}

type ProfileResourceState string

const (
	ProfileResourceStateActive   ProfileResourceState = "Active"
	ProfileResourceStateCreating ProfileResourceState = "Creating"
	ProfileResourceStateDeleting ProfileResourceState = "Deleting"
	ProfileResourceStateDisabled ProfileResourceState = "Disabled"
)

func PossibleValuesForProfileResourceState() []string {
	return []string{
		string(ProfileResourceStateActive),
		string(ProfileResourceStateCreating),
		string(ProfileResourceStateDeleting),
		string(ProfileResourceStateDisabled),
	}
}

func parseProfileResourceState(input string) (*ProfileResourceState, error) {
	vals := map[string]ProfileResourceState{
		"active":   ProfileResourceStateActive,
		"creating": ProfileResourceStateCreating,
		"deleting": ProfileResourceStateDeleting,
		"disabled": ProfileResourceStateDisabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProfileResourceState(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameCustomVerizon                             SkuName = "Custom_Verizon"
	SkuNamePremiumAzureFrontDoor                     SkuName = "Premium_AzureFrontDoor"
	SkuNamePremiumVerizon                            SkuName = "Premium_Verizon"
	SkuNameStandardAkamai                            SkuName = "Standard_Akamai"
	SkuNameStandardAvgBandWidthChinaCdn              SkuName = "Standard_AvgBandWidth_ChinaCdn"
	SkuNameStandardAzureFrontDoor                    SkuName = "Standard_AzureFrontDoor"
	SkuNameStandardChinaCdn                          SkuName = "Standard_ChinaCdn"
	SkuNameStandardMicrosoft                         SkuName = "Standard_Microsoft"
	SkuNameStandardNineFiveFiveBandWidthChinaCdn     SkuName = "Standard_955BandWidth_ChinaCdn"
	SkuNameStandardPlusAvgBandWidthChinaCdn          SkuName = "StandardPlus_AvgBandWidth_ChinaCdn"
	SkuNameStandardPlusChinaCdn                      SkuName = "StandardPlus_ChinaCdn"
	SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn SkuName = "StandardPlus_955BandWidth_ChinaCdn"
	SkuNameStandardVerizon                           SkuName = "Standard_Verizon"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameCustomVerizon),
		string(SkuNamePremiumAzureFrontDoor),
		string(SkuNamePremiumVerizon),
		string(SkuNameStandardAkamai),
		string(SkuNameStandardAvgBandWidthChinaCdn),
		string(SkuNameStandardAzureFrontDoor),
		string(SkuNameStandardChinaCdn),
		string(SkuNameStandardMicrosoft),
		string(SkuNameStandardNineFiveFiveBandWidthChinaCdn),
		string(SkuNameStandardPlusAvgBandWidthChinaCdn),
		string(SkuNameStandardPlusChinaCdn),
		string(SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn),
		string(SkuNameStandardVerizon),
	}
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"custom_verizon":                     SkuNameCustomVerizon,
		"premium_azurefrontdoor":             SkuNamePremiumAzureFrontDoor,
		"premium_verizon":                    SkuNamePremiumVerizon,
		"standard_akamai":                    SkuNameStandardAkamai,
		"standard_avgbandwidth_chinacdn":     SkuNameStandardAvgBandWidthChinaCdn,
		"standard_azurefrontdoor":            SkuNameStandardAzureFrontDoor,
		"standard_chinacdn":                  SkuNameStandardChinaCdn,
		"standard_microsoft":                 SkuNameStandardMicrosoft,
		"standard_955bandwidth_chinacdn":     SkuNameStandardNineFiveFiveBandWidthChinaCdn,
		"standardplus_avgbandwidth_chinacdn": SkuNameStandardPlusAvgBandWidthChinaCdn,
		"standardplus_chinacdn":              SkuNameStandardPlusChinaCdn,
		"standardplus_955bandwidth_chinacdn": SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn,
		"standard_verizon":                   SkuNameStandardVerizon,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}
