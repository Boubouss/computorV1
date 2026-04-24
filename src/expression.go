package src

type Expression struct {
  members [3]Member
  disc    float32
}

func ParseExpression(str string) (*Expression, error) {
  return nil, nil
}

func (e Expression) PrintReducedForm() {}
func (e Expression) PrintPolDegree() {}
func (e Expression) PrintSolution() {}

func (e *Expression) addMember() {}
func (e *Expression) calcDisc() {}
