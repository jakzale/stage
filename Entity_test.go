package stage

/**
 *
 * This is a set of Entity tests
 *
 */

import (
	"testing"
	//"reflect"
)

var (
	e1     = NewEntity()
	script = "Some script"
	e2     = NewEntityWithScript(script)
)

// Basic testing function!!
// Okay, this ones will never be nil
func test(t *testing.T, name string, a, b interface{}) {
	// Turns out Google go has difficulties with comparing nils, becuase there
	// are no nil interfaces -_-
	// This seem like a solution, although it is pretty ugly

	if a != b {
		t.Errorf(name+" %s want %s", a, b)
	}
}

// Need to make test with the following values
// Expected x, resulted y
func TestEntityCreation(t *testing.T) {

	// Empty entity
	var n *Entity = nil

	// Testing entitiy constructor
	test(t, "Entity1", e1.Script(), "")
	test(t, "Entity2", e2.Script(), script)

	// Testing if Action works fine
	test(t, "Action1", e1.IsAction(), false)
	test(t, "Action2", e2.IsAction(), true)

	// Testing the basic connections
	test(t, "Linkin1", n, e1.Find("something"))

	name := "name"
	e1.Link(name, e2)
	test(t, "Linkin2", e1.Find(name), e2)
	test(t, "Linkin3", e1.Hidden(name), false)

	e1.Hide(name)
	test(t, "Linkin4", e1.Hidden(name), true)

	e1.Reveal(name)
	test(t, "Linkin5", e1.Hidden(name), false)

	links := e1.Links()
	test(t, "Linkin6", len(links), 1)
	test(t, "Linkin7", links[0], name)

	e1.Unlink(name)
	links = e1.Links()

	test(t, "Linkin8", e1.Find(name), n)
	test(t, "Linkin8a", e1.Find(name), e1.Find(name))
	test(t, "Linkin9", len(links), 0)
	// This one is not exactly correct
	test(t, "Linkin0", e1.Hidden(name), false)

}
