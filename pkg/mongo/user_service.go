package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"api-boilerplate/pkg"
)

// UserService : user service struct 
type UserService struct {
	collection *mgo.Collection
	crypto 		root.Crypto
}

// NewUserService creating a new user service
// The constructor gets a collection from the 
// session parameter and sets up the user index.
func NewUserService(session *Session, dbName string, collectionName string, crypto root.Crypto) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, crypto}
}

// CreateUser function is responsible for 
// creating a new user model and insert into the db
func (p *UserService) CreateUser(u *root.User) error {
	user := newUserModel(u)
	hashedPassword, err := p.crypto.Generate(user.Password)

	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return p.collection.Insert(&user)


}

// GetByUsername is responsible for getting a specific 
// user name and return the user model and potential error
func(p *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username":username}).One(&model)
	return model.toRootUser(), err
}

