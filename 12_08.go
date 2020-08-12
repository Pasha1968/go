package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}
type Item struct {
	Value interface {
	}
	Created    time.Time
	Expiration int64
}

func New(defaultExpiration, cleanupInterval time.Duration) *Cache {

	// инициализируем карту(map) в паре ключ(string)/значение(Item)
	items := make(map[string]Item)

	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	// Если интервал очистки больше 0, запускаем GC (удаление устаревших элементов)
	if cleanupInterval > 0 {
		cache.StartGC() // данный метод рассматривается ниже
	}

	return &cache
}
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {

	var expiration int64

	// Если продолжительность жизни равна 0 - используется значение по-умолчанию
	if duration == 0 {
		duration = c.defaultExpiration
	}

	// Устанавливаем время истечения кеша
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()

	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}
func (c *Cache) Delete(key string) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("Key not found")
	}

	delete(c.items, key)

	return nil
}
func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {

	for {
		// ожидаем время установленное в cleanupInterval
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		// Ищем элементы с истекшим временем жизни и удаляем из хранилища
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}

}

func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}
func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}

	}

	return item.Value, true
}

func (c *Cache) Opros() map[int]interface{} {
	c.RLock()
	i := 0
	m := make(map[int]interface{})
	defer c.RUnlock()
	for key := range c.items {
		temp, _ := c.Get(key)
		//fmt.Println(temp)
		m[i] = temp
		i++
	}
	//fmt.Println(m)
	return m
}
func (c *Cache) chamform(ch chan map[int]interface{}) {
	ch <- c.Opros()
}
func (c *Cache) cham(ch chan map[int]interface{}, cashe *Cache) {
	for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
		msg = cashe.Opros()
		fmt.Println(msg)
	}
}

func (c *Cache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}

type person struct {
	age     int
	name    string
	surname string
}

// type worker struct {
// 	man *person
// 	//	jobname  string
// 	location string
// }
type worker interface {
	hello()
	take()
	// copy()
}

type butcher struct {
	creature *person
	needed   bool
	shop     string
	jobname  string
}
type doctor struct {
	//creature worker
	creature *person
	type_    string
	hospital string
	killed   int
	healed   int
	jobname  string
}
type slave struct {
	creature person
	type_    string
	days     int
}
type slaveboss struct {
	creature   person
	type_      string
	Department string
}

func Newbutcher(cre *person) *butcher {
	return &butcher{
		creature: cre,
		needed:   true,
		shop:     "",
		jobname:  "Мясничок",
	}
}
func Newdoctor(cre *person) *doctor {
	return &doctor{
		creature: cre,
		type_:    "Unknown",
		hospital: "Unknown",
		killed:   0,
		healed:   0,
		jobname:  "doctor",
	}
}
func (b butcher) hello() {
	fmt.Println("hello I'm")
	fmt.Println(b.creature.name)
	fmt.Println(b.creature.surname)
	fmt.Println(b.jobname)
}
func (b butcher) copy() {
	fmt.Println("hello I'm")
	fmt.Println(b.creature.name)
	fmt.Println(b.creature.surname)
	fmt.Println(b.jobname)
}
func (b butcher) take() (ex string) {
	ex = "hello I'm" + b.jobname
	return
}
func (b doctor) hello() {
	fmt.Println("Hallo I'm")
	fmt.Println(b.creature.name)
	fmt.Println(b.creature.surname)
	fmt.Println(b.jobname)
}
func (b doctor) copy(ent doctor) (exit doctor) {
	exit = ent
	return exit
}
func konn(g worker) {
	fmt.Println(g)
	g.hello()
}

func main() {
	cashe := New(0, 0)
	casheboss := New(0, 0)
	// man := &person{23, "Mihail", "Mihalovich"}
	// woman := &person{20, "Sasha", "Oleksandrovna"}
	// b := butcher{man, true, "silpo"}
	// b := Newbutcher(man)
	// a := Newdoctor(woman)
	// a.killed = 50
	// cashe.Set("doc", a, 0)
	// cashe.Set("but", b, 0)
	//fmt.Println(cashe)
	//i, _ := cashe.Get("doc")
	var slave1 = slave{
		type_: "batrak",
		days:  20,
		creature: person{
			age:     20,
			name:    "orc",
			surname: "orcovich",
		},
	}
	var slave2 = slave{
		type_: "wisp",
		days:  159,
		creature: person{
			age:     200,
			name:    "wisp",
			surname: "wispovich",
		},
	}
	var slave3 = slaveboss{
		type_:      "wisp",
		Department: "Org",
		creature: person{
			age:     150,
			name:    "ocg",
			surname: "ocgovich",
		},
	}
	cashe.Set("slave", slave1, 0)
	cashe.Set("builder", slave2, 0)
	casheboss.Set("boss1", slave3, 0)
	//cashe.Opros()
	//go casheboss.Opros()
	//go cashe.Opros()
	var c chan map[int]interface{} = make(chan map[int]interface{})

	go cashe.chamform(c)
	go cashe.cham(c, casheboss)
	var input string
	fmt.Scanln(&input)
	//fmt.Println(mapa)
	//fmt.Println(cashe.items["slave"])
}
