package stage

/**
 *
 * This is an experimental implementation for Stage for Google Go
 * So far only the Entity Class is implemented
 * Need to include the rest
 */

// TODO: Still needs some functionality to preserve the oreder in the hash map
type Entity struct {
	script string
	links  map[string]*Entity
	hidden map[string]bool
}

// Creating an empty entity
func NewEntity() (e *Entity) {
	e = new(Entity)
	e.script = ""
	e.links = make(map[string]*Entity)
	e.hidden = make(map[string]bool)
	return
}

func NewEntityWithScript(script string) (e *Entity) {
	e = NewEntity()
	e.SetScript(script)
	return
}

func (e *Entity) IsAction() bool {
	// Checking if the script is empty
	return e.script != ""
}

func (e *Entity) Link(name string, target *Entity, isVisible bool) {
	// Linking the entity
	e.links[name] = target
	// There is some inconsistency regarding this
	e.hidden[name] = !isVisible
}

func (e *Entity) Unlink(name string) {
	delete(e.links, name)
	delete(e.hidden, name)
}

func (e *Entity) Find(name string, isPlayer bool) (result *Entity) {
	result = nil
	if isPlayer {
		if !e.hidden[name] {
			result = e.links[name]
		}
	} else {
		if name == "this" || name == "here" {
			result = e
		} else {
			result = e.links[name]
		}
	}
	return
}

// Returning the copy of the links
// This one needs to preserve the order somehow...
func (e *Entity) Links() []string {
	size := len(e.links)
	links := make([]string, size)
	count := 0
	for key := range e.links {
		// For some reason it is impossible to write count++
		links[count] = key
		count++
	}
	return links
}

func (e *Entity) Hidden(name string) bool {
	present := e.hidden[name]
	return present
}

func (e *Entity) Reveal(name string) {
	//delete(e.hidden, name)
	e.hidden[name] = false
}

func (e *Entity) Hide(name string) {
	e.hidden[name] = true
}

// Make we should make entities inmuttable?
func (e *Entity) SetScript(script string) {
	e.script = script
}

func (e *Entity) Script() string {
	script := e.script
	return script
}
