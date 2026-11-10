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

type Resource struct {
	value int
}

type UnitCost struct {
	Wood, Iron int
}

var Costs = map[string]UnitCost{
	"swordsman":    {1, 0},
	"spearman":     {1, 1},
	"light":        {1, 0},
	"heavy":        {1, 2},
	"cataphract":   {3, 1},
	"horse_archer": {2, 2},
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

func NewResource(val int) *Resource { return &Resource{value: val} }

func (w *Resource) GetValue() int { return w.value }

func (w *Resource) Increase(amount int) {
	if amount > 0 {
		w.value += amount
	}
}

func (w *Resource) Decrease(amount int) {
	w.value -= amount
	if w.value < 0 {
		w.value = 0
	}
}
