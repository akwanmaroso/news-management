package interfaces

import "github.com/akwanmaroso/news/utils/mock"

var (
	newsApp mock.NewsAppInterface
	tagApp  mock.TagAppInterface

	newNewsApp = NewNews(&newsApp, &tagApp)
	newTagApp  = NewTag(&tagApp, &newsApp)
)
