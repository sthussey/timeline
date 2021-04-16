package timeline

type Timeline struct {
	syncMap   map[string]func(interface{}) error
	actionMap map[string]func(interface{}) error
	Events    []Event
}

type Event struct {
	Sync   string
	Action string
	Inputs interface{}
	Block  bool
}

func NewTimeline() *Timeline {
	t := &Timeline{Events: make([]Event, 0)}
	t.setDefaultMaps()
	return t
}

func (t *Timeline) AddEvent(e Event) {
	t.Events = append(t.Events, e)
}

func (t *Timeline) MapSync(sync string, syncer func(interface{}) error) {
	t.syncMap[sync] = syncer
}

func (t *Timeline) MapAction(action string, actor func(interface{}) error) {
	t.actionMap[action] = actor
}

func (t *Timeline) setDefaultMaps() {
	t.syncMap = getDefaultSyncMap()
	//	t.actionMap = getDefaultActionMap()
}
