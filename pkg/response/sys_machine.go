package response

// 机器信息响应, 字段含义见models.SysMachine
type MachineListResponseStruct struct {
	BaseData
	Host      string `json:"host"`
	SshPort   int    `json:"sshPort"`
	Version   string `json:"version"`
	Name      string `json:"name"`
	Arch      string `json:"arch"`
	Cpu       string `json:"cpu"`
	Memory    string `json:"memory"`
	Disk      string `json:"disk"`
	LoginName string `json:"loginName"`
	LoginPwd  string `json:"loginPwd"`
	Status    *uint  `json:"status"`
	Remark    string `json:"remark"`
	Creator   string `json:"creator"`
}
