package config

import (
	"github.com/DSiSc/craft/types"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	assert := assert.New(t)
	conf := New("config.json")
	assert.NotNil(conf)
	path := conf.filePath
	ok := strings.Contains(path, "/config/config.json")
	assert.True(ok)
}

func Test_NewNodeConfig(t *testing.T) {
	assert := assert.New(t)
	nodeConf := NewNodeConfig()
	assert.NotNil(nodeConf)
	assert.Equal(uint64(4096), nodeConf.TxPoolConf.GlobalSlots)
	assert.NotNil("solo", nodeConf.ParticipatesConf.PolicyName)
	assert.NotNil("solo_node", nodeConf.Account)
	assert.Equal("tcp://0.0.0.0:47768", nodeConf.ApiGatewayAddr)
	assert.Equal(uint8(2), nodeConf.BlockInterval)

	var address = types.Address{
		0x33, 0x3c, 0x33, 0x10, 0x82, 0x4b, 0x7c, 0x68, 0x51, 0x33,
		0xf2, 0xbe, 0xdb, 0x2c, 0xa4, 0xb8, 0xb4, 0xdf, 0x63, 0x3d,
	}
	assert.Equal(address, nodeConf.Account.Address)
}
