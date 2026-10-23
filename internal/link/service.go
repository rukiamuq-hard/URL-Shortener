package link

type Service struct {
	PDB Postgres
}

func NewService(PDB Postgres) *Service {
	return &Service{
		PDB: PDB,
	}
}
