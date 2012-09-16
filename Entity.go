package stage

/**
 *
 * This is an experimental implementation for Stage for Google Go
 * So far only the Entity Class is implemented
 * Need to include the rest
 */

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

func (e *Entity) Link(name string, target *Entity) {
	// Linking the entity
	e.links[name] = target
	// Removing it from the hidden
	// In case of creating new links
	delete(e.hidden, name)

	return
}

func (e *Entity) Unlink(name string) {
	delete(e.links, name)
	delete(e.hidden, name)
}

func (e *Entity) Find(name string) *Entity {
	return e.links[name]
}

// Returning the copy of the links
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
	delete(e.hidden, name)
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
