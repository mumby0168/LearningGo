package newsfeed

type IAdder interface {
	Add(Item item)
}

type IGetter interface {
	GetAll()
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (repo *Repo) Add(item Item) {
	repo.Items = append(repo.Items, item)
}

func (repo *Repo) GetAll() []Item {
	return repo.Items
}
