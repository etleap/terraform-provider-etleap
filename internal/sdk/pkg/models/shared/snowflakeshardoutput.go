// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SnowflakeShardOutput - Use shards when the database is split amongst several physical machines, but should be treated as a single database.
type SnowflakeShardOutput struct {
	ShardID  string `json:"shardId"`
	Address  string `json:"address"`
	Database string `json:"database"`
	// The virtual warehouse to use once connected.
	Warehouse string `json:"warehouse"`
	Username  string `json:"username"`
	// The role the user will use to connect
	Role *string `json:"role,omitempty"`
	// Snowflake Authentication Types
	Authentication *SnowflakeAuthenticationTypesOutput `json:"authentication,omitempty"`
}

func (o *SnowflakeShardOutput) GetShardID() string {
	if o == nil {
		return ""
	}
	return o.ShardID
}

func (o *SnowflakeShardOutput) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *SnowflakeShardOutput) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SnowflakeShardOutput) GetWarehouse() string {
	if o == nil {
		return ""
	}
	return o.Warehouse
}

func (o *SnowflakeShardOutput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *SnowflakeShardOutput) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *SnowflakeShardOutput) GetAuthentication() *SnowflakeAuthenticationTypesOutput {
	if o == nil {
		return nil
	}
	return o.Authentication
}

func (o *SnowflakeShardOutput) GetAuthenticationPassword() *SnowflakeAuthenticationPasswordOutput {
	if v := o.GetAuthentication(); v != nil {
		return v.SnowflakeAuthenticationPasswordOutput
	}
	return nil
}

func (o *SnowflakeShardOutput) GetAuthenticationKeyPair() *SnowflakeAuthenticationKeyPairOutput {
	if v := o.GetAuthentication(); v != nil {
		return v.SnowflakeAuthenticationKeyPairOutput
	}
	return nil
}
