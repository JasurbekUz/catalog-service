package service

import (

	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
	l "github.com/JasurbekUz/catalog-service/pkg/logger"

)

func (s CatalogService) CreateAuthor (ctx context.Context, req *cs.Author) (*cs.Author, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}

	req.AuthorId = id.String()
	resp, err := s.storage.Catalog().CreateAuthor(*req)
	if err != nil {
		s.logger.Error("falied to create author", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create author")
	}

	return &resp, nil
}

func (s *CatalogService) GetAuthor (ctx context.Context, req *cs.ByIdReq) (*cs.Author, error) {
	resp, err := s.storage.Catalog().GetAuthor(req.Id)
	if err != nil {
		s.logger.Error("failed to get author", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get author")
	}
	return &resp, nil
}