package main

type storage interface {
	Get() (any, error)
	Create() (any, error)
}

type notifier struct {
	s storage
}

func newNotifier(s storage) notifier {
	return notifier{s: s}
}

func (n notifier) getOrCreate() (any, error) {
	val, err := n.s.Get()
	if err != nil {
		newVal, err := n.s.Create()
		if err != nil {
			return nil, err
		}
		return newVal, nil
	}

	return val, nil
}
