// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type DbtScheduleRunTypesType string

const (
	DbtScheduleRunTypesTypeInProgress             DbtScheduleRunTypesType = "IN_PROGRESS"
	DbtScheduleRunTypesTypeIngestCouldNotComplete DbtScheduleRunTypesType = "INGEST_COULD_NOT_COMPLETE"
	DbtScheduleRunTypesTypeDbtError               DbtScheduleRunTypesType = "DBT_ERROR"
	DbtScheduleRunTypesTypeSuccessWithDbtWarnings DbtScheduleRunTypesType = "SUCCESS_WITH_DBT_WARNINGS"
	DbtScheduleRunTypesTypeSuccess                DbtScheduleRunTypesType = "SUCCESS"
)

type DbtScheduleRunTypes struct {
	DbtScheduleRunSuccess    *DbtScheduleRunSuccess
	DbtScheduleRunFailure    *DbtScheduleRunFailure
	DbtScheduleRunInProgress *DbtScheduleRunInProgress

	Type DbtScheduleRunTypesType
}

func CreateDbtScheduleRunTypesInProgress(inProgress DbtScheduleRunInProgress) DbtScheduleRunTypes {
	typ := DbtScheduleRunTypesTypeInProgress

	typStr := DbtScheduleRunInProgressStatus(typ)
	inProgress.Status = typStr

	return DbtScheduleRunTypes{
		DbtScheduleRunInProgress: &inProgress,
		Type:                     typ,
	}
}

func CreateDbtScheduleRunTypesIngestCouldNotComplete(ingestCouldNotComplete DbtScheduleRunFailure) DbtScheduleRunTypes {
	typ := DbtScheduleRunTypesTypeIngestCouldNotComplete

	typStr := DbtScheduleRunFailureStatus(typ)
	ingestCouldNotComplete.Status = typStr

	return DbtScheduleRunTypes{
		DbtScheduleRunFailure: &ingestCouldNotComplete,
		Type:                  typ,
	}
}

func CreateDbtScheduleRunTypesDbtError(dbtError DbtScheduleRunFailure) DbtScheduleRunTypes {
	typ := DbtScheduleRunTypesTypeDbtError

	typStr := DbtScheduleRunFailureStatus(typ)
	dbtError.Status = typStr

	return DbtScheduleRunTypes{
		DbtScheduleRunFailure: &dbtError,
		Type:                  typ,
	}
}

func CreateDbtScheduleRunTypesSuccessWithDbtWarnings(successWithDbtWarnings DbtScheduleRunSuccess) DbtScheduleRunTypes {
	typ := DbtScheduleRunTypesTypeSuccessWithDbtWarnings

	typStr := Status(typ)
	successWithDbtWarnings.Status = typStr

	return DbtScheduleRunTypes{
		DbtScheduleRunSuccess: &successWithDbtWarnings,
		Type:                  typ,
	}
}

func CreateDbtScheduleRunTypesSuccess(success DbtScheduleRunSuccess) DbtScheduleRunTypes {
	typ := DbtScheduleRunTypesTypeSuccess

	typStr := Status(typ)
	success.Status = typStr

	return DbtScheduleRunTypes{
		DbtScheduleRunSuccess: &success,
		Type:                  typ,
	}
}

func (u *DbtScheduleRunTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Status string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Status {
	case "IN_PROGRESS":
		dbtScheduleRunInProgress := new(DbtScheduleRunInProgress)
		if err := utils.UnmarshalJSON(data, &dbtScheduleRunInProgress, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DbtScheduleRunInProgress = dbtScheduleRunInProgress
		u.Type = DbtScheduleRunTypesTypeInProgress
		return nil
	case "INGEST_COULD_NOT_COMPLETE":
		dbtScheduleRunFailure := new(DbtScheduleRunFailure)
		if err := utils.UnmarshalJSON(data, &dbtScheduleRunFailure, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DbtScheduleRunFailure = dbtScheduleRunFailure
		u.Type = DbtScheduleRunTypesTypeIngestCouldNotComplete
		return nil
	case "DBT_ERROR":
		dbtScheduleRunFailure := new(DbtScheduleRunFailure)
		if err := utils.UnmarshalJSON(data, &dbtScheduleRunFailure, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DbtScheduleRunFailure = dbtScheduleRunFailure
		u.Type = DbtScheduleRunTypesTypeDbtError
		return nil
	case "SUCCESS_WITH_DBT_WARNINGS":
		dbtScheduleRunSuccess := new(DbtScheduleRunSuccess)
		if err := utils.UnmarshalJSON(data, &dbtScheduleRunSuccess, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DbtScheduleRunSuccess = dbtScheduleRunSuccess
		u.Type = DbtScheduleRunTypesTypeSuccessWithDbtWarnings
		return nil
	case "SUCCESS":
		dbtScheduleRunSuccess := new(DbtScheduleRunSuccess)
		if err := utils.UnmarshalJSON(data, &dbtScheduleRunSuccess, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DbtScheduleRunSuccess = dbtScheduleRunSuccess
		u.Type = DbtScheduleRunTypesTypeSuccess
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DbtScheduleRunTypes) MarshalJSON() ([]byte, error) {
	if u.DbtScheduleRunSuccess != nil {
		return utils.MarshalJSON(u.DbtScheduleRunSuccess, "", true)
	}

	if u.DbtScheduleRunFailure != nil {
		return utils.MarshalJSON(u.DbtScheduleRunFailure, "", true)
	}

	if u.DbtScheduleRunInProgress != nil {
		return utils.MarshalJSON(u.DbtScheduleRunInProgress, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
