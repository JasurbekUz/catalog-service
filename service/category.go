package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
	l "github.com/JasurbekUz/catalog-service/pkg/logger"
)

func (s *CatalogService) CreateCategory (ctx context.Context, req *cs.Category) (*cs.Category, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}

	req.CategoryId = id.String()
	resp, err := s.storage.Catalog().CreateCategory(*req)
	if err != nil {
		s.logger.Error("falied to create category", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create category")
	}
	return &resp, nil
}

func (s *CatalogService) GetCategory(ctx context.Context, req *cs.ByIdReq) (*cs.Category, error){
	resp, err := s.storage.Catalog().GetCategory(req.Id)
	if err != nil {
		s.logger.Error("falied to get category", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get category")
	}
	return &resp, nil
}

func (s *CatalogService) GetCategories(ctx context.Context, req *cs.ListReq) (*cs.CategoryListResp, error){
	resp, err := s.storage.Catalog().GetCategories(req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to get categories", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get categories")
	}
	return &resp, nil
}

func (s *CatalogService) UpdateCategory(ctx context.Context, req *cs.Category) (*cs.Category, error){
	resp, err := s.storage.Catalog().UpdateCategory(*req)
	if err != nil {
		s.logger.Error("failed to update category", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update category")
	}
	return &resp, nil
}

func (s *CatalogService) DeleteCategory(ctx context.Context, req *cs.ByIdReq) (*cs.Empty, error) {
	err := s.storage.Catalog().DeleteCategory(req.Id)
	if err != nil {
		s.logger.Error("failed to delete category", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete category")
	}
	return &cs.Empty{}, nil
}