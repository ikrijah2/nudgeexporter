package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// BillingFrequency enumerates the values for billing frequency.
type BillingFrequency string

const (
	// Month ...
	Month BillingFrequency = "Month"
	// Quarter ...
	Quarter BillingFrequency = "Quarter"
	// Year ...
	Year BillingFrequency = "Year"
)

// PossibleBillingFrequencyValues returns an array of possible values for the BillingFrequency const type.
func PossibleBillingFrequencyValues() []BillingFrequency {
	return []BillingFrequency{Month, Quarter, Year}
}

// Bound enumerates the values for bound.
type Bound string

const (
	// Lower ...
	Lower Bound = "Lower"
	// Upper ...
	Upper Bound = "Upper"
)

// PossibleBoundValues returns an array of possible values for the Bound const type.
func PossibleBoundValues() []Bound {
	return []Bound{Lower, Upper}
}

// CategoryType enumerates the values for category type.
type CategoryType string

const (
	// Cost ...
	Cost CategoryType = "Cost"
	// Usage ...
	Usage CategoryType = "Usage"
)

// PossibleCategoryTypeValues returns an array of possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{Cost, Usage}
}

// ChargeType enumerates the values for charge type.
type ChargeType string

const (
	// ChargeTypeActual ...
	ChargeTypeActual ChargeType = "Actual"
	// ChargeTypeForecast ...
	ChargeTypeForecast ChargeType = "Forecast"
)

// PossibleChargeTypeValues returns an array of possible values for the ChargeType const type.
func PossibleChargeTypeValues() []ChargeType {
	return []ChargeType{ChargeTypeActual, ChargeTypeForecast}
}

// Datagrain enumerates the values for datagrain.
type Datagrain string

const (
	// DailyGrain Daily grain of data
	DailyGrain Datagrain = "daily"
	// MonthlyGrain Monthly grain of data
	MonthlyGrain Datagrain = "monthly"
)

// PossibleDatagrainValues returns an array of possible values for the Datagrain const type.
func PossibleDatagrainValues() []Datagrain {
	return []Datagrain{DailyGrain, MonthlyGrain}
}

// Grain enumerates the values for grain.
type Grain string

const (
	// Daily ...
	Daily Grain = "Daily"
	// Monthly ...
	Monthly Grain = "Monthly"
	// Yearly ...
	Yearly Grain = "Yearly"
)

// PossibleGrainValues returns an array of possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{Daily, Monthly, Yearly}
}

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// EqualTo ...
	EqualTo OperatorType = "EqualTo"
	// GreaterThan ...
	GreaterThan OperatorType = "GreaterThan"
	// GreaterThanOrEqualTo ...
	GreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{EqualTo, GreaterThan, GreaterThanOrEqualTo}
}

// TimeGrainType enumerates the values for time grain type.
type TimeGrainType string

const (
	// TimeGrainTypeAnnually ...
	TimeGrainTypeAnnually TimeGrainType = "Annually"
	// TimeGrainTypeMonthly ...
	TimeGrainTypeMonthly TimeGrainType = "Monthly"
	// TimeGrainTypeQuarterly ...
	TimeGrainTypeQuarterly TimeGrainType = "Quarterly"
)

// PossibleTimeGrainTypeValues returns an array of possible values for the TimeGrainType const type.
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return []TimeGrainType{TimeGrainTypeAnnually, TimeGrainTypeMonthly, TimeGrainTypeQuarterly}
}
