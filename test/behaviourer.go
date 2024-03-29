package test

import (
	"fmt"
	"github.com/995933447/gobehaviortree"
)

//Behaviourer behaviourer
type Behaviourer struct {
	executor *gobehaviortree.Tree
}

//Init init
func (b *Behaviourer) Init() {
	b.executor = gobehaviortree.NewTree(nil, func(tree *gobehaviortree.Tree, result gobehaviortree.Result, root gobehaviortree.BaseNode) {
		fmt.Printf("node:%s run over!\n", root.WhoAmI())
	})

	action1 := NewAction1()
	action2 := NewAction2()

	selector := gobehaviortree.NewSelector()

	seq := gobehaviortree.NewSequence()
	seq.AddNodes(action2, action1)

	selector.AddNodes(action2, seq, action1)

	b.executor.SetRoot(selector)
}

//Run run
func (b *Behaviourer) Run() {
	b.executor.Run()
}

func NewBehaviour() *Behaviourer {
	return &Behaviourer{}
}
