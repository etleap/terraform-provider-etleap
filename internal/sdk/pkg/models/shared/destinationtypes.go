// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type DestinationTypesType string

const (
	DestinationTypesTypeS3DataLake DestinationTypesType = "S3_DATA_LAKE"
	DestinationTypesTypeDeltaLake  DestinationTypesType = "DELTA_LAKE"
	DestinationTypesTypeIceberg    DestinationTypesType = "ICEBERG"
	DestinationTypesTypeSnowflake  DestinationTypesType = "SNOWFLAKE"
	DestinationTypesTypeRedshift   DestinationTypesType = "REDSHIFT"
)

type DestinationTypes struct {
	DestinationRedshift   *DestinationRedshift
	DestinationSnowflake  *DestinationSnowflake
	DestinationDeltaLake  *DestinationDeltaLake
	DestinationS3DataLake *DestinationS3DataLake
	DestinationIceberg    *DestinationIceberg

	Type DestinationTypesType
}

func CreateDestinationTypesS3DataLake(s3DataLake DestinationS3DataLake) DestinationTypes {
	typ := DestinationTypesTypeS3DataLake

	typStr := DestinationS3DataLakeType(typ)
	s3DataLake.Type = typStr

	return DestinationTypes{
		DestinationS3DataLake: &s3DataLake,
		Type:                  typ,
	}
}

func CreateDestinationTypesDeltaLake(deltaLake DestinationDeltaLake) DestinationTypes {
	typ := DestinationTypesTypeDeltaLake

	typStr := DestinationDeltaLakeType(typ)
	deltaLake.Type = typStr

	return DestinationTypes{
		DestinationDeltaLake: &deltaLake,
		Type:                 typ,
	}
}

func CreateDestinationTypesIceberg(iceberg DestinationIceberg) DestinationTypes {
	typ := DestinationTypesTypeIceberg

	typStr := DestinationIcebergType(typ)
	iceberg.Type = typStr

	return DestinationTypes{
		DestinationIceberg: &iceberg,
		Type:               typ,
	}
}

func CreateDestinationTypesSnowflake(snowflake DestinationSnowflake) DestinationTypes {
	typ := DestinationTypesTypeSnowflake

	typStr := DestinationSnowflakeType(typ)
	snowflake.Type = typStr

	return DestinationTypes{
		DestinationSnowflake: &snowflake,
		Type:                 typ,
	}
}

func CreateDestinationTypesRedshift(redshift DestinationRedshift) DestinationTypes {
	typ := DestinationTypesTypeRedshift

	typStr := DestinationRedshiftType(typ)
	redshift.Type = typStr

	return DestinationTypes{
		DestinationRedshift: &redshift,
		Type:                typ,
	}
}

func (u *DestinationTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Type string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "S3_DATA_LAKE":
		destinationS3DataLake := new(DestinationS3DataLake)
		if err := utils.UnmarshalJSON(data, &destinationS3DataLake, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DestinationS3DataLake = destinationS3DataLake
		u.Type = DestinationTypesTypeS3DataLake
		return nil
	case "DELTA_LAKE":
		destinationDeltaLake := new(DestinationDeltaLake)
		// [VIK-4496] Etleap monkey-patch: the unmarshal function allows unknown fields for ignoring schemaChangingTo and tableChangingTo
		if err := utils.UnmarshalJSON(data, &destinationDeltaLake, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DestinationDeltaLake = destinationDeltaLake
		u.Type = DestinationTypesTypeDeltaLake
		return nil
	case "ICEBERG":
		destinationIceberg := new(DestinationIceberg)
		if err := utils.UnmarshalJSON(data, &destinationIceberg, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DestinationIceberg = destinationIceberg
		u.Type = DestinationTypesTypeIceberg
		return nil
	case "SNOWFLAKE":
		destinationSnowflake := new(DestinationSnowflake)
		if err := utils.UnmarshalJSON(data, &destinationSnowflake, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DestinationSnowflake = destinationSnowflake
		u.Type = DestinationTypesTypeSnowflake
		return nil
	case "REDSHIFT":
		destinationRedshift := new(DestinationRedshift)
		if err := utils.UnmarshalJSON(data, &destinationRedshift, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DestinationRedshift = destinationRedshift
		u.Type = DestinationTypesTypeRedshift
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationTypes) MarshalJSON() ([]byte, error) {
	if u.DestinationRedshift != nil {
		return utils.MarshalJSON(u.DestinationRedshift, "", true)
	}

	if u.DestinationSnowflake != nil {
		return utils.MarshalJSON(u.DestinationSnowflake, "", true)
	}

	if u.DestinationDeltaLake != nil {
		return utils.MarshalJSON(u.DestinationDeltaLake, "", true)
	}

	if u.DestinationS3DataLake != nil {
		return utils.MarshalJSON(u.DestinationS3DataLake, "", true)
	}

	if u.DestinationIceberg != nil {
		return utils.MarshalJSON(u.DestinationIceberg, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
