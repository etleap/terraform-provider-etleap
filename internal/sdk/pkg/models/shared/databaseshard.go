// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DatabaseShard - Use shards when the database is split amongst several physical machines, but should be treated as a single database. For performance reasons, only the first shard is validated.
type DatabaseShard struct {
	ShardID   string     `json:"shardId"`
	Address   string     `json:"address"`
	Port      int64      `json:"port"`
	Database  string     `json:"database"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	SSHConfig *SSHConfig `json:"sshConfig,omitempty"`
}

func (o *DatabaseShard) GetShardID() string {
	if o == nil {
		return ""
	}
	return o.ShardID
}

func (o *DatabaseShard) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *DatabaseShard) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *DatabaseShard) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DatabaseShard) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *DatabaseShard) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *DatabaseShard) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
