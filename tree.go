// tree
package tree

type Tree struct {
	children      [20]*Tree
	data          []interface{}
	isRoot        bool
	totalCpuValue int
	totalRamValue int
}

// NewTree returns a new Tree for the provided entry.
func NewTree(d []interface{}, pIsroot bool, totalCpu int, totalRam int, nassigned string) *Tree {
	return &Tree{
		data:          d,
		children:      [20]*Tree{},
		isRoot:        pIsroot,
		totalCpuValue: totalCpu,
		totalRamValue: totalRam,
	}
}

func (node *Tree) IsRoot() bool {
	return node.isRoot
}

func (node *Tree) Data() []interface{} {
	return node.data
}

func (node *Tree) TotalCpuValue() int {
	return node.totalCpuValue
}

func (node *Tree) TotalRamValue() int {
	return node.totalRamValue
}

func (node *Tree) SetTotalCpuValue(value int) {
	node.totalCpuValue = value
}

func (node *Tree) SetTotalRamValue(value int) {
	node.totalRamValue = value
}

func (node *Tree) Children() [20]*Tree {
	return node.children
}

func (node *Tree) SetChildren(pos int, child *Tree) {
	node.children[pos] = child
}

func (node *Tree) SetData(newData []interface{}) {
	node.data = newData
}

func (node *Tree) IsLeaf() bool {
	return len(node.children) == 0
}

func (node *Tree) Preorder(listado *[][]interface{}) [][]interface{} {
	*listado = append(*listado, node.Data())
	for i := 0; i < len(node.children); i++ {
		if node.children[i] != nil {
			node.children[i].Preorder(listado)
		}
	}
	return *listado
}
