package scanner

type Scanner interface {
	Scan()
}

type ShiroTarget struct {
	Base    string
	Key     string
	Gadget  string
	Command string
}

func InitFactory(module string) func(target *ShiroTarget) Scanner {
	return func(target *ShiroTarget) Scanner {
		switch module {
		case "key":
			return &KeyScanner{Target: target}
		case "gadgetfuzz":
			return &GadgetScanner{Target: target}
		case "gadgetexec":
			return nil
		default:
			return nil
		}
	}
}
