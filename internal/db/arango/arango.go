package arango

import (
	"context"
	"github.com/arangodb/go-driver"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Query(qry string, bindVars map[string]interface{}) (driver.Cursor, error) {
	db := getDbInstance()
	ctx := context.Background()
	return db.Query(ctx, qry, bindVars)
}

func (c *Client) CreateDoc(collName string, doc interface{}) (string, error) {
	ctx := context.Background()
	coll := getCollectionInstance(collName)

	meta, err := coll.CreateDocument(ctx, doc)
	return meta.Key, err
}

func (c *Client) CreateDocs(collectionName string, doc interface{}) (driver.DocumentMetaSlice, error) {
	ctx := context.Background()
	coll := getCollectionInstance(collectionName)

	metas, _, err := coll.CreateDocuments(ctx, doc)
	return metas, err
}

func (c *Client) FindById(collName, id string, result interface{}) error {
	qry := "RETURN DOCUMENT(CONCAT(@collName, '/', @_key))"
	bindVars := map[string]interface{}{
		"collName": collName,
		"_key":     id,
	}

	cur, err := c.Query(qry, bindVars)
	if err != nil {
		return err
	}
	defer cur.Close()

	if _, readErr := cur.ReadDocument(nil, result); readErr != nil {
		return readErr
	}
	return nil
}

func (c *Client) UpdateDoc(collName string, id string, doc interface{}) (string, error) {
	ctx := context.Background()
	coll := getCollectionInstance(collName)

	exist, err := coll.DocumentExists(ctx, id)
	if err != nil {
		return "", err
	}

	var meta driver.DocumentMeta
	var updErr error

	if exist {
		meta, updErr = coll.UpdateDocument(ctx, id, doc)
	} else {
		meta, updErr = coll.CreateDocument(ctx, doc)
	}

	return meta.Key, updErr
}
func (c *Client) UpdateDocs(collName string, ids []string, doc interface{}) error {
	ctx := context.Background()
	coll := getCollectionInstance(collName)

	_, _, err := coll.UpdateDocuments(ctx, ids, doc)
	return err
}

func (c *Client) DeleteDoc(collName string, id string) error {
	ctx := context.Background()
	coll := getCollectionInstance(collName)
	_, err := coll.RemoveDocument(ctx, id)
	return err
}

func (c *Client) DeleteDocs(collName string, ids []string) error {
	ctx := context.Background()
	coll := getCollectionInstance(collName)
	_, _, err := coll.RemoveDocuments(ctx, ids)
	return err
}
