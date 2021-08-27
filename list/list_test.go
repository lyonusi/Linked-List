package list

// Basic imports
import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ListTestSuite struct {
	suite.Suite
	singleNodeValue   string
	emptyList         *linkedList
	listWithOneItem   *linkedList
	listWithTwoItems  *linkedList
	listWithManyItems *linkedList
}

// before each test
func (s *ListTestSuite) SetupTest() {

	s.singleNodeValue = "single"
	s.emptyList = &linkedList{}

	singleNode := &node{
		data: s.singleNodeValue,
		next: nil,
	}
	s.listWithOneItem = &linkedList{
		length: 1,
		head:   singleNode,
		tail:   singleNode,
	}

	s.listWithTwoItems = &linkedList{
		length: 2,
		head: &node{
			data: "1",
			next: singleNode,
		},
		tail: singleNode,
	}

	s.listWithManyItems = &linkedList{
		length: 4,
		head: &node{
			data: "1",
			next: &node{
				data: "2",
				next: &node{
					data: "3",
					next: singleNode,
				},
			}},
		tail: singleNode,
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestListListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *ListTestSuite) TestAddShouldAddNodeToHeadIfListIsEmpty() {
	nodeValue := "new"
	s.emptyList.Add(nodeValue)
	s.Assert().Equal(nodeValue, s.emptyList.head.data)
	s.Assert().Nil(s.emptyList.head.next)

	s.Assert().Equal(nodeValue, s.emptyList.tail.data)
	s.Assert().Nil(s.emptyList.tail.next)

	s.Assert().Equal(1, s.emptyList.length)
}

func (s *ListTestSuite) TestAddShouldAddNodeToTailIfListHasOneNode() {
	nodeValue := "new"
	s.listWithOneItem.Add(nodeValue)
	s.Assert().Equal(s.singleNodeValue, s.listWithOneItem.head.data)

	s.Assert().Equal(nodeValue, s.listWithOneItem.head.next.data)
	s.Assert().Equal(nodeValue, s.listWithOneItem.tail.data)
	s.Assert().Nil(s.listWithOneItem.tail.next)

	s.Assert().Equal(2, s.listWithOneItem.length)
}

func (s *ListTestSuite) TestAddShouldAddNodeToTailIfListHasManyNode() {
	nodeValue := "new"
	s.listWithManyItems.Add(nodeValue)
	s.Assert().Equal("1", s.listWithManyItems.head.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithManyItems.head.next.next.next.data)
	s.Assert().Equal(nodeValue, s.listWithManyItems.head.next.next.next.next.data)
	s.Assert().Equal(nodeValue, s.listWithManyItems.tail.data)
	s.Assert().Nil(s.listWithManyItems.tail.next)
	s.Assert().Equal(5, s.listWithManyItems.length)
}
