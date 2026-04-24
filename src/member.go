package src

type Member struct {
  abs float32
  sign float32
}

func CreateMember(str string) (*Member, error) {
  return nil, nil
}

func (m Member) GetValue() float32 {
  return m.abs * m.sign
}

func (m Member) GetSign() string {
  if m.sign < 0 {
    return "-"
  }
  return "+"
}
