package mgr

import (
	"testing"

	"github.com/alibaba/pouch/apis/types"

	"github.com/stretchr/testify/assert"
)

func TestValidateNvidiaDevice(t *testing.T) {
	r := &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaVisibleDevices: "all",
		},
	}
	assert.NoError(t, validateNvidiaDevice(r))

	r = &types.Resources{
		NvidiaConfig: nil,
	}
	assert.NoError(t, validateNvidiaDevice(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaVisibleDevices: "none",
		},
	}
	assert.NoError(t, validateNvidiaDevice(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaVisibleDevices: "void",
		},
	}
	assert.NoError(t, validateNvidiaDevice(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaVisibleDevices: "",
		},
	}
	assert.NoError(t, validateNvidiaDevice(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaVisibleDevices: "ErrorError",
		},
	}
	assert.Error(t, validateNvidiaDevice(r))
}

func TestValidateNvidiaDriver(t *testing.T) {
	r := &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaDriverCapabilities: "all",
		},
	}
	assert.NoError(t, validateNvidiaDriver(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaDriverCapabilities: "",
		},
	}
	assert.NoError(t, validateNvidiaDriver(r))

	r = &types.Resources{
		NvidiaConfig: nil,
	}
	assert.NoError(t, validateNvidiaDriver(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaDriverCapabilities: "compute,compat32,graphics,utility,video,display",
		},
	}
	assert.NoError(t, validateNvidiaDriver(r))

	r = &types.Resources{
		NvidiaConfig: &types.NvidiaConfig{
			NvidiaDriverCapabilities: "ErrorError",
		},
	}
	assert.NoError(t, validateNvidiaDriver(r))
}
