package justitia

import (
	"github.com/DSiSc/justitia/node"
	"github.com/DSiSc/txpool/common/log"
)

func main() {
	node, err := node.NewNode()
	if nil != err {
		log.Error("Failed to instance a node.")
	}

	err = node.Start()
	if nil != err {
		log.Error("Failed to start node service.")
	}
}
