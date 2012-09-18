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
	e1             = NewEntity()
	script         = "Some script"
	e2             = NewEntityWithScript(script)
	here           = "here"
	this           = "this"
	name           = "name"
	n      *Entity = nil
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

	// Testing entitiy constructor
	test(t, "Basic Entity Script", e1.Script(), "")
	test(t, "Basic Action", e2.Script(), script)

	// Testing if Action works fine
	test(t, "Entity is not Action", e1.IsAction(), false)
	test(t, "Action is Action", e2.IsAction(), true)

	// Testing the basic connections
	test(t, "Player found something", n, e1.Find("something", true))
	test(t, "Imp found something", n, e1.Find("something", false))

	e1.Link(name, e2, true)

	//
	test(t, "Player found a link", e1.Find(name, true), e2)
	test(t, "Link is hidden", e1.Hidden(name), false)

	e1.Hide(name)
	test(t, "Link is now hidden 1", e1.Hidden(name), true)
	test(t, "Player cannot find it", e1.Find(name, true), n)
	test(t, "Imp can find it", e1.Find(name, false), e2)

	e1.Reveal(name)
	test(t, "Link is now visible", e1.Hidden(name), false)

	// Readding test
	e1.Link(name, e2, false)
	test(t, "Link is now hidden", e1.Hidden(name), true)
	test(t, "Player cannot find it", e1.Find(name, true), n)
	test(t, "Imp can find it", e1.Find(name, false), e2)

	links := e1.Links()
	test(t, "Linkin6", len(links), 1)
	test(t, "Linkin7", links[0], name)

	e1.Unlink(name)
	links = e1.Links()

	test(t, "Unlinking works for player", e1.Find(name, true), n)
	test(t, "Unlinking works for imp", e1.Find(name, false), n)

	test(t, "Symmetry for player", e1.Find(name, true), e1.Find(name, true))
	test(t, "Symmetry for imp", e1.Find(name, false), e1.Find(name, false))

	test(t, "Veryfying the links", len(links), 0)
	// This one is not exactly correct
	test(t, "Empty link isHidden test", e1.Hidden(name), false)

	// TODO: Add tests for here and this
	test(t, "Here is not available to player", e1.Find(here, true), n)
	test(t, "This is not available to player", e1.Find(this, true), n)

	test(t, "Here is available to imp", e1.Find(here, false), e1)
	test(t, "This is available to imp", e1.Find(this, false), e1)
}
