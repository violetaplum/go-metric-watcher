package repository

import (
	"context"
	"github.com/violetaplum/go-metric-watcher/domain"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"gorm.io/gorm"
	"time"
)

type AlertHistoryRepository struct {
	db *gorm.DB
}

func (a AlertHistoryRepository) SaveAlert(ctx context.Context, history *model.AlertHistory) error {
	return a.db.WithContext(ctx).Create(history).Error
}

func (a AlertHistoryRepository) UpdateAlert(ctx context.Context, history *model.AlertHistory) error {
	return a.db.WithContext(ctx).Save(history).Error
}

func (a AlertHistoryRepository) GetAlertsByTimeRange(ctx context.Context, start, end time.Time) ([]model.AlertHistory, error) {
	var histories []model.AlertHistory
	err := a.db.WithContext(ctx).
		Where("time BETWEEN ? AND ?", start, end).
		Order("time DESC").
		Find(&histories).Error
	return histories, err
}

func (a AlertHistoryRepository) GetAlertsByRuleID(ctx context.Context, ruleID int64) ([]model.AlertHistory, error) {
	var histories []model.AlertHistory
	err := a.db.WithContext(ctx).
		Where("alert_rule_id = ?", ruleID).
		Order("time DESC").
		Find(&histories).Error
	return histories, err
}

func (a AlertHistoryRepository) GetUnresolvedAlerts(ctx context.Context) ([]model.AlertHistory, error) {
	var histories []model.AlertHistory
	err := a.db.WithContext(ctx).
		Where("status = ?", "triggered").
		Order("time DESC").
		Find(&histories).Error
	return histories, err
}

func NewAlertHistoryRepository(db *gorm.DB) domain.AlertHistoryRepository {
	return &AlertHistoryRepository{db: db}
}
