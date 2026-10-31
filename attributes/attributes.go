package attributes

type Health struct {
	value int
}

type Armor struct {
	value int
}

type Attack struct {
	value int
}

type Point struct {
	X, Y int
}

func NewHealth(val int) *Health { return &Health{value: val} }

func (h *Health) GetValue() int { return h.value }

func (h *Health) Increase(amount int) {
	if amount > 0 {
		h.value += amount
	}
}

func (h *Health) Decrease(amount int) bool {
	h.value -= amount
	if h.value < 0 {
		h.value = 0
		return true
	}
	return false
}

func NewArmor(val int) *Armor { return &Armor{value: val} }

func (a *Armor) GetValue() int { return a.value }

func (a *Armor) ReduceDamage(damage int) int {
	res := damage - a.value
	if res < 0 {
		return 0
	}
	return res
}

func NewAttack(val int) *Attack { return &Attack{value: val} }

func (a *Attack) GetValue() int { return a.value }

func (a *Attack) CalculateDamage(targetArmor *Armor) int {
	return targetArmor.ReduceDamage(a.value)
}
