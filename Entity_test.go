package stage

/**
 * Testing for the Entity class
 */

import (
	"fmt"
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

func format(a, b interface{}) string {
	return fmt.Sprintf("got %s want %s", a, b)
}

func test(t *testing.T, name string, a, b interface{}) {
	// Turns out Google go has difficulties with comparing nils, becuase there
	// are no nil interfaces -_-
	// This seem like a solution, although it is pretty ugly

	if a != b {
		t.Errorf(name+" %s want %s", a, b)
	}
}

// Splitting the test cases into seperate ones
func TestEntityCreation_001(t *testing.T) {
	a, b := e1.script, ""
	if a != b {
		t.Errorf(format(a, b))
	}
}

func TestEntityCreation_002(t *testing.T) {
	a, b := e2.script, script
	if a != b {
		t.Errorf(format(a, b))
	}
}

func TestEntityCreation_003(t *testing.T) {
	a, b := e1.IsAction(), false
	if a != b {
		t.Errorf(format(a, b))
	}
}

func TestEntityCreation_004(t *testing.T) {
	a, b := e2.IsAction(), true
	if a != b {
		t.Errorf(format(a, b))
	}
}

func TestEntityCreation_005(t *testing.T) {
	a, b := e1.Find("something", true), n
	if a != b {
		t.Errorf(format(a, b))
	}
}

func TestEntityCreation_006(t *testing.T) {
	a, b := e1.Find("something", false), n
	if a != b {
		t.Errorf(format(a, b))
	}
}

// The following test need some kind of test preparation
// Leave it for now
func TestEntityCreation(t *testing.T) {

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
