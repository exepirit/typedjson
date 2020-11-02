package options

type Identify struct {
	counter int
}

func (i *Identify) Mutate(object map[string]interface{}) {
	object["$id"] = i.counter
	i.counter++
}
