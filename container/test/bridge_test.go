package container

import (
	"github.com/mikeyfennelly1/mrund/container"
	"testing"
)

func TestBridge(t *testing.T) {
	t.Run("Just run", func(t *testing.T) {
		container.CreateVethPairAndBridge()
	})
}
