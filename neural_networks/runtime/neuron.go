package runtime

type NeuronType int

const (
	Unknown NeuronType = iota
	Input
	Hidden
	Output

	kInputNodeBias = 0.0
)

var (
	kSigmoidTable []float64
)

type Neuron struct {
	nodeType     NeuronType
	id           int
	bias         float64
	activation   float64
	output       float64
	fanIn        []*Neuron
	fanInWeights []float64
	frwdDirty    bool
}

func New() *Neuron {
	n := &Neuron{}
	n.Init()
	return n
}

func (n *Neuron) Clear() {
	n.frwdDirty = true
}

func (n *Neuron) Init() {
	n.nodeType = Unknown
	n.id = -1
	n.bias = 0
	n.activation = 0
	n.output = 0
	n.fanIn = nil
	n.fanInWeights = nil
	n.frwdDirty = false
}

// TODO(ttacon): read from buffer file

// TODO(ttacon): remove fromCnt???
func (n *Neuron) AddFromConnection(neuronVec []*Neuron, wtsOffset []float64, fromCnt int) {
	for i, neur := range neuronVec {
		n.fanIn = append(n.fanIn, neur)
		n.fanInWeights = append(n.fanInWeights, wtsOffset[i])
	}
}

func (n *Neuron) setNodeType(nType NeuronType) {
	n.nodeType = nType
}

func (n *Neuron) FeedForward() {
	if !n.frwdDirty {
		return
	}

	if n.nodeType != Input {
		fanInCnt = len(n.fanIn)
		n.activation = 0

		for i, neur := range n.fanIn {
			if neur.frwdDirty {
				neur.FeedForward()
			}
			n.activation += (n.fanInWeights[i] * neur.output)
		}
		// sigmoid it
		n.output = Sigmoid(n.activation)
	}
	n.frwdDirty = false
}

func (n *Neuron) output() float64 {
	return n.output
}

func (n *Neuron) setOutput(outVal float64) {
	n.output = outVal
}

func (n *Neuron) Id() int {
	return n.id
}

func (n *Neuron) fanInCnt() int {
	return len(n.fanIn)
}

func (n *Neuron) fanInAt(idx int) *Neuron {
	return n.fanIn[idx]
}

func (n *Neuron) fanInWtsAt(idx int) float64 {
	return n.fanInWeights[idx]
}

func (n *Neuron) setId(idVal int) {
	n.id = idVal
}

func (n *Neuron) Bias() float64 {
	return n.Bias
}

func (n *Neuron) NodeTypes() NeuronType {
	return n.nodeType
}

func Sigmoid(activation float64) float64 {
	if n.activation <= -10.0 {
		return 0.0
	} else if n.activation >= 10.0 {
		return 1.0
	} else {
		return kSigmoidTable[100.0*(n.activation+10.0)]
	}
}
