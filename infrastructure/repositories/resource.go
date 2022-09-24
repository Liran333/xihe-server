package repositories

import "github.com/opensourceways/xihe-server/domain"

type ResourceListDO struct {
	Name     string
	RepoType string
}

type ResourceObjectDO struct {
	Owner string
	Type  string
	Id    string
}

func (do *ResourceObjectDO) toResourceObject(r *domain.ResourceObject) (err error) {
	if r.Owner, err = domain.NewAccount(do.Owner); err != nil {
		return
	}

	if r.Type, err = domain.NewResourceType(do.Type); err != nil {
		return
	}

	r.Id = do.Id

	return
}

func toResourceObjectDO(r *domain.ResourceObject) ResourceObjectDO {
	return ResourceObjectDO{
		Owner: r.Owner.Account(),
		Type:  r.Type.ResourceType(),
		Id:    r.Id,
	}
}

type ResourceIndexDO struct {
	Owner string
	Id    string
}

func (do *ResourceIndexDO) toResourceIndex(r *domain.ResourceIndex) (err error) {
	if r.Owner, err = domain.NewAccount(do.Owner); err != nil {
		return
	}

	r.Id = do.Id

	return
}

func toResourceIndexDO(r *domain.ResourceIndex) ResourceIndexDO {
	return ResourceIndexDO{
		Owner: r.Owner.Account(),
		Id:    r.Id,
	}
}
