/*
type OrderedMap struct { ... }

func NewOrderedMap() OrderedMap                      // создать упорядоченный словарь
func (m *OrderedMap) Insert(key, value int)          // добавить элемент в словарь
func (m *OrderedMap) Erase(key int)                  // удалить элемент из словари
func (m *OrderedMap) Contains(key int) bool          // проверить существование элемента в словаре
func (m *OrderedMap) Size() int                      // получить количество элементов в словаре
func (m *OrderedMap) ForEach(action func(int, int))  // применить функцию к каждому элементу словаря от меньшего к большему
*/
package lesson_4

type OrderedMap struct {
	size int
	tree *Tree
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		size: 0,
		tree: nil,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	m.tree = m.tree.Insert(key)
}

func (m *OrderedMap) Erase(key int) {
	m.tree = m.tree.Delete(key)
}

func (m *OrderedMap) Contains(key int) bool {
	return m.tree.Contains(key)
}

func (m *OrderedMap) Size() int {
	return m.tree.Count()
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.tree.InorderTraversal(action)
}
