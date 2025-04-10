// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MysqlShard - Use shards when the database is split amongst several physical machines, but should be treated as a single database. For performance reasons, only the first shard is validated.
type MysqlShard struct {
	ShardID string `json:"shardId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Database  *string    `json:"database,omitempty"`
	Username  string     `json:"username"`
	SSHConfig *SSHConfig `json:"sshConfig,omitempty"`
	Password  string     `json:"password"`
	Port      int64      `json:"port"`
	Address   string     `json:"address"`
}

func (o *MysqlShard) GetShardID() string {
	if o == nil {
		return ""
	}
	return o.ShardID
}

func (o *MysqlShard) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *MysqlShard) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *MysqlShard) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

func (o *MysqlShard) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *MysqlShard) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *MysqlShard) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}
