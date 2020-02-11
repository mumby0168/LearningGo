package newsfeed

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
