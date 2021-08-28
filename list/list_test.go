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

func (s *ListTestSuite) TestInsertAfterShouldReturnErrorIfListIsEmpty() {
	index := 0
	nodeValue := "new"
	s.Assert().EqualError(s.emptyList.InsertAfter(index, nodeValue), "[error: empty list]")
}
func (s *ListTestSuite) TestInsertAfterShouldReturnErrorIfIndexExceededLength() {
	index := 5
	nodeValue := "new"
	s.Assert().EqualError(s.listWithTwoItems.InsertAfter(index, nodeValue), "[invalid index: exceeded maximum index (1)]")
}
func (s *ListTestSuite) TestInsertAfterShouldReturnErrorIfIndexIsNegative() {
	index := -10
	nodeValue := "new"
	s.Assert().EqualError(s.listWithTwoItems.InsertAfter(index, nodeValue), "[invalid index: should be [0 - 1]]")
}
func (s *ListTestSuite) TestInsertAfterShouldAddNodeAtTailIfIndexEqualsLength() {
	nodeValue := "new"

	index := 0
	s.listWithOneItem.InsertAfter(index, nodeValue)
	s.Assert().Equal(nodeValue, s.listWithOneItem.tail.data)
	s.Assert().Equal(nodeValue, s.listWithOneItem.head.next.data)
	s.Assert().NotEqual(s.listWithOneItem.tail.data, s.listWithOneItem.head.data)

	index = 1
	s.listWithTwoItems.InsertAfter(index, nodeValue)
	s.Assert().Equal(nodeValue, s.listWithTwoItems.tail.data)
	s.Assert().Equal(nodeValue, s.listWithTwoItems.head.next.next.data)

	index = 3
	s.listWithManyItems.InsertAfter(index, nodeValue)
	s.Assert().Equal(nodeValue, s.listWithManyItems.tail.data)

	s.Assert().Equal("1", s.listWithManyItems.head.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithManyItems.head.next.next.next.data)
	s.Assert().Equal(nodeValue, s.listWithManyItems.head.next.next.next.next.data)
	s.Assert().Equal(nodeValue, s.listWithManyItems.tail.data)
	s.Assert().Nil(s.listWithManyItems.tail.next)
	s.Assert().Equal(5, s.listWithManyItems.length)
}

func (s *ListTestSuite) TestInsertAfterShouldAddNodeAfterIndexedNode() {
	nodeValue := "new"

	index := 0
	s.listWithTwoItems.InsertAfter(index, nodeValue)
	s.Assert().Equal(nodeValue, s.listWithTwoItems.head.next.data)

	index = 2
	s.listWithManyItems.InsertAfter(index, nodeValue)

	s.Assert().Equal("1", s.listWithManyItems.head.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithManyItems.head.next.next.next.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithManyItems.tail.data)
	s.Assert().Equal(nodeValue, s.listWithManyItems.head.next.next.next.data)
	s.Assert().Nil(s.listWithManyItems.tail.next)
	s.Assert().Equal(5, s.listWithManyItems.length)
}

func (s *ListTestSuite) TestPopShouldReturnErrorIfListIsEmpty() {
	result, err := s.emptyList.Pop()
	s.Assert().Equal("", result)
	s.Assert().EqualError(err, "[error: empty list]")
}

func (s *ListTestSuite) TestPopShouldReturnHeadIfListHasOneNode() {
	result, err := s.listWithOneItem.Pop()
	s.Assert().Equal("single", result)
	s.Assert().Nil(err)
	s.Assert().Equal(0, s.listWithOneItem.length)
	s.Assert().Equal(s.emptyList, s.listWithOneItem)
}

func (s *ListTestSuite) TestPopShouldReturnHeadIfListHasManyNodes() {
	result, err := s.listWithManyItems.Pop()
	s.Assert().Equal("1", result)
	s.Assert().Nil(err)
	s.Assert().Equal(3, s.listWithManyItems.length)
	s.Assert().Equal("2", s.listWithManyItems.head.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithManyItems.head.next.next.data)
	s.Assert().Nil(s.listWithManyItems.tail.next)
}

func (s *ListTestSuite) TesContainsShouldReturnFalseIfListDoesNotContainItem() {
	testValue := "a"
	s.Assert().Equal(false, s.emptyList.Contains(testValue))
	s.Assert().Equal(false, s.listWithOneItem.Contains(testValue))
	s.Assert().Equal(false, s.listWithTwoItems.Contains(testValue))
	s.Assert().Equal(false, s.listWithManyItems.Contains(testValue))

}

func (s *ListTestSuite) TestContainsShouldReturnTrueIfListDoesNotContainItem() {
	testValue := "single"
	s.Assert().Equal(true, s.listWithOneItem.Contains(testValue))
	s.Assert().Equal(true, s.listWithTwoItems.Contains(testValue))
	s.Assert().Equal(true, s.listWithManyItems.Contains(testValue))
}

func (s *ListTestSuite) TestIndexOfShouldReturnNegativeOneIfListDoesNotContainItem() {
	testValue := "a"
	s.Assert().Equal(-1, s.emptyList.IndexOf(testValue))
	s.Assert().Equal(-1, s.listWithOneItem.IndexOf(testValue))
	s.Assert().Equal(-1, s.listWithTwoItems.IndexOf(testValue))
	s.Assert().Equal(-1, s.listWithManyItems.IndexOf(testValue))
}

func (s *ListTestSuite) TestIndexOfShouldReturnIndexIfListContainsItem() {
	testValue := "single"
	s.Assert().Equal(0, s.listWithOneItem.IndexOf(testValue))
	s.Assert().Equal(1, s.listWithTwoItems.IndexOf(testValue))
	s.Assert().Equal(3, s.listWithManyItems.IndexOf(testValue))
}

func (s *ListTestSuite) TestPushShouldAddNodeToHead() {
	nodeValue := "new"

	s.emptyList.Push(nodeValue)
	s.Assert().Equal(nodeValue, s.emptyList.head.data)
	s.Assert().Nil(s.emptyList.head.next)
	s.Assert().Equal(nodeValue, s.emptyList.tail.data)
	s.Assert().Nil(s.emptyList.tail.next)
	s.Assert().Equal(1, s.emptyList.length)

	s.listWithOneItem.Push(nodeValue)
	s.Assert().Equal(nodeValue, s.listWithOneItem.head.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithOneItem.head.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithOneItem.tail.data)
	s.Assert().Equal(s.listWithOneItem.head.next, s.listWithOneItem.tail)
	s.Assert().Nil(s.listWithOneItem.tail.next)
	s.Assert().Equal(2, s.listWithOneItem.length)

	s.listWithTwoItems.Push(nodeValue)

	s.listWithManyItems.Push(nodeValue)
	s.Assert().Equal(nodeValue, s.listWithManyItems.head.data)
	s.Assert().Equal("1", s.listWithManyItems.head.next.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.next.next.data)
	s.Assert().Equal(s.singleNodeValue, s.listWithManyItems.head.next.next.next.next.data)
	s.Assert().Nil(s.listWithManyItems.tail.next)
	s.Assert().Equal(5, s.listWithManyItems.length)
}

func (s *ListTestSuite) TestRemoveShouldReturnErrorIfListDoesNotContainItem() {
	nodeValue := "0"
	s.Assert().EqualError(s.emptyList.Remove(nodeValue), "[error: \"0\" not found]")
	s.Assert().EqualError(s.listWithOneItem.Remove(nodeValue), "[error: \"0\" not found]")
	s.Assert().EqualError(s.listWithTwoItems.Remove(nodeValue), "[error: \"0\" not found]")
	s.Assert().EqualError(s.listWithManyItems.Remove(nodeValue), "[error: \"0\" not found]")
}

func (s *ListTestSuite) TestRemoveShouldRemoveItemIfListContainsItem() {
	nodeValue := "single"
	err := s.listWithOneItem.Remove(nodeValue)
	s.Assert().Equal(s.emptyList, s.listWithOneItem)
	s.Assert().Nil(err)
	s.Assert().Equal(0, s.listWithOneItem.length)

	err = s.listWithTwoItems.Remove(nodeValue)
	s.Assert().Nil(err)
	s.Assert().Equal("1", s.listWithTwoItems.head.data)
	s.Assert().Equal("1", s.listWithTwoItems.tail.data)
	s.Assert().Nil(s.listWithTwoItems.head.next)
	s.Assert().Nil(s.listWithTwoItems.tail.next)

	err = s.listWithManyItems.Remove(nodeValue)
	s.Assert().Nil(err)
	s.Assert().Equal("1", s.listWithManyItems.head.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal("3", s.listWithManyItems.tail.data)
	s.Assert().Equal(3, s.listWithManyItems.length)
	s.Assert().Nil(s.listWithManyItems.tail.next)
}

func (s *ListTestSuite) TestRemoveByIndexShouldReturnErrorIfIndexExceededLength() {
	index := 5
	s.Assert().EqualError(s.listWithTwoItems.RemoveByIndex(index), "[invalid index: exceeded maximum index (1)]")
}
func (s *ListTestSuite) TestRemoveByIndexShouldReturnErrorIfIndexIsNegative() {
	index := -10
	s.Assert().EqualError(s.listWithTwoItems.RemoveByIndex(index), "[invalid index: should be [0 - 1]]")
}

func (s *ListTestSuite) TestRemoveByIndexShouldReturnErrorIfListIsEmpty() {
	index := 0
	s.Assert().EqualError(s.emptyList.RemoveByIndex(index), "[error: empty list]")
}

func (s *ListTestSuite) TestRemoveByIndexShouldRemoveIndexedItem() {
	index := 0
	err := s.listWithOneItem.RemoveByIndex(index)
	s.Assert().Equal(s.emptyList, s.listWithOneItem)
	s.Assert().Nil(err)
	s.Assert().Equal(0, s.listWithOneItem.length)

	err = s.listWithTwoItems.RemoveByIndex(index)
	s.Assert().Nil(err)
	s.Assert().Equal("single", s.listWithTwoItems.head.data)
	s.Assert().Equal("single", s.listWithTwoItems.tail.data)
	s.Assert().Nil(s.listWithTwoItems.head.next)
	s.Assert().Nil(s.listWithTwoItems.tail.next)

	index = 3
	err = s.listWithManyItems.RemoveByIndex(index)
	s.Assert().Nil(err)
	s.Assert().Equal("1", s.listWithManyItems.head.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.data)
	s.Assert().Equal("3", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal("3", s.listWithManyItems.tail.data)
	s.Assert().Equal(3, s.listWithManyItems.length)
	s.Assert().Nil(s.listWithManyItems.tail.next)
}

func (s *ListTestSuite) TestSetShouldReturnErrorIfIndexExceededLength() {
	index := 5
	nodeValue := "0"
	s.Assert().EqualError(s.listWithTwoItems.Set(index, nodeValue), "[invalid index: exceeded maximum index (1)]")
}
func (s *ListTestSuite) TestSetShouldReturnErrorIfIndexIsNegative() {
	index := -10
	nodeValue := "0"

	s.Assert().EqualError(s.listWithTwoItems.Set(index, nodeValue), "[invalid index: should be [0 - 1]]")
}

func (s *ListTestSuite) TestSetShouldReturnErrorIfListIsEmpty() {
	index := 0
	nodeValue := "0"

	s.Assert().EqualError(s.emptyList.Set(index, nodeValue), "[error: empty list]")
}

func (s *ListTestSuite) TestSetShouldSetItemAsHeadAndTailIfListHasOneNode() {

	index := 0
	nodeValue := "0"
	err := s.listWithOneItem.Set(index, nodeValue)
	s.Assert().Nil(err)
	s.Assert().Equal("0", s.listWithOneItem.head.data)
	s.Assert().Equal("0", s.listWithOneItem.tail.data)
	s.Assert().Nil(s.listWithOneItem.head.next)
	s.Assert().Nil(s.listWithOneItem.tail.next)
	s.Assert().Equal(1, s.listWithOneItem.length)
}

func (s *ListTestSuite) TestSetShouldSetItemAsHeadOrTailIfListHasTwoNodes() {

	nodeValue := "0"

	index := 0
	err := s.listWithTwoItems.Set(index, nodeValue)
	s.Assert().Nil(err)
	s.Assert().Equal("0", s.listWithTwoItems.head.data)
	s.Assert().Equal("single", s.listWithTwoItems.head.next.data)
	s.Assert().Equal("single", s.listWithTwoItems.tail.data)
	s.Assert().Nil(s.listWithTwoItems.tail.next)
	s.Assert().Equal(2, s.listWithTwoItems.length)

	index = 1
	err = s.listWithTwoItems.Set(index, nodeValue)
	s.Assert().Nil(err)
	s.Assert().Equal("0", s.listWithTwoItems.head.data)
	s.Assert().Equal("0", s.listWithTwoItems.head.next.data)
	s.Assert().Equal("0", s.listWithTwoItems.tail.data)
	s.Assert().Nil(s.listWithTwoItems.tail.next)
	s.Assert().Equal(2, s.listWithTwoItems.length)
}

func (s *ListTestSuite) TestSetShouldSetItemAsIndexedIfListHasManyNodes() {

	nodeValue := "0"
	index := 2
	err := s.listWithManyItems.Set(index, nodeValue)
	s.Assert().Nil(err)
	s.Assert().Equal("1", s.listWithManyItems.head.data)
	s.Assert().Equal("2", s.listWithManyItems.head.next.data)
	s.Assert().Equal("0", s.listWithManyItems.head.next.next.data)
	s.Assert().Equal("single", s.listWithManyItems.head.next.next.next.data)
	s.Assert().Equal("single", s.listWithManyItems.tail.data)
	s.Assert().Equal(4, s.listWithManyItems.length)
	s.Assert().Nil(s.listWithManyItems.tail.next)
}
