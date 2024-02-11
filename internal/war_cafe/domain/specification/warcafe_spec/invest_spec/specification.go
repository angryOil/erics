package invest_spec

type Spec interface {
	IsSatisfy() bool
}

var _ Spec = (*AndSpecification)(nil)

type AndSpecification struct {
	specs []Spec
}

func NewAndSpecification(specs ...Spec) Spec {
	return &AndSpecification{
		specs: specs,
	}
}

func (a AndSpecification) IsSatisfy() bool {
	for _, spec := range a.specs {
		if !spec.IsSatisfy() {
			return false
		}
	}
	return true
}

var _ Spec = (*MaxSum)(nil)

type MaxSum struct {
	CurrentCount int
	AddCount     int
	MaxSumNumber int
}

func (m MaxSum) IsSatisfy() bool {
	if m.CurrentCount+m.AddCount > m.MaxSumNumber {
		return false
	}
	return true
}

var _ Spec = (*MinAddNumber)(nil)

type MinAddNumber struct {
	AddCount  int
	MinNumber int
}

func (m MinAddNumber) IsSatisfy() bool {
	if m.AddCount < m.MinNumber {
		return false
	}
	return true
}

var _ Spec = (*MaxAddNumber)(nil)

type MaxAddNumber struct {
	AddCount  int
	MaxNumber int
}

func (m MaxAddNumber) IsSatisfy() bool {
	if m.AddCount > m.MaxNumber {
		return false
	}
	return true
}
