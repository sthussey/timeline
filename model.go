package timeline

type Timeline struct {
	syncMap   map[string]func(interface{}, map[string]interface{}) error
	actionMap map[string]func(interface{}, map[string]interface{}) error
	Events    []Event
	Variables map[string]interface{}
}

type Event struct {
	Sync   string
	Action string
	Inputs interface{}
	Block  bool
}

func NewTimeline() *Timeline {
	t := &Timeline{Events: make([]Event, 0), Variables: make(map[string]interface{})}
	t.setDefaultMaps()
	return t
}

func (t *Timeline) Execute() {
	for _, e := range t.Events {
		var f func(interface{}, map[string]interface{}) error
		var ok bool

		if e.Sync != "" {
			f, ok = t.syncMap[e.Sync]
		} else if e.Action != "" {
			f, ok = t.actionMap[e.Action]
		}

		if !ok {
			continue
		}

		if e.Block {
			_ = f(e.Inputs, t.Variables)
		} else {
			go f(e.Inputs, t.Variables)
		}
	}
}

func (t *Timeline) AddEvent(e Event) {
	t.Events = append(t.Events, e)
}

func (t *Timeline) SetVariable(name string, val interface{}) {
	t.Variables[name] = val
}

func (t *Timeline) GetVariable(name string) (interface{}, bool) {
	v, ok := t.Variables[name]
    return v, ok
}

func (t *Timeline) MapSync(sync string, syncer func(interface{}, map[string]interface{}) error) {
	t.syncMap[sync] = syncer
}

func (t *Timeline) MapAction(action string, actor func(interface{}, map[string]interface{}) error) {
	t.actionMap[action] = actor
}

func (t *Timeline) setDefaultMaps() {
	t.syncMap = getDefaultSyncMap()
	t.actionMap = getDefaultActionMap()
}
